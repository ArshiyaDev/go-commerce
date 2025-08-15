package users

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Insert(ctx context.Context, user *User) (uuid.UUID, error) {

	user.ID = uuid.New()
	query := `INSERT INTO users (id, name, last_name, email, status, tel, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		user.ID,
		user.Name,
		user.LastName,
		user.Email,
		user.Status,
		user.Tel,
		time.Now(), // created_at
		time.Now(), // updated_at
	).Scan(&user.ID)

	if err != nil {
		log.Printf("Failed to insert user : %v", err)

	}

	// if err != nil {
	// 	// Check for UNIQUE constraint violation (email)
	// 	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
	// 		return uuid.UUID{}, fmt.Errorf("email already exists: %s", user.Email)
	// 	}
	// 	return uuid.UUID{}, fmt.Errorf("failed to insert user: %w", err)
	// }

	return user.ID, nil
}
