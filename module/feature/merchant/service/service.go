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

// CreateMerchantItems implements merchant.MerchantServiceInterface.
func (m *merchantService) CreateMerchantItems(ctx context.Context, items *dto.ReqCreateMerchantItem) (string, error) {
	res, err := m.merchantRepository.CreateMerchantItems(ctx, &entities.MerchantItem{
		MerchantId:      items.MerchantId,
		Name:            items.Name,
		ProductCategory: items.ProductCategory,
		Price:           items.Price,
		ImageUrl:        items.ImageUrl,
	})

	if err != nil {
		return "", err
	}

	return res, nil
}

// GetMerchantItemsByFilters implements merchant.MerchantServiceInterface.
func (m *merchantService) GetMerchantItemsByFilters(ctx context.Context, payload *dto.ReqGetMerchantItemsByFilters) (*dto.Meta, []*entities.MerchantItem, error) {
	return m.merchantRepository.FindMerchantItemsByFilters(ctx, payload)
}

// GetMerchantByFilters implements merchant.MerchantServiceInterface.
func (m *merchantService) GetMerchantByFilters(ctx context.Context, payload *dto.ReqGetMerchantByFilters) (*dto.Meta, []*dto.ResGetMerchant, error) {
	return m.merchantRepository.FindMerchantByFilters(ctx, payload)
}

// CreateMerchant implements merchant.MerchantServiceInterface.
func (m *merchantService) CreateMerchant(ctx context.Context, payload *dto.ReqCreateMerchant) (string, error) {
	id, err := m.merchantRepository.Create(ctx, &entities.Merchant{
		Name:             payload.Name,
		MerchantCategory: payload.MerchantCategory,
		ImageUrl:         payload.ImageUrl,
		Location:         []float64{*payload.Location.Latitude, *payload.Location.Longitude},
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
