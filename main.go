package main

import (
	"fmt"
	"io"
	"log"
	"monopoly-web/game"
	"net/http"
)

func main() {
	fmt.Println("monopoly-web backend")

	game.Users = append(game.Users, game.User{UUID: "abc", Money: 100, CurrentSpaceID: 0})
	fmt.Println(game.Users)

	// register routes
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/v1/roll", rollDiceHandler)
	http.HandleFunc("POST /api/v1/turn", endTurnHandler)

	// listen and serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Status: healthy\n")
}

const UUID = "abc" // TODO: UUID in cookie

func rollDiceHandler(w http.ResponseWriter, req *http.Request) {
	if game.ValidateCanRoll(UUID) {
		game.RollDice()
		game.ProcessMovement()
	}
}

func endTurnHandler(w http.ResponseWriter, req *http.Request) {
	if game.ValidateCanEndTurn(UUID) {
		game.EndTurn()
	}
}
