package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*

This file is for role model
You can adjust the structure's field as your need

*/

type (
	GetRoleRequest struct {
		Page   int    `json:"page" query:"page" validate:"gte=1"`
		Limit  int    `json:"limit" query:"limit" validate:"gte=1"`
		Search string `json:"search" query:"search"`
	}

	Role struct {
		ID   uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
		Name string    `json:"name" gorm:"type:varchar(100)"`
		Model
	}
)

func (u *Role) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
