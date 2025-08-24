package users

import (
	"github.com/ArshiyaDev/go-commerce/modules/users/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, controller *controller.UserController) {

	users := rg.Group("/users")
	{
		users.POST("", controller.Create)
		users.GET("", controller.GetUserByUsername)
	}

}
