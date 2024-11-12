package api

import (
	user_controller "user-service/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func Router(userController user_controller.UserController) *fiber.App {
	app := fiber.New()
	api := app.Group("/api")

	signup := api.Group("/signup")
	signup.Post("/user", userController.CreateUser)

	return app
}
