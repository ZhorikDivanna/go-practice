package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, connectionString)
}