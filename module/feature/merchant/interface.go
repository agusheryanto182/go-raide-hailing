package merchant

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant/dto"
	"github.com/gofiber/fiber/v2"
)

type MerchantRepositoryInterface interface {
	Create(ctx context.Context, merchant *entities.Merchant) (string, error)
	FindMerchantByFilters(ctx context.Context, payload *dto.ReqGetMerchantByFilters) (*dto.Meta, []*dto.ResGetMerchant, error)
	CreateMerchantItems(ctx context.Context, items *entities.MerchantItem) (string, error)
	FindMerchantItemsByFilters(ctx context.Context, payload *dto.ReqGetMerchantItemsByFilters) (*dto.Meta, []*entities.MerchantItem, error)
}

type MerchantServiceInterface interface {
	CreateMerchant(ctx context.Context, payload *dto.ReqCreateMerchant) (string, error)
	GetMerchantByFilters(ctx context.Context, payload *dto.ReqGetMerchantByFilters) (*dto.Meta, []*dto.ResGetMerchant, error)
	CreateMerchantItems(ctx context.Context, items *dto.ReqCreateMerchantItem) (string, error)
	GetMerchantItemsByFilters(ctx context.Context, payload *dto.ReqGetMerchantItemsByFilters) (*dto.Meta, []*entities.MerchantItem, error)
}

type MerchantControllerInterface interface {
	CreateMerchant(ctx *fiber.Ctx) error
	GetMerchantByFilters(ctx *fiber.Ctx) error
	CreateMerchantItems(ctx *fiber.Ctx) error
	GetMerchantItemsByFilters(ctx *fiber.Ctx) error
}
