package service

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/agusheryanto182/go-raide-hailing/utils/helper"
)

type purchaseService struct {
	purchaseRepository purchase.PurchaseRepositoryInterface
}

// PostEstimate implements purchase.PurchaseServiceInterface.
func (p *purchaseService) PostEstimate(ctx context.Context, payload *dto.ReqPostEstimate) (*dto.ResPostEstimate, error) {
	var merchantUUIDParams []string
	itemParams := &dto.ItemParams{}

	for _, uuid := range payload.Orders {
		merchantUUIDParams = append(merchantUUIDParams, uuid.MerchantId)
		for _, itemID := range uuid.Items {
			itemParams.ItemIDParams = append(itemParams.ItemIDParams, itemID.ItemId)
			itemParams.MerchantId = append(itemParams.MerchantId, uuid.MerchantId)
			itemParams.Quantity = append(itemParams.Quantity, itemID.Quantity)
		}
	}

	merchants, items, err := p.purchaseRepository.PostEstimate(ctx, merchantUUIDParams, itemParams)
	if err != nil {
		return nil, err
	}

	res := &dto.ResPostEstimate{}
	res.TotalPrice, res.EstimatedDeliveryTimeInMinutes, err = helper.CalculateTotalPriceAndDeliveryTime(merchants, items, payload.UserLocation.Lat, payload.UserLocation.Long)
	if err != nil {
		return nil, customErr.NewBadRequestError("failed to calculate total price and delivery time : " + err.Error())
	}

	estimateId, err := p.purchaseRepository.CreateEstimate(ctx, &entities.Estimate{
		UserId:                payload.UserId,
		UserLat:               payload.UserLocation.Lat,
		UserLon:               payload.UserLocation.Long,
		TotalPrice:            res.TotalPrice,
		EstimatedDeliveryTime: res.EstimatedDeliveryTimeInMinutes,
	})

	if err != nil {
		return nil, err
	}

	return &dto.ResPostEstimate{
		TotalPrice:                     res.TotalPrice,
		EstimatedDeliveryTimeInMinutes: res.EstimatedDeliveryTimeInMinutes,
		CalculatedEstimateId:           estimateId,
	}, nil

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
