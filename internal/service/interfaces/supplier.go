package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type SupplierService interface {
	Get(ctx context.Context, params *model.GetSupplierRequest) (data []model.Supplier, total int64, err error)
	GetById(ctx context.Context, id uuid.UUID) (data *model.Supplier, err error)
	Create(ctx context.Context, data *model.Supplier) (err error)
	Update(ctx context.Context, data *model.Supplier, id uuid.UUID) (err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
