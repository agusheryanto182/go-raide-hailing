package repository

import (
	statementutil "github.com/agusheryanto182/go-raide-hailing/utils/statementUtils"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	CreateMerchant *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		CreateMerchant: statementutil.MustPrepareNamed(`
		INSERT INTO merchants(
			name,
			merchant_category,
			image_url,
			location
		) VALUES (
			:name,
			:merchant_category,
			:image_url,
			:location
		) RETURNING id
		`),
	}
}
