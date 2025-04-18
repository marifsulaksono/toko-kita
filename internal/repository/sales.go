package repository

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"gorm.io/gorm"
)

type saleRepository struct {
	DB *gorm.DB
}

func NewSaleRepository(db *gorm.DB) interfaces.SaleRepository {
	return &saleRepository{
		DB: db,
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
	db := r.DB.WithContext(ctx)

	err = db.Where("id = ?", id).Preload("User").Preload("Details.Item").
		Preload("Details.Item.StockBatchItems").
		First(&data).Error
	if err != nil {
		return nil, err
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
