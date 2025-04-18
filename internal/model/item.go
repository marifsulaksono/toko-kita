package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	GetItemRequest struct {
		Page   int    `json:"page"`
		Limit  int    `json:"limit"`
		Search string `json:"search"`
	}

	Item struct {
		ID              uuid.UUID        `json:"id" gorm:"primaryKey;type:varchar(36)"`
		SKU             string           `json:"sku" gorm:"type:varchar(25);unique;not null"`
		Name            string           `json:"name" gorm:"type:varchar(255);not null"`
		Unit            string           `json:"unit" gorm:"type:varchar(255);not null"`
		SellingPrice    float64          `json:"selling_price" gorm:"not null"`
		Stock           int              `json:"stock" gorm:"-:migration"`
		StockBatchItems []StockBatchItem `json:"stock_batch_items,omitempty" gorm:"foreignKey:ItemID"`
		Model
	}
)

func (s *Item) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
