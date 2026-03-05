package domain

import "time"

type UserRole string

const (
	RoleAdmin   UserRole = "admin"
	RoleRelawan UserRole = "relawan"
)

type User struct {
	ID           string     `json:"id" db:"id"`
	Name         string     `json:"name" db:"name"`
	Email        string     `json:"email" db:"email"`
	PasswordHash string     `json:"-" db:"password_hash"` 
	Role         UserRole   `json:"role" db:"role"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
