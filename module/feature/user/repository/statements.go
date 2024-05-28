package repository

import (
	statementutil "github.com/agusheryanto182/go-raide-hailing/utils/statementUtils"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	checkUsernameAndEmail *sqlx.Stmt
	registerUser          *sqlx.Stmt
	findByUsernameAndRole *sqlx.Stmt
	checkUser             *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		checkUsernameAndEmail: statementutil.MustPrepare(
			`SELECT EXISTS(SELECT 1 FROM users WHERE username = $1) AS user_exists,
			EXISTS(SELECT 1 FROM users WHERE email = $2 AND role = $3) AS email_exists`,
		),
		registerUser: statementutil.MustPrepare(
			`INSERT INTO users (username, password, role, email)
			VALUES ($1, $2, $3, $4)
			RETURNING user_id`,
		),

		findByUsernameAndRole: statementutil.MustPrepare(
			`SELECT user_id, username, password, role, email
			FROM users
			WHERE username = $1 AND role = $2`,
		),

		checkUser: statementutil.MustPrepare(
			`SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1 AND username = $2 AND role = $3) AS user_exists`,
		),
	}
}
