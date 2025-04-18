package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type StockBatchRepository interface {
	Get(ctx context.Context, params *model.GetStockBatchRequest) (data []model.StockBatchItem, total int64, err error)
	CreateBulk(ctx context.Context, payload []model.StockBatchItem) (err error)
	Update(ctx context.Context, payload *model.StockBatchItem, id uuid.UUID) (err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
