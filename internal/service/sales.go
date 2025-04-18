package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/common"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	sinterface "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type saleService struct {
	SaleRepository interfaces.SaleRepository
	ItemRepository interfaces.ItemRepository
	StockBatch     interfaces.StockBatchRepository
	Common         *common.Contract
}

func NewSaleService(r *repository.Contract, c *common.Contract) sinterface.SaleService {
	return &saleService{
		SaleRepository: r.Sale,
		ItemRepository: r.Item,
		StockBatch:     r.StockBatch,
		Common:         c,
	}
}

func (s *saleService) Get(ctx context.Context, params *model.GetSaleRequest) (data []model.Sale, total int64, err error) {
	return s.SaleRepository.Get(ctx, params)
}

func (s *saleService) GetById(ctx context.Context, id uuid.UUID) (data *model.Sale, err error) {
	return s.SaleRepository.GetByID(ctx, id)
}

func (s *saleService) Create(ctx context.Context, data *model.Sale) (err error) {
	var (
		total float64
	)

	// create database transaction
	tx := s.Common.DB.Begin()
	ctxTx := helper.WithTx(ctx, tx)

	for i := range data.Details {
		detail := &data.Details[i]

		item, err := s.ItemRepository.GetById(ctx, detail.ItemID)
		if err != nil {
			tx.Rollback()
			return err
		}

		stockBatches, err := s.StockBatch.GetByItemID(ctx, detail.ItemID.String(), false)
		if err != nil {
			tx.Rollback()
			return err
		}

		if len(stockBatches) == 0 {
			tx.Rollback()
			return response.NewCustomError(http.StatusBadRequest, fmt.Sprintf("Item %s tidak memiliki stock batch", item.Name), nil)
		}

		var totalStock int
		for _, batch := range stockBatches {
			totalStock += batch.RemainingQty
		}
		if totalStock < detail.Qty {
			tx.Rollback()
			return response.NewCustomError(http.StatusBadRequest, "Stock tidak mencukupi", nil)
		}

		// FIFO logic
		remainingToFulfill := detail.Qty
		var usedBatches []model.SalesBatchItemDetail

		for _, batch := range stockBatches {
			if remainingToFulfill == 0 {
				break
			}
			if batch.RemainingQty <= 0 {
				continue
			}

			useQty := min(batch.RemainingQty, remainingToFulfill)

			// Update batch stock
			newRemaining := batch.RemainingQty - useQty
			if err := s.StockBatch.UpdateStock(ctxTx, batch.ID, newRemaining); err != nil {
				tx.Rollback()
				return err
			}

			// Record sales_batch_link
			usedBatches = append(usedBatches, model.SalesBatchItemDetail{
				BatchNo:        batch.BatchNo,
				Qty:            useQty,
				PurchasedPrice: batch.PurchasedPrice,
			})

			remainingToFulfill -= useQty
		}

		detail.SellingPrice = item.SellingPrice
		detail.SalesBatchItems = usedBatches
		total += float64(detail.Qty) * detail.SellingPrice
	}

	data.Total = total
	if err := s.SaleRepository.Create(ctx, data); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (s *saleService) Delete(ctx context.Context, id uuid.UUID) (err error) {
	return s.SaleRepository.Delete(ctx, id)
}
