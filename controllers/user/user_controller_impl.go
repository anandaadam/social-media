package user_controller

import (
	"user-service/helpers"
	"user-service/models/user"
	user_service "user-service/services/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService user_service.UserService
}

func NewUserController(userService user_service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (userController *UserControllerImpl) CreateUser(ctx *fiber.Ctx) error {
	reqBody := &user.CreateUserRequest{}

	if err := ctx.BodyParser(reqBody); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(reqBody)

	if err != nil {
		errorMessages := errorMessageValidation(err.(validator.ValidationErrors))

		return helpers.FailResponse(ctx, "Error validation", errorMessages, 422)
	}

	user, err := userController.UserService.CreateUser(ctx, reqBody)

	if err != nil {
		return helpers.FailResponse(ctx, "Failed to create user", err.Error(), 500)
	}

	return helpers.SuccessResponse(ctx, "Success to signup", user, 201)
}

func errorMessageValidation(errors validator.ValidationErrors) []string {
	messages := make([]string, 0)

	for _, err := range errors {
		switch err.Field() {
		case "Email":
			switch err.Tag() {
			case "required":
				messages = append(messages, "Email wajib diisi")
			case "email":
				messages = append(messages, "Email format salah")
			}

		case "Password":
			switch err.Tag() {
			case "required":
				messages = append(messages, "Password wajib diisi")
			case "min":
				messages = append(messages, "Password minimal 8 karakter")
			}
		}
	}

	return messages
}
