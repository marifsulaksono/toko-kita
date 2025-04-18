package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	GetSaleRequest struct {
		Page      int    `json:"page"`
		Limit     int    `json:"limit"`
		Search    string `json:"search"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	Sale struct {
		ID           uuid.UUID     `json:"id" gorm:"primaryKey;type:varchar(36)"`
		UserID       uuid.UUID     `json:"user_id" gorm:"type:uuid;not null"`
		User         User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
		CustomerName string        `json:"customer_name" gorm:"type:varchar(255);not null"`
		Date         time.Time     `json:"date" gorm:"not null"`
		Total        float64       `json:"total" gorm:"not null"`
		Details      []SalesDetail `json:"sales_details,omitempty" gorm:"foreignKey:SaleID"`
		Model
	}

	SalesDetail struct {
		ID              uuid.UUID              `json:"id" gorm:"primaryKey;type:varchar(36)"`
		SaleID          uuid.UUID              `json:"sale_id" gorm:"type:uuid;not null"`
		ItemID          uuid.UUID              `json:"item_id" gorm:"type:uuid;not null"`
		Item            Item                   `json:"item" gorm:"foreignKey:ItemID;references:ID"`
		Qty             int                    `json:"qty" gorm:"not null"`
		SellingPrice    float64                `json:"selling_price" gorm:"not null"`
		SalesBatchItems []SalesBatchItemDetail `json:"sales_batch_items,omitempty" gorm:"foreignKey:SalesDetailID"`
	}

	SalesBatchItemDetail struct {
		ID             uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
		SalesDetailID  uuid.UUID `json:"sales_detail_id" gorm:"type:uuid;not null"`
		BatchNo        string    `json:"batch_id" gorm:"varchar(50);not null"`
		Qty            int       `json:"qty" gorm:"not null"`
		PurchasedPrice float64   `json:"purchased_price" gorm:"not null"`
	}
)

func (s *Sale) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}

func (s *SalesDetail) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}

func (s *SalesBatchItemDetail) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
