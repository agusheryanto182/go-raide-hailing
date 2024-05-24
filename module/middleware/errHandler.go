package middleware

import (
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := ctx.Response().StatusCode()
	if customError, ok := err.(customErr.CustomError); ok {
		code = customError.Status()
		return ctx.Status(code).JSON(fiber.Map{
			"message": customError.Error(),
		})
	}

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		return ctx.Status(code).JSON(fiber.Map{
			"message": e.Message,
		})
	}

	return ctx.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
}
