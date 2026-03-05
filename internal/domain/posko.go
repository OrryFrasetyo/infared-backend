package domain

import "time"

type Posko struct {
	ID            string     `json:"id" db:"id"`
	Name          string     `json:"name" db:"name"`
	Address       *string    `json:"address" db:"address"`
	Latitude      *float64   `json:"latitude" db:"latitude"`
	Longitude     *float64   `json:"longitude" db:"longitude"`
	CoordinatorID *string    `json:"coordinator_id" db:"coordinator_id"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
