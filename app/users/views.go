package users

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type View struct {
	flow *FLow
}

func NewView(f *FLow) *View {
	return &View{flow: f}
}

func (v *View) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("/", v.createUser)
	}
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required" example:"John"`
	LastName string `json:"last_name" binding:"required" example:"Doe"`
	Email    string `json:"email" binding:"required,email" example:"john.doe@example.com"`
	Tel      string `json:"telephone" binding:"required" example:"+1234567890"`
	Status   Status `json:"status" binding:"required" example:"active"`
}

// CreateUserResponse represents the response for user creation
type CreateUserResponse struct {
	Message string `json:"message" example:"user created"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"validation error"`
}

// createUser creates a new user
// @Summary Create a new user
// @Description Create a new user with name, last name, email, telephone, and status
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User creation request"
// @Success 201 {object} CreateUserResponse "User created successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users [post]
func (f *View) createUser(c *gin.Context) {
	var input CreateUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := f.flow.Insert(context.Background(), input.Name, input.LastName, input.Email, input.Tel, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
