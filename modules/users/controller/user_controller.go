package controller

import (
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

func (controller *UserController) Create(ctx *gin.Context) {

	createUserReq := entities.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&createUserReq)

	if err != nil {
		panic(err)
	}
	controller.userService.CreateUser(
		, Name string, LastName string, Email string, Tel string)

	webResponse := Response{Code: http.StatusOK, Status: "Ok"}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
