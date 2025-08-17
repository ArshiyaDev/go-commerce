package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"  db:"id"`
	Name     string    `json:"name" db:"name"`
	LastName string    `json:"last_name" db:"last_name"`
	// Need Validation for email format
	Email     string    `json:"email" db:"email"`
	Status    Status    `json:"status" db:"status"`
	Tel       string    `json:"telephone" db:"telephone"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
