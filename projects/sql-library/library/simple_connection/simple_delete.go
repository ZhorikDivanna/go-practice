package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteBooks(ctx context.Context, conn *pgx.Conn, ids []int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE id = ANY($1);`

	_, err := conn.Exec(ctx, sqlQuery, ids)

	return err
}
