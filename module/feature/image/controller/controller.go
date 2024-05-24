package controller

import (
	"github.com/agusheryanto182/go-raide-hailing/module/feature/image"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/image/service"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/agusheryanto182/go-raide-hailing/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

type ImageHandler struct {
	imageService service.ImageService
}

func NewImageController(imageService service.ImageService) image.ImageControllerInterface {
	return &ImageHandler{
		imageService: imageService,
	}
}

func (c *ImageHandler) UploadImage(ctx *fiber.Ctx) error {
	currentUser := ctx.Locals("CurrentUser").(*jwt.JWTPayload)
	if currentUser.Role != "admin" {
		return customErr.NewUnauthorizedError("Access denied: invalid token")
	}

	fileHeader, err := ctx.FormFile("file")
	if fileHeader == nil {
		return customErr.NewBadRequestError("file should not be empty")
	}
	if err != nil {
		return customErr.NewInternalServerError("failed to retrieve file")
	}

	url, err := c.imageService.UploadImage(fileHeader)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "image uploaded successfully",
		"data": fiber.Map{
			"imageUrl": url,
		},
	})

}
