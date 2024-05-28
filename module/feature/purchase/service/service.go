package service

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/dto"
)

type purchaseService struct {
	purchaseRepository purchase.PurchaseRepositoryInterface
}

// GetNearbyMerchants implements purchase.PurchaseServiceInterface.
func (p *purchaseService) GetNearbyMerchants(ctx context.Context, payload *dto.ReqNearbyMerchants) ([]entities.NearbyMerchant, error) {
	return p.purchaseRepository.FindNearbyMerchants(ctx, payload)
}

func NewPurchaseService(purchaseRepository purchase.PurchaseRepositoryInterface) purchase.PurchaseServiceInterface {
	return &purchaseService{
		purchaseRepository: purchaseRepository,
	}
}
