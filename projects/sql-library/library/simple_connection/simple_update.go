package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateBook(ctx context.Context, conn *pgx.Conn, book Book) error {
	sqlQuery := `
	UPDATE books
	SET name = $1, 
	    author = $2, 
	    review = $3, 
	    year_of_publication = $4, 
	    completed = $5, 
	    on_the_shelf_at = $6, 
	    completed_at = $7
	WHERE id = $8;`

	_, err := conn.Exec(ctx, sqlQuery,
		book.Name,
		book.Author,
		book.Review,
		book.YearOfPublication,
		book.Completed,
		book.OnTheShelfAt,
		book.CompletedAt,
		book.ID,
	)

	return err
}
