package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	GetSupplierRequest struct {
		Page   int    `json:"page"`
		Limit  int    `json:"limit"`
		Search string `json:"search"`
		Sort   string `json:"sort"`
		Order  string `json:"order"`
	}

	Supplier struct {
		ID          uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
		Name        string    `json:"name" gorm:"type:varchar(255);not null"`
		PhoneNumber string    `json:"phone_number" gorm:"type:varchar(30);not null"`
		Address     string    `json:"address" gorm:"type:text;not null"`
		Model
	}
)

func (s *Supplier) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
