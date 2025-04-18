package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	sinterface "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type purchaseService struct {
	SupplierRepository  interfaces.SupplierRepository
	StockBatcRepository interfaces.StockBatchRepository
}

func NewPurchaseService(r *repository.Contract) sinterface.PurchaseService {
	return &purchaseService{
		SupplierRepository:  r.Supplier,
		StockBatcRepository: r.StockBatch,
	}
}

func (s *purchaseService) Get(ctx context.Context, params *model.GetStockBatchRequest) (data []model.StockBatchItem, total int64, err error) {
	return s.StockBatcRepository.Get(ctx, params)
}

func (s *purchaseService) CreateBulk(ctx context.Context, data []model.StockBatchItem) (err error) {
	return s.StockBatcRepository.CreateBulk(ctx, data)
}

func (s *purchaseService) Update(ctx context.Context, data *model.StockBatchItem, id uuid.UUID) (err error) {
	return s.StockBatcRepository.Update(ctx, data, id)
}

func (s *purchaseService) Delete(ctx context.Context, id uuid.UUID) (err error) {
	return s.StockBatcRepository.Delete(ctx, id)
}
