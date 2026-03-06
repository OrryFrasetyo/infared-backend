package usecase

import (
	"context"
	"infared-backend/internal/domain"
	"infared-backend/internal/repository"
	"infared-backend/pkg/utils"
	"time"
)

type ItemUsecase interface {
	CreateItem(ctx context.Context, name, unit string) (*domain.Item, error)
	GetAllItems(ctx context.Context) ([]domain.Item, error)
}

type itemUsecase struct {
	itemRepo repository.ItemRepository
}

func NewItemUsecase(repo repository.ItemRepository) ItemUsecase {
	return &itemUsecase{itemRepo: repo}
}

func (u *itemUsecase) CreateItem(ctx context.Context, name, unit string) (*domain.Item, error) {
	item := &domain.Item{
		ID:        utils.GenerateID("itm"),
		Name:      name,
		Unit:      unit,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.itemRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (u *itemUsecase) GetAllItems(ctx context.Context) ([]domain.Item, error) {
	return u.itemRepo.GetAll(ctx)
}
