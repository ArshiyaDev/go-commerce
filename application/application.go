package application

import (
	"github.com/ArshiyaDev/go-commerce/modules/users"
	"github.com/ArshiyaDev/go-commerce/modules/users/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApllication(userController *controller.UserController) *gin.Engine {

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api/v1")
	{
		users.RegisterRoutes(api, userController)
	}

	return router

}
