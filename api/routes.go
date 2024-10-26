package api

import (
	user_controller "social-media/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	app := fiber.New()
	api := app.Group("/api")

	signup := api.Group("/signup")
	signup.Post("/user", user_controller.CreateUser)

	app.Listen(":3000")
}
