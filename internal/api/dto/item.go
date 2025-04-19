package dto

import (
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type (
	GetItemRequest struct {
		Page   int    `json:"page" query:"page" validate:"gte=1"`
		Limit  int    `json:"limit" query:"limit" validate:"gte=1"`
		Search string `json:"search" query:"search"`
		Sort   string `json:"sort" query:"sort" validate:"omitempty,oneof=name sku unit selling_price created_at"`
		Order  string `json:"order" query:"order" validate:"omitempty,oneof=asc desc"`
	}

	CreateItemRequest struct {
		SKU             string           `json:"sku" validate:"required,lte=25"`
		Name            string           `json:"name" validate:"required,lte=255"`
		Unit            string           `json:"unit" validate:"required,lte=255"`
		SellingPrice    float64          `json:"selling_price" validate:"required,gte=0"`
		StockBatchItems []StockBatchItem `json:"stock_batch_items"`
	}

	UpdateItemRequest struct {
		SKU          string  `json:"sku" validate:"required,lte=25"`
		Name         string  `json:"name" validate:"required,lte=255"`
		Unit         string  `json:"unit" validate:"required,lte=255"`
		SellingPrice float64 `json:"selling_price" validate:"required,gte=0"`
	}
)

func (u *GetItemRequest) ParseToModel() *model.GetItemRequest {
	return &model.GetItemRequest{
		Page:   u.Page,
		Limit:  u.Limit,
		Search: u.Search,
		Sort:   u.Sort,
		Order:  u.Order,
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
