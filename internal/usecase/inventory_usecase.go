package usecase

import (
	"context"
	"infared-backend/internal/domain"
	"infared-backend/internal/repository"
)

type InventoryUsecase interface {
	GetInventoryByPosko(ctx context.Context, poskoID string) ([]domain.PoskoInventoryDetail, error)
}

type inventoryUsecase struct {
	inventoryRepo repository.InventoryRepository
}

func NewInventoryUsecase(repo repository.InventoryRepository) InventoryUsecase {
	return &inventoryUsecase{inventoryRepo: repo}
}

func (u *inventoryUsecase) GetInventoryByPosko(ctx context.Context, poskoID string) ([]domain.PoskoInventoryDetail, error) {
	return u.inventoryRepo.GetByPoskoID(ctx, poskoID)
}
