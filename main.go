package main

import (
	"log"

	"github.com/ArshiyaDev/go-commerce/application"
	user_controlelr "github.com/ArshiyaDev/go-commerce/modules/users/controller"
	user_repo "github.com/ArshiyaDev/go-commerce/modules/users/repos"
	user_service "github.com/ArshiyaDev/go-commerce/modules/users/services"

	_ "github.com/ArshiyaDev/go-commerce/docs"
	"github.com/ArshiyaDev/go-commerce/pkg/db"
)

func main() {

	database, err := db.Init()

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Println("error closing DB:", err)
		} else {
			log.Println("DB connection closed.")
		}
	}()

	userRepo := user_repo.NewRepo(database)
	userService := user_service.NewService(userRepo)
	userController := user_controlelr.NewUserController(*userService)

	r := application.NewApllication(userController)

	r.Run(":8080")
}
