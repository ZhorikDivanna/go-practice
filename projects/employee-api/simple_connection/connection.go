package simple_connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnection(ctx context.Context) (*pgxpool.Pool, error) {

	return pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
}
