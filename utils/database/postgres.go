package database

import (
	"fmt"

	"github.com/agusheryanto182/go-raide-hailing/config"
	"github.com/agusheryanto182/go-raide-hailing/utils/logging"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitDatabase(cfg config.Global) (*sqlx.DB, error) {

	db, err := sqlx.Connect("pgx", cfg.DbString)

	db.SetMaxOpenConns(cfg.DbMaxOpenConns)
	db.SetConnMaxIdleTime(cfg.DbMaxConnIdleTime)
	// db.SetMaxIdleConns(cfg.DbMaxIdleConns)
	// db.SetConnMaxLifetime(cfg.DbMaxConnLifetime)

	logging.GetLogger("db").Debug(fmt.Sprintf("Connected to database: %s", cfg.DbString))
	return db, err
}
