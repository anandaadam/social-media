package user_service

import (
	"user-service/models/user"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	CreateUser(ctx *fiber.Ctx, userRequest *user.CreateUserRequest) (string, error)
}
