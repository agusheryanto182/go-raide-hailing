package request

import (
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BindValidate(ctx *fiber.Ctx, req interface{}, validator *validator.Validate) error {
	if err := ctx.BodyParser(req); err != nil {
		return customErr.NewBadRequestError("failed to parse request body : " + err.Error())
	}

	if err := validator.Struct(req); err != nil {
		return customErr.NewBadRequestError("failed to validate request body : " + err.Error())
	}

	return nil
}
