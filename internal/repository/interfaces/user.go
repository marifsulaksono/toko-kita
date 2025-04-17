package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type UserRepository interface {
	Get(ctx context.Context, payload *model.GetUserRequest) ([]model.User, int64, error)
	GetById(ctx context.Context, id uuid.UUID) (data *model.User, err error)
	GetByEmail(ctx context.Context, email string) (data *model.User, err error)
	Create(ctx context.Context, payload *model.User) (string, error)
	Update(ctx context.Context, payload *model.User, id uuid.UUID) (string, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
