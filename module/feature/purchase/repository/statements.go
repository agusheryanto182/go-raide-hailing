package repository

import (
	statementutil "github.com/agusheryanto182/go-raide-hailing/utils/statementUtils"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	GetMerchants   *sqlx.Stmt
	GetItems       *sqlx.Stmt
	CreateEstimate *sqlx.NamedStmt
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
	}
}
