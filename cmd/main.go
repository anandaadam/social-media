package main

import (
	"log"
	"user-service/api"
	user_controller "user-service/controllers/user"
	"user-service/database"
	user_repository "user-service/repositories/user"
	user_service "user-service/services/user"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := database.GetDBConnection()

	userRepository := user_repository.NewUserRepository()
	userService := user_service.NewUserService(userRepository, db)
	userController := user_controller.NewUserController(userService)

	router := api.Router(userController)

	router.Listen(":3000")
}
