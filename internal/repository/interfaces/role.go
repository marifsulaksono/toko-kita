package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type RoleRepository interface {
	Get(ctx context.Context, payload *model.GetRoleRequest) (data []model.Role, total int64, err error)
	GetById(ctx context.Context, id uuid.UUID) (data *model.Role, err error)
	Create(ctx context.Context, payload *model.Role) (string, error)
	Update(ctx context.Context, payload *model.Role, id uuid.UUID) (string, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
