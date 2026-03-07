package repository

import (
	"context"
	"infared-backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type RequestRepository interface {
	CreateRequestWithItems(ctx context.Context, req *domain.LogisticsRequest, items []domain.RequestItem) error
}

type requestRepository struct {
	db *sqlx.DB
}

func NewRequestRepository(db *sqlx.DB) RequestRepository {
	return &requestRepository{db: db}
}

func (r *requestRepository) CreateRequestWithItems(ctx context.Context, req *domain.LogisticsRequest, items []domain.RequestItem) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	queryHeader := `
		INSERT INTO logistics_requests (id, posko_id, requested_by, original_prompt, status, created_at, updated_at)
		VALUES (:id, :posko_id, :requested_by, :original_prompt, :status, :created_at, :updated_at)
	`
	_, err = tx.NamedExecContext(ctx, queryHeader, req)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryItem := `
		INSERT INTO request_items (id, request_id, item_id, quantity, urgency, created_at, updated_at)
		VALUES (:id, :request_id, :item_id, :quantity, :urgency, :created_at, :updated_at)
	`
	for _, item := range items {
		_, err = tx.NamedExecContext(ctx, queryItem, item)
		if err != nil {
			tx.Rollback() 
			return err
		}
	}

	return tx.Commit()
}
