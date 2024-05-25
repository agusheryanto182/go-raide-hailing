package merchant

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant/dto"
	"github.com/gofiber/fiber/v2"
)

type MerchantRepositoryInterface interface {
	Create(ctx context.Context, merchant *entities.Merchant) (string, error)
}

type MerchantServiceInterface interface {
	CreateMerchant(ctx context.Context, payload *dto.ReqCreateMerchant) (string, error)
}

type MerchantControllerInterface interface {
	CreateMerchant(ctx *fiber.Ctx) error
}
