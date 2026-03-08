package repository

import (
	"context"
	"infared-backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type RequestRepository interface {
	CreateRequestWithItems(ctx context.Context, req *domain.LogisticsRequest, items []domain.RequestItem) error
	GetAllWithDetails(ctx context.Context) ([]domain.RequestDetail, error)
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

func (r *requestRepository) GetAllWithDetails(ctx context.Context) ([]domain.RequestDetail, error) {
	var requests []domain.RequestDetail

	queryHeader := `
		SELECT 
			r.*, 
			p.name AS posko_name, 
			u.name AS user_name
		FROM logistics_requests r
		JOIN posko p ON r.posko_id = p.id
		JOIN users u ON r.requested_by = u.id
		ORDER BY r.created_at DESC
	`
	err := r.db.SelectContext(ctx, &requests, queryHeader)
	if err != nil {
		return nil, err
	}

	queryItems := `
		SELECT 
			ri.*, 
			i.name AS item_name, 
			i.unit AS item_unit
		FROM request_items ri
		JOIN items i ON ri.item_id = i.id
		WHERE ri.request_id = $1
	`

	for i, req := range requests {
		var items []domain.RequestItemDetail
		err = r.db.SelectContext(ctx, &items, queryItems, req.ID)
		if err != nil {
			return nil, err
		}
		requests[i].Items = items
	}

	return requests, nil
}
