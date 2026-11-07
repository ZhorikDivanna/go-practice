package simple_connection

import (
	"context"
	"fmt"
)

func CheckConnection() {
	ctx := context.Background()

	conn, err := CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подключение к базе данных произошло успешно!")
}
