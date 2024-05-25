package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BindValidate(ctx *fiber.Ctx, req interface{}, validator *validator.Validate) error {
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	if err := validator.Struct(req); err != nil {
		return err
	}

	return nil
}
