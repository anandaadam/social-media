package user

import (
	"regexp"

	"github.com/go-playground/validator"
)

type CreateUserRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Minimum 8 characters, at least one uppercase letter, one number, and one special character
	regex := `^(?=.*[A-Z])(?=.*\d)(?=.*[^a-zA-Z0-9]).{8,}$`
	match, _ := regexp.MatchString(regex, password)

	return match
}
