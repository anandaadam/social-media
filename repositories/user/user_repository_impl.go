package user_repository

import (
	"time"
	"user-service/models/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) CreateUser(ctx *fiber.Ctx, tx *gorm.DB, useReq *user.CreateUserRequest) (string, error) {
	insertedUser := &user.User{
		Email:     useReq.Email,
		Password:  useReq.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := tx.Create(insertedUser).Error

	return insertedUser.Email, err
}
