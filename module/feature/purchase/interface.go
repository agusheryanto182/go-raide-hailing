package purchase

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/dto"
	"github.com/gofiber/fiber/v2"
)

type PurchaseRepositoryInterface interface {
	FindNearbyMerchants(ctx context.Context, payload *dto.ReqNearbyMerchants) ([]*entities.NearbyMerchant, error)
	PostEstimate(ctx context.Context, merchantUUIDParams []string, itemParams *dto.ItemParams) ([]*dto.ResEstimateMerchant, []*dto.ResEstimateItem, error)
	CreateEstimate(ctx context.Context, estimate *entities.Estimate) (string, error)
	SaveOrderItems(ctx context.Context, payload *dto.ReqPostEstimate) error
	PostOrders(ctx context.Context, payload *dto.ReqPostOrders) (string, error)
	GetOrders(ctx context.Context, payload *dto.ReqGetOrders) ([]*dto.ResGetOrders, error)
}

type PurchaseServiceInterface interface {
	GetNearbyMerchants(ctx context.Context, payload *dto.ReqNearbyMerchants) ([]*entities.NearbyMerchant, error)
	PostEstimate(ctx context.Context, payload *dto.ReqPostEstimate) (*dto.ResPostEstimate, error)
	PostOrders(ctx context.Context, payload *dto.ReqPostOrders) (string, error)
	GetOrders(ctx context.Context, payload *dto.ReqGetOrders) ([]*dto.ResGetOrders, error)
}

type PurchaseControllerInterface interface {
	GetNearbyMerchants(ctx *fiber.Ctx) error
	PostEstimate(ctx *fiber.Ctx) error
	PostOrders(ctx *fiber.Ctx) error
	GetOrders(ctx *fiber.Ctx) error
}
