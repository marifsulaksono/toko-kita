package repository

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"gorm.io/gorm"
)

type itemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) interfaces.ItemRepository {
	return &itemRepository{
		DB: db,
	}
}

func (r *itemRepository) Get(ctx context.Context, params *model.GetItemRequest) (data []model.Item, total int64, err error) {
	var (
		offset = (params.Page - 1) * params.Limit
	)

	db := r.DB
	if params.Search != "" {
		db = db.Where("name ILIKE ? OR sku ILIKE ?", "%"+params.Search+"%", "%"+params.Search+"%")
	}

	err = db.Offset(offset).Limit(params.Limit).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&model.Item{}).Where("deleted_at IS NULL").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return
}

func (r *itemRepository) GetById(ctx context.Context, id uuid.UUID) (data *model.Item, err error) {
	err = r.DB.Preload("StockBatchItems").Preload("StockBatchItems.Item").Preload("StockBatchItems.Supplier").Where("id = ?", id).First(&data, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewCustomError(http.StatusNotFound, "Data item tidak ditemukan", nil)
		}
		return nil, response.NewCustomError(http.StatusInternalServerError, "Terjadi kesalahan pada server", err)
	}
	return
}

func (r *itemRepository) Create(ctx context.Context, payload *model.Item) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		userID = ""
	}

	payload.CreatedBy = userID
	if len(payload.StockBatchItems) > 0 {
		for i := range payload.StockBatchItems {
			payload.StockBatchItems[i].ItemID = payload.ID
			payload.StockBatchItems[i].RemainingQty = payload.StockBatchItems[i].PurchasedQty
			payload.StockBatchItems[i].CreatedBy = userID
		}
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(&payload).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *itemRepository) Update(ctx context.Context, payload *model.Item, id uuid.UUID) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		payload.UpdatedBy = ""
	} else {
		payload.UpdatedBy = userID
	}

	var existingItem model.Item
	if err := r.DB.Where("id = ?", id).First(&existingItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewCustomError(http.StatusNotFound, "Item tidak ditemukan", err)
		}
		return response.NewCustomError(http.StatusInternalServerError, "Gagal mengambil data item", err)
	}

	updates := map[string]interface{}{
		"sku":           payload.SKU,
		"name":          payload.Name,
		"unit":          payload.Unit,
		"selling_price": payload.SellingPrice,
		"updated_by":    userID,
		"updated_at":    time.Now(),
	}

	if err := r.DB.Model(&model.Item{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (r *itemRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		userID = ""
	}

	// Set deleted_by field to the user who deleted the record
	err = r.DB.Model(&model.Item{}).
		Where("id = ?", id).
		Update("deleted_by", userID).Error
	if err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal menghapus data", err)
	}

	if err := r.DB.Delete(&model.Item{}, id).Error; err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal menghapus data", err)
	}

	return nil
}
