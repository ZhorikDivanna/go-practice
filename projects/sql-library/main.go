package main

import (
	"Home_SQL/library/simple_connection"
	"context"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	if err := simple_connection.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	books, err := simple_connection.GetAllBooks(ctx, conn)
	if err != nil {
		panic(err)
	}

	if len(books) == 0 {
		if err := simple_connection.InsertBook(ctx, conn); err != nil {
			panic(err)
		}
	}

	if _, err := simple_connection.ListPages(ctx, conn, 4); err != nil {
		panic(err)
	}
}