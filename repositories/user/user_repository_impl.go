package user_repository

import (
	"user-service/models/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) CreateUser(ctx *fiber.Ctx, tx *gorm.DB, userInput *user.User) (string, error) {
	err := tx.Create(userInput).Error

	return userInput.Email, err
}
