package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func LoadDataBase(url string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(url)

	if err != nil {
		slog.Info("Unable to parse connection URL")
		return nil, err
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		slog.Info("Unable to create pool")
		return nil, err
	}

	if err = dbpool.Ping(context.Background()); err != nil {
		slog.Info("Unable to ping database")
		return nil, err
	}

	slog.Info("connection to the database made successfully")

	return dbpool, nil
}
