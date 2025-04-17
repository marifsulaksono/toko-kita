package dto

import "github.com/marifsulaksono/go-echo-boilerplate/internal/model"

type (
	GetRoleRequest struct {
		Page   int    `json:"page" query:"page" validate:"gte=1"`
		Limit  int    `json:"limit" query:"limit" validate:"gte=1"`
		Search string `json:"search" query:"search"`
	}

	RoleRequest struct {
		Name string `json:"name" validate:"required"`
	}
)

func (r *GetRoleRequest) ParseToModel() *model.GetRoleRequest {
	return &model.GetRoleRequest{
		Page:   r.Page,
		Limit:  r.Limit,
		Search: r.Search,
	}
}

func (r *RoleRequest) ParseToModel() *model.Role {
	return &model.Role{
		Name: r.Name,
	}
}
