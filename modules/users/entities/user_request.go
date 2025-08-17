package entities

import (
	"github.com/ArshiyaDev/go-commerce/common/helpers"
	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required" example:"John"`
	LastName string `json:"last_name" binding:"required" example:"Doe"`
	Email    string `json:"email" binding:"required,email" example:"john.doe@example.com"`
	Tel      string `json:"telephone" binding:"required" example:"+1234567890"`
	Status   Status `json:"status" binding:"required" example:"active"`
}

func (req CreateUserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)

	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidatorFormatErrors(errs)
		}
		return err
	}

	return nil
}
