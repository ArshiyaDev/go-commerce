package repo

import (
	"context"

	"github.com/ArshiyaDev/go-commerce/modules/users/entities"
	"github.com/google/uuid"
)

type UserInterface interface {
	Insert(ctx context.Context, user *entities.User) (uuid.UUID, error)
}
