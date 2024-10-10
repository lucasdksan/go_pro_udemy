package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func LoadDataBase(url string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), url)

	if err != nil {
		return nil, err
	}

	slog.Info("connection to the database made successfully")

	return dbpool, nil
}
