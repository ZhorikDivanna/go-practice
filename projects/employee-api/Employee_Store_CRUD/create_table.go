package employee_store_crud

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTable(ctx context.Context, conn *pgxpool.Pool) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS employees (
	id SERIAL PRIMARY KEY,
	full_name VARCHAR(200) NOT NULL,
	position VARCHAR(300) NOT NULL
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
