package domain

import "time"

type Item struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Unit      string     `json:"unit" db:"unit"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type PoskoInventory struct {
	ID        string     `json:"id" db:"id"`
	PoskoID   string     `json:"posko_id" db:"posko_id"`
	ItemID    string     `json:"item_id" db:"item_id"`
	Quantity  int        `json:"quantity" db:"quantity"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
