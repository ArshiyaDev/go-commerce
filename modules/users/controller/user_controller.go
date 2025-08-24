package controller

import (
	"fmt"
	"net/http"

	"github.com/ArshiyaDev/go-commerce/modules/users/entities"
	users "github.com/ArshiyaDev/go-commerce/modules/users/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService users.UserService
}

func NewUserController(service users.UserService) *UserController {
	return &UserController{userService: service}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a user with name, last name, email, and telephone
// @Tags users
// @Accept json
// @Produce json
// @Param request body entities.CreateUserRequest true "Create User Request"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/v1/users [post]
func (controller *UserController) Create(ctx *gin.Context) {

	createUserReq := entities.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&createUserReq)

	if err != nil {
		panic(err)
	}
	controller.userService.CreateUser(ctx,
		createUserReq.Name, createUserReq.LastName, createUserReq.Email, createUserReq.Tel)

	webResponse := Response{Code: http.StatusOK, Status: "Ok"}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// GetUserByUsername godoc
// @Summary Get user by username
// @Description Retrieve a user by their username
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "Email"
// @Success 200 {object} entities.User
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /api/v1/users/{email} [get]
func (controller *UserController) GetUserByUsername(ctx *gin.Context) {
	email := ctx.Param("email")
	fmt.Println(email)
	user, err := controller.userService.GetByUsername(ctx, email)
	fmt.Println(user, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
		})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, Response{
			Code:   http.StatusNotFound,
			Status: "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code:   http.StatusOK,
		Status: "OK",
	})
}
