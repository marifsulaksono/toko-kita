package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type (
	StockBatchItem struct {
		ItemID         uuid.UUID `json:"item_id"`
		SupplierID     uuid.UUID `json:"supplier_id" validate:"required"`
		BatchNo        string    `json:"batch_no" validate:"required,lte=50"`
		PurchasedPrice float64   `json:"purchased_price" validate:"required,gte=0"`
		PurchasedQty   int       `json:"purchased_qty" validate:"required,gte=1"`
		RemainingQty   int       `json:"remaining_qty" validate:"gte=1"`
		PurchasedAt    time.Time `json:"purchased_at" validate:"required"`
	}
)

func (s *StockBatchItem) ParseToModel() model.StockBatchItem {
	return model.StockBatchItem{
		ItemID:         s.ItemID,
		SupplierID:     s.SupplierID,
		BatchNo:        s.BatchNo,
		PurchasedPrice: s.PurchasedPrice,
		PurchasedQty:   s.PurchasedQty,
		RemainingQty:   s.RemainingQty,
		PurchasedAt:    s.PurchasedAt,
	}
}
