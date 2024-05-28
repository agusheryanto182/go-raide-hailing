package repository

import (
	statementutil "github.com/agusheryanto182/go-raide-hailing/utils/statementUtils"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	CreateMerchant      *sqlx.NamedStmt
	GetCountMerchant    *sqlx.Stmt
	CreateMerchantItems *sqlx.NamedStmt
	CheckMerchant       *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		CreateMerchant: statementutil.MustPrepareNamed(`
		INSERT INTO merchants(
			name,
			merchant_category,
			image_url,
			location_lat,
			location_long
		) VALUES (
			:name,
			:merchant_category,
			:image_url,
			:location_lat,
			:location_long
		) RETURNING merchant_id
		`),

		GetCountMerchant: statementutil.MustPrepare(`SELECT COUNT(*) FROM merchants`),

		CreateMerchantItems: statementutil.MustPrepareNamed(`
		WITH check_merchant AS (
			SELECT EXISTS(SELECT 1 FROM merchants WHERE merchant_id = :merchant_id) AS merchant_exists
		)
		INSERT INTO merchant_items (
			merchant_id,
			name,
			product_category,
			price,
			image_url
		)
		SELECT
			:merchant_id,
			:name,
			:product_category,
			:price,
			:image_url
		FROM check_merchant
		WHERE check_merchant.merchant_exists = true
		RETURNING item_id		
		`),

		CheckMerchant: statementutil.MustPrepare(`SELECT EXISTS(SELECT 1 FROM merchants WHERE merchant_id = $1) AS merchant_exists`),
	}
}
