package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	GetStockBatchRequest struct {
		Page       int    `json:"page"`
		Limit      int    `json:"limit"`
		Search     string `json:"search"`
		ItemID     string `json:"item_id"`
		SupplierID string `json:"supplier_id"`
		StartDate  string `json:"start_date"`
		EndDate    string `json:"end_date"`
	}

	StockBatchItem struct {
		ID             uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
		ItemID         uuid.UUID `json:"item_id" gorm:"type:uuid;not null"`
		Item           Item      `json:"item" gorm:"foreignKey:ItemID;references:ID"`
		SupplierID     uuid.UUID `json:"supplier_id" gorm:"type:uuid;not null"`
		Supplier       Supplier  `json:"supplier" gorm:"foreignKey:SupplierID;references:ID"`
		BatchNo        string    `json:"batch_no" gorm:"type:varchar(50);not null"`
		PurchasedPrice float64   `json:"purchased_price" gorm:"not null"`
		PurchasedQty   int       `json:"purchased_qty" gorm:"not null"`
		RemainingQty   int       `json:"remaining_qty" gorm:"not null"`
		PurchasedAt    time.Time `json:"purchased_at" gorm:"not null"`
		Model
	}
)

func (s *StockBatchItem) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
