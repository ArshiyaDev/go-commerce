package users

import (
	"context"
	"time"

	"github.com/ArshiyaDev/go-commerce/modules/users/entities"
	repo "github.com/ArshiyaDev/go-commerce/modules/users/repos"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo repo.UserInterface
}

func NewService(r repo.UserInterface) *UserService {
	return &UserService{userRepo: r}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	Name string,
	LastName string,
	Email string,
	Tel string) (uuid.UUID, error) {

	u := entities.User{
		Name:      Name,
		LastName:  LastName,
		Email:     Email,
		Status:    entities.Inactive,
		Tel:       Tel,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.userRepo.Insert(ctx, &u)

}

func (s *UserService) GetByUsername(ctx context.Context, email string) (*entities.User, error) {

	return s.userRepo.GetByUsername(ctx, email)
}
