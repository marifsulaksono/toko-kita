package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type SaleRepository interface {
	Get(ctx context.Context, params *model.GetSaleRequest) (data []model.Sale, total int64, err error)
	GetByID(ctx context.Context, id uuid.UUID) (data *model.Sale, err error)
	Create(ctx context.Context, payload *model.Sale) error
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
