package validator

import (
	"slices"

	validatorV10 "github.com/go-playground/validator/v10"
)

func NewValidator() *validatorV10.Validate {
	validate := validatorV10.New()

	// Register custom validation functions if needed
	validate.RegisterValidation("oneof", func(fl validatorV10.FieldLevel) bool {
		// Example custom validation for "oneof" to check if the value is one of the allowed values
		allowedValues := []string{"admin", "murid"}
		value := fl.Field().String()
		return slices.Contains(allowedValues, value)
	})

	return validate
}
