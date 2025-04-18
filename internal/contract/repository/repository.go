package repository

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/common"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
)

type Contract struct {
	User       interfaces.UserRepository
	Auth       interfaces.AuthRepository
	Role       interfaces.RoleRepository
	Supplier   interfaces.SupplierRepository
	Item       interfaces.ItemRepository
	StockBatch interfaces.StockBatchRepository
}

func NewRepository(ctx context.Context, common *common.Contract) (*Contract, error) {
	role := repository.NewRoleRepository(common.DB)
	user := repository.NewUserRepository(common.DB)
	auth := repository.NewAuthRepository(common.DB)
	supplier := repository.NewSupplierRepository(common.DB)
	item := repository.NewItemRepository(common.DB)
	stockBatch := repository.NewStockBatchRepository(common.DB)

	return &Contract{
		User:       user,
		Auth:       auth,
		Role:       role,
		Supplier:   supplier,
		Item:       item,
		StockBatch: stockBatch,
	}, nil
}
