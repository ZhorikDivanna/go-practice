package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	author VARCHAR(100) NOT NULL,
	review VARCHAR(1000), 
	year_of_publication INTEGER NOT NULL,
	completed BOOLEAN NOT NULL, 
	on_the_shelf_at TIMESTAMP NOT NULL, 
	completed_at TIMESTAMP
	);`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
