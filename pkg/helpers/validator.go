package helpers

import (
	"context"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}
