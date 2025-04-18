package model

import "github.com/google/uuid"

type (
	GetMonthlySalesReport struct {
		ItemID uuid.UUID `json:"item_id"`
		Month  int       `json:"month"`
		Year   int       `json:"year"`
	}

	MonthlySalesReport struct {
		Month      int     `json:"month"`
		Year       int     `json:"year"`
		TotalSales float64 `json:"total_sales"`
		TotalHPP   float64 `json:"total_hpp"`
		Profit     float64 `json:"profit"`
	}
)
