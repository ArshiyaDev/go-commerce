package users

import (
	"context"

	"github.com/google/uuid"
)

type UserInterface interface {
	Insert(ctx context.Context, user *User) (uuid.UUID, error)
}
