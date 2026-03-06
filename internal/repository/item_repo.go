package repository

import (
	"context"
	"infared-backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type ItemRepository interface {
	Create(ctx context.Context, item *domain.Item) error
	GetAll(ctx context.Context) ([]domain.Item, error)
}

type itemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Create(ctx context.Context, item *domain.Item) error {
	query := `
		INSERT INTO items (id, name, unit, created_at, updated_at)
		VALUES (:id, :name, :unit, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, item)
	return err
}

func (r *itemRepository) GetAll(ctx context.Context) ([]domain.Item, error) {
	items := []domain.Item{}

	query := `SELECT id, name, unit, created_at, updated_at FROM items WHERE deleted_at IS NULL ORDER BY name ASC`

	err := r.db.SelectContext(ctx, &items, query)
	return items, err
}
