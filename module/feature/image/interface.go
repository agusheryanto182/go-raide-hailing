package image

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type ImageServicInterface interface {
	UploadImage(*multipart.FileHeader) (string, error)
}

type ImageControllerInterface interface {
	UploadImage(ctx *fiber.Ctx) error
}
