package http

import (
	customValidation "go-todo-list/internal/validator"

	"github.com/go-playground/validator/v10"
)

// CustomValidator is struct used to create request validator
type CustomValidator struct {
	Validator *validator.Validate
}

// NewValidator is function to create custom validator struct
func NewValidator() *CustomValidator {
	validate := validator.New()

	// Register validation functions
	validate.RegisterValidation("AfterToday", customValidation.AfterTodayValidator)
	validate.RegisterValidation("IsDdMmYyyy", customValidation.IsDdMmYyyyValidator)

	return &CustomValidator{
		Validator: validate,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
