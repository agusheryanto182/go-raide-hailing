package service

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant/dto"
)

type merchantService struct {
	merchantRepository merchant.MerchantRepositoryInterface
}

// CreateMerchant implements merchant.MerchantServiceInterface.
func (m *merchantService) CreateMerchant(ctx context.Context, payload *dto.ReqCreateMerchant) (string, error) {
	id, err := m.merchantRepository.Create(ctx, &entities.Merchant{
		Name:             payload.Name,
		MerchantCategory: payload.MerchantCategory,
		ImageUrl:         payload.ImageUrl,
		Location: []float64{
			payload.Location.Latitude,
			payload.Location.Longitude,
		},
	})

	if err != nil {
		return "", err
	}

	return id, nil
}

func NewMerchantService(merchantRepository merchant.MerchantRepositoryInterface) merchant.MerchantServiceInterface {
	return &merchantService{
		merchantRepository: merchantRepository,
	}
}
