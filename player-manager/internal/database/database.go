package database

import (
	"database/sql"
	"fmt"
	"player-manager/internal/config"

	_ "github.com/lib/pq"
)

func EstablishConnection(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port,
	)

	return sql.Open("postgres", dsn)
}
