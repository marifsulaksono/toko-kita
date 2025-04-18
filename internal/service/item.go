package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	sinterface "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type itemService struct {
	ItemRepository interfaces.ItemRepository
}

func NewItemService(r *repository.Contract) sinterface.ItemService {
	return &itemService{
		ItemRepository: r.Item,
	}
}

func (s *itemService) Get(ctx context.Context, params *model.GetItemRequest) (data []model.Item, total int64, err error) {
	return s.ItemRepository.Get(ctx, params)
}

func (s *itemService) GetById(ctx context.Context, id uuid.UUID) (data *model.Item, err error) {
	return s.ItemRepository.GetById(ctx, id)
}

func (s *itemService) Create(ctx context.Context, data *model.Item) (err error) {
	return s.ItemRepository.Create(ctx, data)
}

func (s *itemService) Update(ctx context.Context, data *model.Item, id uuid.UUID) (err error) {
	return s.ItemRepository.Update(ctx, data, id)
}

func (s *itemService) Delete(ctx context.Context, id uuid.UUID) (err error) {
	return s.ItemRepository.Delete(ctx, id)
}
