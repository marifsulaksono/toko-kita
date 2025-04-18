package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type (
	GetSaleRequest struct {
		Page      int    `json:"page" query:"page" validate:"gte=1"`
		Limit     int    `json:"limit" query:"limit" validate:"gte=1"`
		Search    string `json:"search" query:"search"`
		StartDate string `json:"start_date" query:"start_date" validate:"datetime=2006-01-02"`
		EndDate   string `json:"end_date" query:"end_date" validate:"datetime=2006-01-02"`
	}

	SaleRequest struct {
		CustomerName string        `json:"customer_name"`
		Date         time.Time     `json:"date" validate:"required"`
		Details      []SalesDetail `json:"sales_details,omitempty" gorm:"foreignKey:SaleID"`
	}

	SalesDetail struct {
		ItemID uuid.UUID `json:"item_id" validate:"required"`
		Qty    int       `json:"qty" validate:"required"`
	}
)

func (s *GetSaleRequest) ParseToModel() *model.GetSaleRequest {
	return &model.GetSaleRequest{
		Page:      s.Page,
		Limit:     s.Limit,
		Search:    s.Search,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
	}
}

func (s *SaleRequest) ParseToModel(userID string) *model.Sale {
	var details []model.SalesDetail
	for _, detail := range s.Details {
		details = append(details, model.SalesDetail{
			ItemID: detail.ItemID,
			Qty:    detail.Qty,
		})
	}

	return &model.Sale{
		UserID:       uuid.MustParse(userID),
		CustomerName: s.CustomerName,
		Date:         s.Date,
		Details:      details,
	}
}
