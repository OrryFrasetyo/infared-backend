package usecase

import (
	"context"
	"errors"
	"infared-backend/internal/domain"
	"infared-backend/internal/repository"
	"infared-backend/pkg/utils"
	"time"
)

type UserUsecase interface {
	RegisterRelawan(ctx context.Context, name, email, plainPassword string) error
	Login(ctx context.Context, email, plainPassword string) (string, *domain.User, error)
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) RegisterRelawan(ctx context.Context, name, email, plainPassword string) error {
	hashedPassword, err := utils.HashPassword(plainPassword)
	if err != nil {
		return errors.New("gagal memproses keamanan password")
	}

	user := &domain.User{
		ID:           utils.GenerateID("usr"),
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
		Role:         domain.RoleRelawan,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, email, plainPassword string) (string, *domain.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, errors.New("email atau password salah")
	}

	if !utils.CheckPasswordHash(plainPassword, user.PasswordHash) {
		return "", nil, errors.New("email atau password salah")
	}

	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return "", nil, errors.New("gagal membuat sesi login")
	}

	return token, user, nil
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return u.userRepo.GetAll(ctx)
}
