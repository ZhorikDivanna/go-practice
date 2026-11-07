package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertBook(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		INSERT INTO books (
			name,
			author,
			year_of_publication,
			completed,
			on_the_shelf_at
		)
		VALUES
			('Книга 1', 'Автор 1', 2001, FALSE, NOW()),
			('Книга 2', 'Автор 2', 2002, FALSE, NOW()),
			('Книга 3', 'Автор 3', 2003, FALSE, NOW()),
			('Книга 4', 'Автор 4', 2004, FALSE, NOW()),
			('Книга 5', 'Автор 5', 2005, FALSE, NOW()),
			('Книга 6', 'Автор 6', 2006, FALSE, NOW()),
			('Книга 7', 'Автор 7', 2007, FALSE, NOW()),
			('Книга 8', 'Автор 8', 2008, FALSE, NOW()),
			('Книга 9', 'Автор 9', 2009, FALSE, NOW()),
			('Книга 10', 'Автор 10', 2010, FALSE, NOW()),
			('Книга 11', 'Автор 11', 2011, FALSE, NOW()),
			('Книга 12', 'Автор 12', 2012, FALSE, NOW()),
			('Книга 13', 'Автор 13', 2013, FALSE, NOW()),
			('Книга 14', 'Автор 14', 2014, FALSE, NOW()),
			('Книга 15', 'Автор 15', 2015, FALSE, NOW()),
			('Книга 16', 'Автор 16', 2016, FALSE, NOW()),
			('Книга 17', 'Автор 17', 2017, FALSE, NOW()),
			('Книга 18', 'Автор 18', 2018, FALSE, NOW()),
			('Книга 19', 'Автор 19', 2019, FALSE, NOW()),
			('Книга 20', 'Автор 20', 2020, FALSE, NOW()),
			('Книга 21', 'Автор 21', 2021, FALSE, NOW()),
			('Книга 22', 'Автор 22', 2022, FALSE, NOW()),
			('Книга 23', 'Автор 23', 2023, FALSE, NOW()),
			('Книга 24', 'Автор 24', 2024, FALSE, NOW());
	`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}