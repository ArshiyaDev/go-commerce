package users

import (
	"net/http"

	"github.com/ArshiyaDev/go-commerce/modules/users/entities"
	"github.com/gin-gonic/gin"
)

type View struct {
	service *Service
}

func NewView(s *Service) *View {
	return &View{service: s}
}

func (v *View) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("/", v.createUser)
	}
}

// CreateUserRequest represents the request body for creating a user

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
func (h *Service) CreateUser(c *gin.Context) {
	name := c.Query("name") // returns empty string if not provided
	lastName := c.Query("last_name")
	email := c.Query("email")

	if name == "" || lastName == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name, last_name, and email are required"})
		return
	}

	user := entities.User{
		Name:     name,
		LastName: lastName,
		Email:    email,
		Status:   "active", // default
	}

	id, err := h.service.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
