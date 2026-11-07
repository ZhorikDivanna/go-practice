package simple_connection

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp"
)

func ListPages(ctx context.Context, conn *pgx.Conn, N int) ([]Book, error) {
	page := 1
	offset := 0

	sqlQuery := `
	SELECT * FROM books
	ORDER BY id
	LIMIT $1
	OFFSET $2;
	`

	if N <= 0 {
		return nil, errors.New("Размер страницы должен быть больше 0")
	}

	var allBooks []Book

	for {
		rows, err := conn.Query(ctx, sqlQuery, N, offset)
		if err != nil {
			return nil, err
		}

		var pageBooks []Book

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
				rows.Close()
				return nil, err
			}

			pageBooks = append(pageBooks, b)
			allBooks = append(allBooks, b)

		}

		if err := rows.Err(); err != nil {
			rows.Close()
			return nil, err
		}
		rows.Close()

		if len(pageBooks) == 0 {
			break
		}

		pp.Println("Страница", page, ": ", pageBooks)

		offset += N
		page++
	}

	return allBooks, nil

}
