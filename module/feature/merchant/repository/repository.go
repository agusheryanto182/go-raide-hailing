package repository

import (
	"context"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
)

type merchantsRepository struct {
	statements statements
}

// Create implements merchant.MerchantRepositoryInterface.
func (m *merchantsRepository) Create(ctx context.Context, merchant *entities.Merchant) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	var id string
	err := m.statements.CreateMerchant.QueryRowxContext(ctx, merchant).Scan(&id)
	if err != nil {
		return "", customErr.NewInternalServerError("failed to create merchant : " + err.Error())
	}

	return id, nil
}

func NewMerchantRepository() merchant.MerchantRepositoryInterface {
	return &merchantsRepository{
		statements: prepareStatements(),
	}
}
