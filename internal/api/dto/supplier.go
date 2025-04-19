package dto

import "github.com/marifsulaksono/go-echo-boilerplate/internal/model"

type (
	GetSupplierRequest struct {
		Page   int    `json:"page" query:"page" validate:"gte=1"`
		Limit  int    `json:"limit" query:"limit" validate:"gte=1"`
		Search string `json:"search" query:"search"`
		Sort   string `json:"sort" query:"sort" validate:"omitempty,oneof=name phone_number created_at"`
		Order  string `json:"order" query:"order" validate:"omitempty,oneof=asc desc"`
	}

	SupplierRequest struct {
		Name        string `json:"name" validate:"required"`
		PhoneNumber string `json:"phone_number" validate:"required"`
		Address     string `json:"address" validate:"required"`
	}
)

func (u *GetSupplierRequest) ParseToModel() *model.GetSupplierRequest {
	return &model.GetSupplierRequest{
		Page:   u.Page,
		Limit:  u.Limit,
		Search: u.Search,
		Sort:   u.Sort,
		Order:  u.Order,
	}
}

func (u *SupplierRequest) ParseToModel() *model.Supplier {
	return &model.Supplier{
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
	}
}
