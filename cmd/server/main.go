package main

import (
	_ "github.com/ArshiyaDev/go-commerce/cmd/server/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary      Ping example
// @Description  Responds with pong message
// @Tags         example
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /ping [get]
func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 👈 This serves Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
