package repository

import (
	"context"
	"infared-backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type InventoryRepository interface {
	GetByPoskoID(ctx context.Context, poskoID string) ([]domain.PoskoInventoryDetail, error)
}

type inventoryRepository struct {
	db *sqlx.DB
}

func NewInventoryRepository(db *sqlx.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetByPoskoID(ctx context.Context, poskoID string) ([]domain.PoskoInventoryDetail, error) {
	inventory := []domain.PoskoInventoryDetail{}

	query := `
		SELECT 
			pi.*, 
			i.name AS item_name, 
			i.unit AS item_unit
		FROM posko_inventory pi
		JOIN items i ON pi.item_id = i.id
		WHERE pi.posko_id = $1 AND pi.deleted_at IS NULL
		ORDER BY i.name ASC
	`

	err := r.db.SelectContext(ctx, &inventory, query, poskoID)
	return inventory, err
}
