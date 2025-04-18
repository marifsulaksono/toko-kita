package repository

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"gorm.io/gorm"
)

type stockBatchRepository struct {
	DB *gorm.DB
}

func NewStockBatchRepository(db *gorm.DB) interfaces.StockBatchRepository {
	return &stockBatchRepository{
		DB: db,
	}
}

func (r *stockBatchRepository) Get(ctx context.Context, params *model.GetStockBatchRequest) (data []model.StockBatchItem, total int64, err error) {
	var (
		offset = (params.Page - 1) * params.Limit
	)

	db := r.DB.WithContext(ctx)
	if params.ItemID != "" {
		db = db.Where("item_id = ?", params.ItemID)
	}

	if params.SupplierID != "" {
		db = db.Where("supplier_id = ?", params.SupplierID)
	}

	var startDate, endDate time.Time
	if params.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", params.StartDate)
		if err != nil {
			return nil, 0, response.NewCustomError(http.StatusBadRequest, "Invalid start date format", err)
		}
		db = db.Where("purchased_at >= ?", startDate)
	}

	if params.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", params.EndDate)
		if err != nil {
			return nil, 0, response.NewCustomError(http.StatusBadRequest, "Invalid end date format", err)
		}

		if !startDate.IsZero() && endDate.Before(startDate) {
			return nil, 0, response.NewCustomError(http.StatusBadRequest, "End date must be greater than or equal to start date", nil)
		}

		db = db.Where("purchased_at < ?", endDate.AddDate(0, 0, 1))
	}

	if params.Search != "" {
		db = db.Joins("LEFT JOIN items ON items.id = stock_batch_items.item_id").
			Joins("LEFT JOIN suppliers ON suppliers.id = stock_batch_items.supplier_id").
			Where("items.name ILIKE ? OR items.sku ILIKE ? OR suppliers.name ILIKE ?",
				"%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%")
	}

	err = db.Preload("Item").Preload("Supplier").Offset(offset).Limit(params.Limit).Order("purchased_at ASC").Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&model.StockBatchItem{}).Where("deleted_at IS NULL").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return
}

func (r *stockBatchRepository) GetByItemID(ctx context.Context, itemID string, isShowZeroStock bool) (data []model.StockBatchItem, err error) {
	db := r.DB.WithContext(ctx)

	if isShowZeroStock {
		db = db.Where("remaining_qty > 0")
	}
	err = db.Where("item_id = ? AND deleted_at IS NULL", itemID).Preload("Item").Order("purchased_at ASC").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return
}

func (r *stockBatchRepository) CreateBulk(ctx context.Context, payload []model.StockBatchItem) error {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		userID = ""
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for i := range payload {
		payload[i].CreatedBy = userID
		payload[i].RemainingQty = payload[i].PurchasedQty

		var itemCount int64
		if err := tx.Model(&model.Item{}).Where("id = ?", payload[i].ItemID).Count(&itemCount).Error; err != nil {
			tx.Rollback()
			return err
		}
		if itemCount == 0 {
			tx.Rollback()
			return response.NewCustomError(http.StatusBadRequest, "Item tidak ditemukan", nil)
		}

		var supplierCount int64
		if err := tx.Model(&model.Supplier{}).Where("id = ?", payload[i].SupplierID).Count(&supplierCount).Error; err != nil {
			tx.Rollback()
			return err
		}
		if supplierCount == 0 {
			tx.Rollback()
			return response.NewCustomError(http.StatusBadRequest, "Supplier tidak ditemukan", nil)
		}
	}

	if err := tx.Create(&payload).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *stockBatchRepository) Update(ctx context.Context, payload *model.StockBatchItem, id uuid.UUID) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		userID = ""
	}
	payload.UpdatedBy = userID

	var existingStockBatch model.StockBatchItem
	if err := r.DB.Where("id = ?", id).First(&existingStockBatch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewCustomError(http.StatusNotFound, "Stock batch tidak ditemukan", err)
		}
		return response.NewCustomError(http.StatusInternalServerError, "Gagal mengambil data item", err)
	}

	var itemCount int64
	if err := r.DB.Model(&model.Item{}).Where("id = ?", payload.ItemID).Count(&itemCount).Error; err != nil {
		return err
	}
	if itemCount == 0 {
		return response.NewCustomError(http.StatusBadRequest, "Item tidak ditemukan", nil)
	}

	var supplierCount int64
	if err := r.DB.Model(&model.Supplier{}).Where("id = ?", payload.SupplierID).Count(&supplierCount).Error; err != nil {
		return err
	}
	if supplierCount == 0 {
		return response.NewCustomError(http.StatusBadRequest, "Supplier tidak ditemukan", nil)
	}

	existingStockBatch.ItemID = payload.ItemID
	existingStockBatch.SupplierID = payload.SupplierID
	existingStockBatch.BatchNo = payload.BatchNo
	existingStockBatch.PurchasedPrice = payload.PurchasedPrice
	existingStockBatch.PurchasedQty = payload.PurchasedQty
	existingStockBatch.RemainingQty = payload.RemainingQty
	existingStockBatch.PurchasedAt = payload.PurchasedAt
	existingStockBatch.UpdatedBy = payload.UpdatedBy

	if err := r.DB.Save(&existingStockBatch).Error; err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal mengupdate stock batch", err)
	}

	return nil

}

func (r *stockBatchRepository) UpdateStock(ctx context.Context, id uuid.UUID, finalQty int) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		userID = ""
	}

	db := helper.GetTx(ctx, r.DB)
	var existingStockBatch model.StockBatchItem
	if err := db.Where("id = ?", id).First(&existingStockBatch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewCustomError(http.StatusNotFound, "Stock batch tidak ditemukan", err)
		}
		return response.NewCustomError(http.StatusInternalServerError, "Gagal mengambil data item", err)
	}

	existingStockBatch.RemainingQty = finalQty
	existingStockBatch.UpdatedBy = userID

	if err := db.Save(&existingStockBatch).Error; err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal mengupdate stock batch", err)
	}

	return nil

}

func (r *stockBatchRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		userID = ""
	}

	var itemCount int64
	if err := r.DB.Model(&model.StockBatchItem{}).Where("id = ?", id).Count(&itemCount).Error; err != nil {
		return err
	}
	if itemCount == 0 {
		return response.NewCustomError(http.StatusBadRequest, "Stock batch tidak ditemukan", nil)
	}

	// Set deleted_by field to the user who deleted the record
	err = r.DB.Model(&model.StockBatchItem{}).
		Where("id = ?", id).
		Update("deleted_by", userID).Error
	if err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal menghapus data", err)
	}

	if err := r.DB.Delete(&model.StockBatchItem{}, id).Error; err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal menghapus data", err)
	}

	return nil
}
