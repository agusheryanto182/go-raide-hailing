package config

import (
	"fmt"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

type Global struct {
	ServerHost string
	ServerPort int

	DbString       string
	DbMaxOpenConns int
	// DbMaxIdleConns    int
	DbMaxConnIdleTime time.Duration
	// DbMaxConnLifetime time.Duration

	JwtSecretKey   string
	BcryptSaltCost int

	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsS3BucketName    string
	AwsRegion          string

	LogLevel string
}

func LoadConfig() (cfg Global, err error) {
	conf := viper.New()
	conf.SetConfigFile(".env")
	conf.SetConfigType("env")
	conf.AutomaticEnv()

	err = conf.ReadInConfig()
	if err != nil {
		return
	}

	conf.SetDefault("HOST", "0.0.0.0")
	conf.SetDefault("PORT", 8080)

	cfg.ServerHost = conf.GetString("HOST")
	cfg.ServerPort = conf.GetInt("PORT")

	cfg.DbString = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		conf.GetString("DB_USERNAME"),
		url.QueryEscape(conf.GetString("DB_PASSWORD")),
		conf.GetString("DB_HOST"),
		conf.GetInt("DB_PORT"),
		conf.GetString("DB_NAME"),
		conf.GetString("DB_PARAMS"),
	)

	conf.SetDefault("DB_MAX_OPEN_CONNS", 100)
	cfg.DbMaxOpenConns = conf.GetInt("DB_MAX_OPEN_CONNS")

	// conf.SetDefault("DB_MAX_IDLE_CONNS", 25)
	// cfg.DbMaxIdleConns = conf.GetInt("DB_MAX_IDLE_CONNS")

	conf.SetDefault("DB_MAX_CONN_IDLE_TIME_S", 5)
	cfg.DbMaxConnIdleTime = time.Duration(
		conf.GetInt64("DB_MAX_CONN_IDLE_TIME_S"),
	) * time.Second

	// conf.SetDefault("DB_MAX_CONN_LIFETIME_MS", 0)
	// cfg.DbMaxConnLifetime = time.Duration(
	// 	conf.GetInt64("DB_MAX_CONN_LIFETIME_MS"),
	// ) * time.Millisecond

	cfg.JwtSecretKey = conf.GetString("JWT_SECRET")
	cfg.BcryptSaltCost = conf.GetInt("BCRYPT_SALT")

	cfg.AwsAccessKeyID = conf.GetString("AWS_ACCESS_KEY_ID")
	cfg.AwsSecretAccessKey = conf.GetString("AWS_SECRET_ACCESS_KEY")
	cfg.AwsS3BucketName = conf.GetString("AWS_S3_BUCKET_NAME")
	cfg.AwsRegion = conf.GetString("AWS_REGION")

	cfg.LogLevel = conf.GetString("LOG_LEVEL")
	return
}
