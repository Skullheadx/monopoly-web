package main

import (
	"fmt"
	"io"
	"log"
	"monopoly-web/game"
	"monopoly-web/types"
	"net/http"
)

func main() {
	fmt.Println("monopoly-web backend")

	game.Users[0] = types.User{UUID: "abc", Money: 100}
	fmt.Println(game.Users)

	// register routes
	http.HandleFunc("/health", healthHandler)

	// listen and serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Status: healthy\n")
}
