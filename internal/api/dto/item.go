package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type (
	GetItemRequest struct {
		Page   int    `json:"page" query:"page" validate:"gte=1"`
		Limit  int    `json:"limit" query:"limit" validate:"gte=1"`
		Search string `json:"search" query:"search"`
	}

	CreateItemRequest struct {
		SKU             string           `json:"sku" validate:"required,lte=25"`
		Name            string           `json:"name" validate:"required,lte=255"`
		Unit            string           `json:"unit" validate:"required,lte=255"`
		SellingPrice    float64          `json:"selling_price" validate:"required"`
		StockBatchItems []StockBatchItem `json:"stock_batch_items"`
	}

	UpdateItemRequest struct {
		SKU          string  `json:"sku" validate:"required,lte=25"`
		Name         string  `json:"name" validate:"required,lte=255"`
		Unit         string  `json:"unit" validate:"required,lte=255"`
		SellingPrice float64 `json:"selling_price" validate:"required"`
	}

	StockBatchItem struct {
		ItemID         uuid.UUID `json:"item_id"`
		SupplierID     uuid.UUID `json:"supplier_id" validate:"required"`
		BatchNo        string    `json:"batch_no" validate:"required,lte=50"`
		PurchasedPrice float64   `json:"purchased_price" validate:"required"`
		PurchasedQty   int       `json:"purchased_qty" validate:"required"`
		PurchasedAt    time.Time `json:"purchased_at" validate:"required"`
	}
)

func (u *GetItemRequest) ParseToModel() *model.GetItemRequest {
	return &model.GetItemRequest{
		Page:   u.Page,
		Limit:  u.Limit,
		Search: u.Search,
	}
}

func (i *CreateItemRequest) ParseToModel() *model.Item {
	var stockBatches []model.StockBatchItem
	for _, batch := range i.StockBatchItems {
		stockBatches = append(stockBatches, model.StockBatchItem{
			ItemID:         batch.ItemID,
			SupplierID:     batch.SupplierID,
			BatchNo:        batch.BatchNo,
			PurchasedPrice: batch.PurchasedPrice,
			PurchasedQty:   batch.PurchasedQty,
			PurchasedAt:    batch.PurchasedAt,
		})
	}

	return &model.Item{
		SKU:             i.SKU,
		Name:            i.Name,
		Unit:            i.Unit,
		SellingPrice:    i.SellingPrice,
		StockBatchItems: stockBatches,
	}
}

func (i *UpdateItemRequest) ParseToModel() *model.Item {
	return &model.Item{
		SKU:          i.SKU,
		Name:         i.Name,
		Unit:         i.Unit,
		SellingPrice: i.SellingPrice,
	}
}
