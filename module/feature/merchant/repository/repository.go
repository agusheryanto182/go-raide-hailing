package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/merchant/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/jmoiron/sqlx"
)

type merchantsRepository struct {
	db         *sqlx.DB
	statements statements
}

// CreateMerchantItems implements merchant.MerchantRepositoryInterface.
func (m *merchantsRepository) CreateMerchantItems(ctx context.Context, items *entities.MerchantItem) (string, error) {
	var id string
	if err := ctx.Err(); err != nil {
		return "", err
	}

	err := m.statements.CreateMerchantItems.QueryRowxContext(ctx, items).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", customErr.NewNotFoundError("merchant not found")
		}
		return "", customErr.NewInternalServerError("failed to create merchant items : " + err.Error())
	}

	return id, nil
}

// FindMerchantItemsByFilters implements merchant.MerchantRepositoryInterface.
func (m *merchantsRepository) FindMerchantItemsByFilters(ctx context.Context, payload *dto.ReqGetMerchantItemsByFilters) (*dto.Meta, []*entities.MerchantItem, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}

	var checkMerchant bool
	if err := m.statements.CheckMerchant.QueryRowxContext(ctx, payload.MerchantId).Scan(&checkMerchant); err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to check merchant : " + err.Error())
	}

	if !checkMerchant {
		return nil, nil, customErr.NewNotFoundError("merchant not found")
	}

	meta := &dto.Meta{
		Limit:  payload.Limit,
		Offset: payload.Offset,
		Total:  0,
	}

	merchant_items := make([]*entities.MerchantItem, 0)

	query := ` 
	SELECT * FROM merchant_items
	WHERE 1 = 1
	AND merchant_id = :merchant_id
	`

	if payload.ItemId != "" {
		query += ` AND item_id = :item_id `
	}

	if payload.Name != "" {
		query += ` AND name ILIKE '%' || :name || '%'`
	}

	if payload.ProductCategory != "" {
		query += ` AND product_category = :product_category `
	}

	query += ` ORDER BY created_at ` + payload.CreatedAt + ` LIMIT :limit OFFSET :offset`

	rows, err := m.db.NamedQueryContext(ctx, query, payload)
	if err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to find merchant_item : " + err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var tempCreatedAt *time.Time
		merchant_item := &entities.MerchantItem{}
		if err := rows.Scan(
			&merchant_item.ID,
			&merchant_item.MerchantId,
			&merchant_item.Name,
			&merchant_item.ProductCategory,
			&merchant_item.Price,
			&merchant_item.ImageUrl,
			&tempCreatedAt,
		); err != nil {
			return nil, nil, customErr.NewInternalServerError("failed to scan merchant items : " + err.Error())
		}

		merchant_item.CreatedAt = tempCreatedAt.Format("2006-01-02T15:04:05.000000000Z")
		merchant_items = append(merchant_items, merchant_item)
	}

	if err := m.statements.GetCountMerchant.GetContext(ctx, &meta.Total); err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to get total merchant : " + err.Error())
	}

	return meta, merchant_items, nil
}

// FindMerchantByFilters implements merchant.MerchantRepositoryInterface.
func (m *merchantsRepository) FindMerchantByFilters(ctx context.Context, payload *dto.ReqGetMerchantByFilters) (*dto.Meta, []*dto.ResGetMerchant, error) {

	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}

	meta := &dto.Meta{
		Limit:  payload.Limit,
		Offset: payload.Offset,
		Total:  0,
	}

	merchants := make([]*dto.ResGetMerchant, 0)

	query := ` 
	SELECT * FROM merchants
	WHERE 1 = 1
	`

	if payload.MerchantId != "" {
		query += ` AND merchant_id = :merchant_id `
	}

	if payload.Name != "" {
		query += ` AND name ILIKE '%' || :name || '%'`
	}

	if payload.MerchantCategory != "" {
		query += ` AND merchant_category = :merchant_category `
	}

	query += ` ORDER BY created_at ` + payload.CreatedAt + ` LIMIT :limit OFFSET :offset`

	rows, err := m.db.NamedQueryContext(ctx, query, payload)
	if err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to find merchants : " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var tempCreatedAt *time.Time
		merchant := &dto.ResGetMerchant{}
		if err := rows.Scan(
			&merchant.ID,
			&merchant.Name,
			&merchant.MerchantCategory,
			&merchant.ImageUrl,
			&merchant.Location.Latitude,
			&merchant.Location.Longitude,
			&tempCreatedAt,
		); err != nil {
			return nil, nil, customErr.NewInternalServerError("failed to scan merchant : " + err.Error())
		}

		merchant.CreatedAt = tempCreatedAt.Format("2006-01-02T15:04:05.000000000Z")
		merchants = append(merchants, merchant)
	}

	if err := m.statements.GetCountMerchant.GetContext(ctx, &meta.Total); err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to get total merchant : " + err.Error())
	}

	return meta, merchants, nil

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

func NewMerchantRepository(db *sqlx.DB) merchant.MerchantRepositoryInterface {
	return &merchantsRepository{
		db:         db,
		statements: prepareStatements(),
	}
}
