package service

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type Contract struct {
	User     interfaces.UserService
	Auth     interfaces.AuthService
	Role     interfaces.RoleService
	Supplier interfaces.SupplierService
	Item     interfaces.ItemService
}

func NewService(ctx context.Context, r *repository.Contract) (*Contract, error) {
	user := service.NewUserService(r)
	auth := service.NewAuthService(r)
	role := service.NewRoleService(r)
	supplier := service.NewSupplierService(r)
	item := service.NewItemService(r)

	return &Contract{
		User:     user,
		Auth:     auth,
		Role:     role,
		Supplier: supplier,
		Item:     item,
	}, nil
}
