package repository

import (
	statementutil "github.com/agusheryanto182/go-raide-hailing/utils/statementUtils"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	GetMerchants                 *sqlx.Stmt
	GetItems                     *sqlx.Stmt
	CreateEstimate               *sqlx.NamedStmt
	SaveOrderItemsWithoutOrderId *sqlx.Stmt
	PostOrders                   *sqlx.NamedStmt
	CreateOrders                 *sqlx.NamedStmt
	GetDetailOrders              *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		GetMerchants: statementutil.MustPrepare(`
		SELECT
		merchant_id, location_lat, location_long
		FROM merchants 
		WHERE merchant_id = $1`,
		),

		GetItems: statementutil.MustPrepare(`
		SELECT
		item_id, merchant_id, price * $2 AS total_price
		FROM merchant_items
		WHERE item_id = $1 AND merchant_id = $3`,
		),

		CreateEstimate: statementutil.MustPrepareNamed(`
		INSERT INTO estimates
		(user_id, user_lat, user_lon, total_price, estimated_delivery_time)
		VALUES
		(:user_id, :user_lat, :user_lon, :total_price, :estimated_delivery_time)
		RETURNING estimate_id`,
		),

		SaveOrderItemsWithoutOrderId: statementutil.MustPrepare(`
		INSERT INTO order_items
		(estimate_id, merchant_id, merchant_item_id, quantity)
		VALUES
		($1, $2, $3, $4)`,
		),

		PostOrders: statementutil.MustPrepareNamed(
			`
			UPDATE order_items
			SET order_id = :order_id
			WHERE estimate_id = :estimate_id
			`,
		),

		CreateOrders: statementutil.MustPrepareNamed(`
			WITH check_estimate AS (
				SELECT EXISTS (
					SELECT 1
					FROM estimates
					WHERE estimate_id = :estimate_id
					AND user_id = :user_id
				) AS exists_estimate
			)
			INSERT INTO orders
			(estimate_id, user_id)
			SELECT :estimate_id, :user_id
			FROM check_estimate
			WHERE exists_estimate
			RETURNING order_id
			`,
		),

		GetDetailOrders: statementutil.MustPrepare(`
			SELECT 
			COALESCE(json_agg(oi) FILTER (WHERE oi.order_id IS NOT NULL), '[]') AS order_items,
			COALESCE(json_agg(m) FILTER (WHERE m.merchant_id IS NOT NULL), '[]') AS merchants,
			COALESCE(json_agg(i) FILTER (WHERE i.item_id IS NOT NULL), '[]') AS items
			FROM order_items oi
			LEFT JOIN merchants m ON oi.merchant_id = m.merchant_id
			LEFT JOIN merchant_items i ON oi.merchant_item_id = i.item_id
			WHERE oi.order_id = $1
			`,
		),
	}
}
