package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	sinterface "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type supplierService struct {
	SupplierRepository interfaces.SupplierRepository
}

func NewSupplierService(r *repository.Contract) sinterface.SupplierService {
	return &supplierService{
		SupplierRepository: r.Supplier,
	}
}

func (s *supplierService) Get(ctx context.Context, params *model.GetSupplierRequest) (data []model.Supplier, total int64, err error) {
	return s.SupplierRepository.Get(ctx, params)
}

func (s *supplierService) GetById(ctx context.Context, id uuid.UUID) (data *model.Supplier, err error) {
	return s.SupplierRepository.GetById(ctx, id)
}

func (s *supplierService) Create(ctx context.Context, data *model.Supplier) (err error) {
	return s.SupplierRepository.Create(ctx, data)
}

func (s *supplierService) Update(ctx context.Context, data *model.Supplier, id uuid.UUID) (err error) {
	return s.SupplierRepository.Update(ctx, data, id)
}

func (s *supplierService) Delete(ctx context.Context, id uuid.UUID) (err error) {
	return s.SupplierRepository.Delete(ctx, id)
}
