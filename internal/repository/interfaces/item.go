package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type ItemRepository interface {
	Get(ctx context.Context, params *model.GetItemRequest) (data []model.Item, total int64, err error)
	GetById(ctx context.Context, id uuid.UUID) (data *model.Item, err error)
	Create(ctx context.Context, data *model.Item) (err error)
	Update(ctx context.Context, data *model.Item, id uuid.UUID) (err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
