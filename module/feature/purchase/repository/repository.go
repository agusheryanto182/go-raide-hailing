package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/agusheryanto182/go-raide-hailing/module/entities"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase"
	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
	"github.com/jmoiron/sqlx"
)

type purchaseRepository struct {
	db         *sqlx.DB
	statements statements
}

func (p *purchaseRepository) GetOrders(ctx context.Context, payload *dto.ReqGetOrders) ([]*dto.ResGetOrders, error) {
	query := `
        SELECT
            order_id,
            estimate_id,
            user_id
        FROM
            orders
        WHERE
            user_id = $1
    `
	res := []*dto.ResGetOrders{}
	tempOrders := []*dto.TempOrders{}

	rows, err := p.db.QueryxContext(ctx, query, payload.UserId)
	if err != nil {
		return nil, customErr.NewInternalServerError("failed to get orders : " + err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		r := &dto.TempOrders{}
		if err := rows.StructScan(&r); err != nil {
			return nil, customErr.NewInternalServerError("failed to get orders : " + err.Error())
		}
		tempOrders = append(tempOrders, r)
	}

	for i := range tempOrders {
		rows, err := p.statements.GetDetailOrders.QueryxContext(ctx, tempOrders[i].OrderId)
		if err != nil {
			return nil, customErr.NewInternalServerError("failed to get merchants: " + err.Error())
		}
		defer rows.Close()

		var (
			orderIitemsJSONStr string
			merchantJSONStr    string
			itemsJSONStr       string
		)

		for rows.Next() {
			if err := rows.Scan(
				&orderIitemsJSONStr,
				&merchantJSONStr,
				&itemsJSONStr,
			); err != nil {
				return nil, customErr.NewInternalServerError("failed to scan merchants: " + err.Error())
			}

			orderItems := []*dto.TempOrderItems{}
			if err := json.Unmarshal([]byte(orderIitemsJSONStr), &orderItems); err != nil {
				return nil, customErr.NewInternalServerError("failed to unmarshal order items: " + err.Error())
			}
			merchants := []*dto.TempMerchant{}
			if err := json.Unmarshal([]byte(merchantJSONStr), &merchants); err != nil {
				return nil, customErr.NewInternalServerError("failed to unmarshal merchants: " + err.Error())
			}
			items := []*dto.TempItems{}
			if err := json.Unmarshal([]byte(itemsJSONStr), &items); err != nil {
				return nil, customErr.NewInternalServerError("failed to unmarshal items: " + err.Error())
			}

			if len(orderItems) != len(merchants) || len(orderItems) != len(items) {
				return nil, customErr.NewInternalServerError("length of order items, merchants, and items do not match")
			}

			merchantMap := make(map[string]*dto.Orders)
			itemMap := make(map[string]bool) // declare itemMap here
			for _, orderItem := range orderItems {
				merchantID := orderItem.MerchantId
				if _, ok := merchantMap[merchantID]; !ok {
					merchant := findMerchantByID(merchants, merchantID)
					if merchant == nil {
						return nil, customErr.NewInternalServerError("failed to find merchant with ID " + merchantID)
					}
					merchantMap[merchantID] = &dto.Orders{
						Merchant: &entities.Merchant{
							ID:               merchant.ID,
							Name:             merchant.Name,
							MerchantCategory: merchant.MerchantCategory,
							ImageUrl:         merchant.ImageUrl,
							Location: entities.Location{
								Latitude:  merchant.LocationLat,
								Longitude: merchant.LocationLong,
							},
							CreatedAt: merchant.CreatedAt,
						},
						Items: []*dto.Items{},
					}
				}
				itemID := orderItem.MerchantItemId
				if _, ok := itemMap[itemID]; !ok {
					for _, item := range items {
						if item.MerchantId == orderItem.MerchantId && item.ID == itemID {
							merchantMap[merchantID].Items = append(merchantMap[merchantID].Items, &dto.Items{
								ID:              item.ID,
								Name:            item.Name,
								ProductCategory: item.ProductCategory,
								Price:           item.Price,
								Quantity:        orderItem.Quantity,
								ImageUrl:        item.ImageUrl,
								CreatedAt:       item.CreatedAt,
							})
							itemMap[itemID] = true
							break
						}
					}
				}
			}

			resItem := &dto.ResGetOrders{
				OrderId: tempOrders[i].OrderId,
				Orders:  []*dto.Orders{},
			}
			for _, orders := range merchantMap {
				resItem.Orders = append(resItem.Orders, orders)
			}

			res = append(res, resItem)
		}
	}

	return res, nil
}

func findMerchantByID(merchants []*dto.TempMerchant, id string) *dto.TempMerchant {
	for _, merchant := range merchants {
		if merchant.ID == id {
			return merchant
		}
	}
	return nil
}

// PostOrders implements purchase.PurchaseRepositoryInterface.
func (p *purchaseRepository) PostOrders(ctx context.Context, payload *dto.ReqPostOrders) (string, error) {
	var id string
	if err := p.statements.CreateOrders.QueryRowxContext(ctx, payload).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", customErr.NewNotFoundError("calculatedEstimatedId is not found : " + err.Error())
		}
		return "", customErr.NewInternalServerError("failed to post orders : " + err.Error())
	}

	payload.OrderId = id
	rows := p.statements.PostOrders.MustExecContext(ctx, payload)
	_, err := rows.RowsAffected()
	if err != nil {
		return "", customErr.NewInternalServerError("rows affected : " + err.Error())
	}

	return id, nil
}

// SaveOrderItems implements purchase.PurchaseRepositoryInterface.
func (p *purchaseRepository) SaveOrderItems(ctx context.Context, payload *dto.ReqPostEstimate) error {
	for _, order := range payload.Orders {
		currMerchant := order.MerchantId
		estimateId := payload.EstimateId

		for _, item := range order.Items {
			row := p.statements.SaveOrderItemsWithoutOrderId.MustExecContext(ctx, estimateId, currMerchant, item.ItemId, item.Quantity)
			_, err := row.RowsAffected()
			if err != nil {
				return customErr.NewInternalServerError("rows affected : " + err.Error())
			}
		}
	}

	return nil
}

// CreateEstimate implements purchase.PurchaseRepositoryInterface.
func (p *purchaseRepository) CreateEstimate(ctx context.Context, estimate *entities.Estimate) (string, error) {
	var id string
	if err := p.statements.CreateEstimate.QueryRowxContext(ctx, estimate).Scan(&id); err != nil {
		return "", customErr.NewInternalServerError("failed to create estimate : " + err.Error())
	}
	return id, nil
}

// PostEstimate implements purchase.PurchaseRepositoryInterface.
func (p *purchaseRepository) PostEstimate(ctx context.Context, merchantUUIDParams []string, itemParams *dto.ItemParams) ([]*dto.ResEstimateMerchant, []*dto.ResEstimateItem, error) {
	merchants := []*dto.ResEstimateMerchant{}
	items := []*dto.ResEstimateItem{}

	for _, uuid := range merchantUUIDParams {
		rows, err := p.statements.GetMerchants.QueryxContext(ctx, uuid)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil, customErr.NewNotFoundError("merchant not found")
			}
			return nil, nil, customErr.NewInternalServerError("failed to get merchants : " + err.Error())
		}

		defer rows.Close()
		for rows.Next() {
			merchant := &dto.ResEstimateMerchant{}
			if err := rows.Scan(
				&merchant.MerchantId,
				&merchant.LocationLat,
				&merchant.LocationLong,
			); err != nil {
				return nil, nil, customErr.NewInternalServerError("failed to scan merchants : " + err.Error())
			}

			merchants = append(merchants, merchant)
		}
	}

	for i := range itemParams.ItemIDParams {
		rows, err := p.statements.GetItems.QueryxContext(ctx, itemParams.ItemIDParams[i], itemParams.Quantity[i], itemParams.MerchantId[i])
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil, customErr.NewNotFoundError("item not found")
			}
			return nil, nil, customErr.NewInternalServerError("failed to get items : " + err.Error())
		}

		defer rows.Close()
		for rows.Next() {
			item := &dto.ResEstimateItem{}
			if err := rows.Scan(
				&item.ItemId,
				&item.MerchantId,
				&item.TotalPrice,
			); err != nil {
				return nil, nil, customErr.NewInternalServerError("failed to scan items : " + err.Error())
			}

			items = append(items, item)
		}
	}

	if len(merchants) != len(merchantUUIDParams) || len(items) != len(itemParams.ItemIDParams) {
		return nil, nil, customErr.NewNotFoundError("merchant or item not found")
	}

	return merchants, items, nil
}

// FindNearbyMerchants implements purchase.PurchaseRepositoryInterface.
func (p *purchaseRepository) FindNearbyMerchants(ctx context.Context, payload *dto.ReqNearbyMerchants) ([]*entities.NearbyMerchant, error) {
	query := `
	SELECT
    m.merchant_id,
    m.name,
    m.merchant_category,
    m.image_url,
    m.location_lat,
    m.location_long,
    m.created_at,
	COALESCE(json_agg(i) FILTER (WHERE i.item_id IS NOT NULL), '[]') AS items
	FROM merchants m
	LEFT JOIN merchant_items i ON m.merchant_id = i.merchant_id
	WHERE 1 = 1
	`

	if payload.MerchantId != "" {
		query += ` AND m.merchant_id = :merchant_id `
	}

	if payload.Name != "" {
		query += ` AND m.name ILIKE '%' || :name || '%' `
	}

	if payload.MerchantCategory != "" {
		query += ` AND m.merchant_category = :merchant_category `
	}

	query += ` 
        GROUP BY m.merchant_id
    `

	merchants := []*entities.NearbyMerchant{}
	rows, err := p.db.NamedQueryContext(ctx, query, payload)

	if err != nil {
		return nil, customErr.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var (
			itemsJSONStr string
			items        = []*dto.TempItems{}
		)

		merchant := &entities.NearbyMerchant{}

		if err := rows.Scan(
			&merchant.Merchant.ID,
			&merchant.Merchant.Name,
			&merchant.Merchant.MerchantCategory,
			&merchant.Merchant.ImageUrl,
			&merchant.Merchant.Location.Latitude,
			&merchant.Merchant.Location.Longitude,
			&merchant.Merchant.CreatedAt,
			&itemsJSONStr,
		); err != nil {
			return nil, customErr.NewInternalServerError(err.Error())
		}
		if err := json.Unmarshal([]byte(itemsJSONStr), &items); err != nil {
			return nil, customErr.NewInternalServerError(err.Error())
		}

		if len(items) == 0 {
			merchant.Items = []entities.Item{}
		}

		for i := 0; i < len(items); i++ {
			if items[i] == nil || len(items) == 0 {
				merchant.Items = []entities.Item{}
				continue
			}

			merchant.Items = append(merchant.Items, entities.Item{
				ID:              items[i].ID,
				Name:            items[i].Name,
				ProductCategory: items[i].ProductCategory,
				Price:           items[i].Price,
				ImageUrl:        items[i].ImageUrl,
				CreatedAt:       items[i].CreatedAt,
			})
		}

		merchants = append(merchants, merchant)
	}
	return merchants, nil
}

func NewPurchaseRepository(db *sqlx.DB) purchase.PurchaseRepositoryInterface {
	return &purchaseRepository{
		db:         db,
		statements: prepareStatements(),
	}
}
