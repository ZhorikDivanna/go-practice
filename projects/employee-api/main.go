package main

import (
	"context"
	employee_store_crud "home-3-4-5/Employee_Store_CRUD"
	"home-3-4-5/simple_connection"
	"net/http"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	if err := employee_store_crud.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	store := &employee_store_crud.Store{Conn: conn}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /employees", store.CreateEmployeeHandler)
	mux.HandleFunc("GET /employees", store.GetEmployeeHandler)
	mux.HandleFunc("DELETE /employees", store.DeleteEmployeeHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
