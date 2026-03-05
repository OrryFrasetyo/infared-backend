package domain

import "time"

type RequestStatus string
type UrgencyLevel string

const (
	StatusPending    RequestStatus = "pending"
	StatusProcessing RequestStatus = "processing"
	StatusFulfilled  RequestStatus = "fulfilled"

	UrgencyRendah UrgencyLevel = "rendah"
	UrgencySedang UrgencyLevel = "sedang"
	UrgencyTinggi UrgencyLevel = "tinggi"
	UrgencyKritis UrgencyLevel = "kritis"
)

type LogisticsRequest struct {
	ID             string        `json:"id" db:"id"`
	PoskoID        string        `json:"posko_id" db:"posko_id"`
	RequestedBy    string        `json:"requested_by" db:"requested_by"`
	OriginalPrompt *string       `json:"original_prompt" db:"original_prompt"`
	Status         RequestStatus `json:"status" db:"status"`
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time    `json:"deleted_at,omitempty" db:"deleted_at"`
}

type RequestItem struct {
	ID        string       `json:"id" db:"id"`
	RequestID string       `json:"request_id" db:"request_id"`
	ItemID    string       `json:"item_id" db:"item_id"`
	Quantity  int          `json:"quantity" db:"quantity"`
	Urgency   UrgencyLevel `json:"urgency" db:"urgency"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}
