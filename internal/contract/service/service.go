package service

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/common"
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
	Purchase interfaces.PurchaseService
	Sale     interfaces.SaleService
}

func NewService(ctx context.Context, r *repository.Contract, c *common.Contract) (*Contract, error) {
	user := service.NewUserService(r)
	auth := service.NewAuthService(r)
	role := service.NewRoleService(r)
	supplier := service.NewSupplierService(r)
	item := service.NewItemService(r)
	purchase := service.NewPurchaseService(r)
	sale := service.NewSaleService(r, c)

	return &Contract{
		User:     user,
		Auth:     auth,
		Role:     role,
		Supplier: supplier,
		Item:     item,
		Purchase: purchase,
		Sale:     sale,
	}, nil
}
