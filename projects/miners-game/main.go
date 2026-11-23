package main

import (
	"fmt"
	"log"
	"project-miners/game"
	"project-miners/http"
)

func main() {
	enterprise := game.NewEnterPrise()

	handlers := http.NewGameHandlers(enterprise)

	server := http.NewHTTPServer(handlers)

	fmt.Println("Игра началась")
	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}
