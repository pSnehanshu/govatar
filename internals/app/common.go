package app

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type FieldError struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
	Message     string
}

type ErrorResponse struct {
	Errors []FieldError
}

func sendError(status int, c *fiber.Ctx, errors *[]FieldError) error {
	return c.Status(status).JSON(ErrorResponse{
		Errors: *errors,
	})
}

func validateReq(body interface{}, c *fiber.Ctx) []FieldError {
	validationErrors := []FieldError{}

	if err := c.BodyParser(body); err != nil {
		fieldErr := FieldError{
			Error:       true,
			FailedField: "[all]",
			Message:     fmt.Sprintf("%v", err),
		}

		validationErrors = append(validationErrors, fieldErr)
	}

	if errs := validate.Struct(body); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			elem := FieldError{
				FailedField: err.Field(), // Export struct field name
				Tag:         err.Tag(),   // Export struct tag
				Value:       err.Value(), // Export field value
				Error:       true,
				Message:     err.Error(),
			}

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
