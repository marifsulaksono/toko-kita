package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type SaleService interface {
	Get(ctx context.Context, params *model.GetSaleRequest) (data []model.Sale, total int64, err error)
	GetMonthlySalesReport(ctx context.Context, params *model.GetMonthlySalesReport) (model.MonthlySalesReport, error)
	GetById(ctx context.Context, id uuid.UUID) (data *model.Sale, err error)
	Create(ctx context.Context, data *model.Sale) (err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
