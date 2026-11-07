package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllBooks(ctx context.Context, conn *pgx.Conn) ([]Book, error) {
	sqlQuery := `
	SELECT id, name, author, review, year_of_publication, completed, on_the_shelf_at, completed_at 
	FROM books
	ORDER BY id
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var b Book

		err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.Author,
			&b.Review,
			&b.YearOfPublication,
			&b.Completed,
			&b.OnTheShelfAt,
			&b.CompletedAt,
		)

		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
