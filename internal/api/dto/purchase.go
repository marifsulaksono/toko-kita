package dto

import "github.com/marifsulaksono/go-echo-boilerplate/internal/model"

type (
	GetPurchaseRequest struct {
		Page       int    `json:"page" query:"page" validate:"gte=1"`
		Limit      int    `json:"limit" query:"limit" validate:"gte=1"`
		Search     string `json:"search" query:"search"`
		ItemID     string `json:"item_id" query:"item_id"`
		SupplierID string `json:"supplier_id" query:"supplier_id"`
		StartDate  string `json:"start_date" query:"start_date" validate:"datetime=2006-01-02"`
		EndDate    string `json:"end_date" query:"end_date" validate:"datetime=2006-01-02"`
	}
)

func (u *GetPurchaseRequest) ParseToModel() *model.GetStockBatchRequest {
	return &model.GetStockBatchRequest{
		Page:      u.Page,
		Limit:     u.Limit,
		Search:    u.Search,
		StartDate: u.StartDate,
		EndDate:   u.EndDate,
	}
}
