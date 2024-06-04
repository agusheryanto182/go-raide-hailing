package database

import (
	"fmt"
	"log"

	"github.com/agusheryanto182/go-raide-hailing/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitDatabase(cfg *config.Global) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DbName,
		cfg.Database.Params,
	)

	db, err := sqlx.Connect("pgx", dsn)

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(15)
	db.SetConnMaxLifetime(0)

	log.Println("Database connected")

	return db, err
}
