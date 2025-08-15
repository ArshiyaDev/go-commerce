package users

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type FLow struct {
	repo *Repo
}

func NewFlow(r *Repo) *FLow {
	return &FLow{repo: r}
}

func (f *FLow) Insert(ctx context.Context, user User) (uuid.UUID, error) {
	u := User{
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		Status:    user.Status,
		Tel:       user.Tel,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return f.repo.Insert(ctx, &u)

}
