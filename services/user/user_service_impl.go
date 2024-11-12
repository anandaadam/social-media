package user_service

import (
	"user-service/helpers"
	"user-service/models/user"
	user_repository "user-service/repositories/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserServiceImpl adalah struct yang mengimplementasikan interface UserService.
type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *gorm.DB
	// Tambahkan dependensi yang diperlukan di sini, seperti repository atau logger.
}

// NewUserService membuat instance baru dari UserServiceImpl.
func NewUserService(userRepository user_repository.UserRepository, db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

// CreateUser adalah implementasi dari method CreateUser di UserService.
func (userService *UserServiceImpl) CreateUser(ctx *fiber.Ctx, userRequest *user.CreateUserRequest) (string, error) {
	var email string
	password, err := helpers.HashPassword(userRequest.Password)
	userRequest.Password = password

	err = userService.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		email, err = userService.UserRepository.CreateUser(ctx, tx, userRequest)

		return err
	})

	return email, err
}
