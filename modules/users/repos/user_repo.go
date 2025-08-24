package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ArshiyaDev/go-commerce/modules/users/entities"
	"github.com/google/uuid"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) UserInterface {
	return &Repo{db: db}
}

func (r *Repo) Insert(ctx context.Context, user *entities.User) (uuid.UUID, error) {

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

	return user.ID, nil
}
func (r *Repo) GetById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	query := `SELECT id, name, last_name, email, status, tel, created_at, updated_at 
	          FROM users WHERE id = $1 LIMIT 1`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Email,
		&user.Status,
		&user.Tel,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repo) GetByUsername(ctx context.Context, email string) (*entities.User, error) {
	query := `SELECT id, name, last_name, email, status, tel, created_at, updated_at 
	          FROM users WHERE email = $1 LIMIT 1`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Email,
		&user.Status,
		&user.Tel,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println(err)
	return &user, nil
}
