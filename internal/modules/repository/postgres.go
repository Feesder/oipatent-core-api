package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Username string
	Password string
	Port     string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.Port, cfg.DBName, cfg.SSLMode,
	))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
