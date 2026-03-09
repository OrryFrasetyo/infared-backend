package repository

import (
	"context"
	"infared-backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type PoskoRepository interface {
	GetAll(ctx context.Context) ([]domain.Posko, error)
}

type poskoRepository struct {
	db *sqlx.DB
}

func NewPoskoRepository(db *sqlx.DB) PoskoRepository {
	return &poskoRepository{db: db}
}

func (r *poskoRepository) GetAll(ctx context.Context) ([]domain.Posko, error) {
	poskos := []domain.Posko{}

	query := `
		SELECT id, name, address, latitude, longitude, coordinator_id, created_at, updated_at 
		FROM posko 
		WHERE deleted_at IS NULL 
		ORDER BY name ASC
	`

	err := r.db.SelectContext(ctx, &poskos, query)
	return poskos, err
}
