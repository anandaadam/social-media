package user_repository

import (
	"user-service/models/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx *fiber.Ctx, db *gorm.DB, userInput *user.User) (string, error)
}
