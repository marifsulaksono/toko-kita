package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type PurchaseService interface {
	Get(ctx context.Context, params *model.GetStockBatchRequest) (data []model.StockBatchItem, total int64, err error)
	CreateBulk(ctx context.Context, data []model.StockBatchItem) (err error)
	Update(ctx context.Context, data *model.StockBatchItem, id uuid.UUID) (err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
