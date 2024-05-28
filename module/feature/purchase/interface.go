package purchase

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/dto"
	"github.com/gofiber/fiber/v2"
)

type PurchaseRepositoryInterface interface {
	FindNearbyMerchants(ctx context.Context, payload *dto.ReqNearbyMerchants) ([]entities.NearbyMerchant, error)
}

type PurchaseServiceInterface interface {
	GetNearbyMerchants(ctx context.Context, payload *dto.ReqNearbyMerchants) ([]entities.NearbyMerchant, error)
}

type PurchaseControllerInterface interface {
	GetNearbyMerchants(ctx *fiber.Ctx) error
}
