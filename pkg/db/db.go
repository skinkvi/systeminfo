package db

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/skinkvi/systeminfo/internal/config"
)

func NewDB(config *config.Config, logger *zerolog.Logger) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"),
	)

	db, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to the database")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logger.Error().Err(err).Msg("Failed to ping the database")
		return nil, err
	}

	logger.Info().Msg("Successfully connected to the database")
	return db, nil
}
