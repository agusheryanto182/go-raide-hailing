package repository

import (
	"context"
	"encoding/json"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/jmoiron/sqlx"
)

type purchaseRepository struct {
	db *sqlx.DB
}

// FindNearbyMerchants implements purchase.PurchaseRepositoryInterface.
func (p *purchaseRepository) FindNearbyMerchants(ctx context.Context, payload *dto.ReqNearbyMerchants) ([]entities.NearbyMerchant, error) {
	query := `
	SELECT
    m.merchant_id,
    m.name,
    m.merchant_category,
    m.image_url,
    m.location_lat,
    m.location_long,
    m.created_at,
	COALESCE(json_agg(i) FILTER (WHERE i.item_id IS NOT NULL), '[]') AS items,
    (
        6371 *
        acos(
            cos(radians(:user_lat)) *
            cos(radians(location_lat)) *
            cos(radians(location_long) - radians(:user_long)) +
            sin(radians(:user_lat)) *
            sin(radians(location_lat))
        )
    ) AS distance
	FROM merchants m
	LEFT JOIN merchant_items i ON m.merchant_id = i.merchant_id
	WHERE 1 = 1
	`

	if payload.MerchantId != "" {
		query += ` AND merchant_id = :merchant_id `
	}

	if payload.Name != "" {
		query += ` AND name ILIKE '%' || :name || '%' `
	}

	if payload.MerchantCategory != "" {
		query += ` AND merchant_category = :merchant_category `
	}

	query += ` 
        GROUP BY m.merchant_id
        ORDER BY distance 
        LIMIT :limit OFFSET :offset 
    `

	var merchants []entities.NearbyMerchant
	rows, err := p.db.NamedQueryContext(ctx, query, payload)

	if err != nil {
		return nil, customErr.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var (
			merchant     entities.NearbyMerchant
			itemsJSONStr string
			distance     float64
			items        []dto.TempItems
		)

		if err := rows.Scan(
			&merchant.Merchant.ID,
			&merchant.Merchant.Name,
			&merchant.Merchant.MerchantCategory,
			&merchant.Merchant.ImageUrl,
			&merchant.Merchant.Location.Latitude,
			&merchant.Merchant.Location.Longitude,
			&merchant.Merchant.CreatedAt,
			&itemsJSONStr,
			&distance,
		); err != nil {
			return nil, customErr.NewInternalServerError(err.Error())
		}
		if err := json.Unmarshal([]byte(itemsJSONStr), &items); err != nil {
			return nil, customErr.NewInternalServerError(err.Error())
		}

		for i := 0; i < len(items); i++ {
			merchant.Items = append(merchant.Items, struct {
				ID              string "json:\"itemId\""
				Name            string "json:\"name\""
				ProductCategory string "json:\"productCategory\""
				Price           int    "json:\"price\""
				ImageUrl        string "json:\"imageUrl\""
				CreatedAt       string "json:\"createdAt\""
			}(items[i]))
		}

		merchants = append(merchants, merchant)
	}
	return merchants, nil
}

func NewPurchaseRepository(db *sqlx.DB) purchase.PurchaseRepositoryInterface {
	return &purchaseRepository{
		db: db,
	}
}
