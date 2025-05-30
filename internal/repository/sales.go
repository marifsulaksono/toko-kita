package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	getSalesDataKey = "get-sales-data"
)

type saleRepository struct {
	DB  *gorm.DB
	rdb *redis.Client
}

func NewSaleRepository(db *gorm.DB, rdb *redis.Client) interfaces.SaleRepository {
	return &saleRepository{
		DB:  db,
		rdb: rdb,
	}
}

func (r *saleRepository) Get(ctx context.Context, params *model.GetSaleRequest) (data []model.Sale, total int64, err error) {
	var (
		offset = (params.Page - 1) * params.Limit
		db     = r.DB.WithContext(ctx)
	)

	var startDate, endDate time.Time
	if params.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", params.StartDate)
		if err != nil {
			return nil, 0, response.NewCustomError(http.StatusBadRequest, "Invalid start date format", err)
		}
		db = db.Where("date >= ?", startDate)
	}

	if params.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", params.EndDate)
		if err != nil {
			return nil, 0, response.NewCustomError(http.StatusBadRequest, "Invalid end date format", err)
		}

		if !startDate.IsZero() && endDate.Before(startDate) {
			return nil, 0, response.NewCustomError(http.StatusBadRequest, "End date must be greater than or equal to start date", nil)
		}

		db = db.Where("date < ?", endDate.AddDate(0, 0, 1))
	}

	if params.Search != "" {
		db = db.Where("customer_name ILIKE ?", "%"+params.Search+"%")
	}

	err = db.Preload("User").Offset(offset).Limit(params.Limit).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&model.StockBatchItem{}).Where("deleted_at IS NULL").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return
}

func (r *saleRepository) GetByID(ctx context.Context, id uuid.UUID) (data *model.Sale, err error) {
	redisKey := fmt.Sprintf("%s:%s", getSalesDataKey, id.String())
	cachedSalesID, err := r.rdb.Get(ctx, redisKey).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(cachedSalesID), &data); err != nil {
			log.Printf("Error unmarshaling user data from Redis: %v", err)
		} else {
			fmt.Println("get data from redis")
			return data, nil
		}
	} else if err != redis.Nil {
		log.Printf("Error fetching user data from Redis: %v", err)
	}

	db := r.DB.WithContext(ctx)

	err = db.Where("id = ?", id).Preload("User").Preload("Details.Item").
		Preload("Details.Item.StockBatchItems").
		First(&data).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("get data from database")

	err = helper.SetRedisJSONCache(ctx, r.rdb, redisKey, data, time.Duration(300)*time.Second)
	if err != nil {
		log.Printf("Error setting user data in Redis: %v", err)
	}

	return
}

func (r *saleRepository) Create(ctx context.Context, payload *model.Sale) error {
	err := r.DB.WithContext(ctx).Create(&payload).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *saleRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		userID = ""
	}

	var itemCount int64
	if err := r.DB.Model(&model.Sale{}).Where("id = ?", id).Count(&itemCount).Error; err != nil {
		return err
	}
	if itemCount == 0 {
		return response.NewCustomError(http.StatusBadRequest, "Data penjualan tidak ditemukan", nil)
	}

	// Set deleted_by field to the user who deleted the record
	err = r.DB.Model(&model.Sale{}).
		Where("id = ?", id).
		Update("deleted_by", userID).Error
	if err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal menghapus data", err)
	}

	if err := r.DB.Delete(&model.Sale{}, id).Error; err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal menghapus data", err)
	}

	return nil
}

func (r *saleRepository) GetMonthlySalesReport(ctx context.Context, params *model.GetMonthlySalesReport) (result model.MonthlySalesReport, err error) {

	fmt.Println("Month:", params.Month)
	fmt.Println("Year:", params.Year)

	query := r.DB.WithContext(ctx).Debug().Table("sales s").
		Select(`
			EXTRACT(MONTH FROM s.date) AS month,
			EXTRACT(YEAR FROM s.date) AS year,
			SUM(sd.qty * sd.selling_price) AS total_sales,
			SUM(sbid.qty * sbid.purchased_price) AS total_hpp,
			SUM(sd.qty * sd.selling_price) - SUM(sbid.qty * sbid.purchased_price) AS profit
		`).
		Joins("JOIN sales_details sd ON sd.sale_id = s.id").
		Joins("JOIN sales_batch_item_details sbid ON sbid.sales_detail_id = sd.id")

	// filter by ItemID
	if params.ItemID != uuid.Nil {
		query = query.Where("item_id = ?", params.ItemID)
	}

	// filter by Month
	if params.Month > 0 {
		query = query.Where("CAST(EXTRACT(MONTH FROM s.date) AS INT) = ?", params.Month)
	}

	// filter by Year
	if params.Year > 0 {
		query = query.Where("CAST(EXTRACT(YEAR FROM s.date) AS INT) = ?", params.Year)
	}

	err = query.Group("1, 2").Order("1, 2").Scan(&result).Error
	if err != nil {
		return model.MonthlySalesReport{}, err
	}

	return result, nil
}
