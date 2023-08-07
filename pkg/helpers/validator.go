package helpers

import (
	"context"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Custom error messages for validation tags
var validationErrors = map[string]string{
	"required": "The {field} field is required.",
	"email":    "The {field} field must be a valid email address.",
	"lte":      "The {field} field must be less than or equal to {param}.",
	"gte":      "The {field} field must be greater than or equal to {param}.",
}

func init() {
	validate = validator.New()
}

// Validate validates the input structure and returns a descriptive error message
func Validate(ctx context.Context, s interface{}) error {
	if err := validate.StructCtx(ctx, s); err != nil {
		var validationErrs []string
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			param := err.Param()
			tag := err.Tag()

			// Get the custom error message if available, otherwise use the default error message
			msg := validationErrors[tag]
			if msg == "" {
				msg = "The {field} field is invalid."
			}

			// Replace placeholders with actual values
			msg = strings.ReplaceAll(msg, "{field}", field)
			msg = strings.ReplaceAll(msg, "{param}", param)

			validationErrs = append(validationErrs, msg)
		}

		return errors.New(strings.Join(validationErrs, " "))
	}

	return nil
}
