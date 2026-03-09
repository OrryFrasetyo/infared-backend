package usecase

import (
	"context"
	"infared-backend/internal/domain"
	"infared-backend/internal/repository"
)

type PoskoUsecase interface {
	GetAllPoskos(ctx context.Context) ([]domain.Posko, error)
}

type poskoUsecase struct {
	poskoRepo repository.PoskoRepository
}

func NewPoskoUsecase(repo repository.PoskoRepository) PoskoUsecase {
	return &poskoUsecase{poskoRepo: repo}
}

func (u *poskoUsecase) GetAllPoskos(ctx context.Context) ([]domain.Posko, error) {
	return u.poskoRepo.GetAll(ctx)
}
