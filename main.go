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

	game.Users = append(game.Users, game.User{UUID: "abc", Money: 100, CurrentSpaceID: 0, GetOutOfJailCards: 0})
	fmt.Println(game.Users)

	// register routes
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/v1/roll", rollDiceHandler)
	http.HandleFunc("POST /api/v1/turn", endTurnHandler)
	http.HandleFunc("POST /api/v1/exit-jail", exitJailHandler)

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

func exitJailHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request: Failed to parse form data", http.StatusBadRequest)
		return
	}

	if !game.ValidateCanExitJail(UUID) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"status": "forbidden", "message": "Not your turn or not in jail"}`))
	}

	method := req.PostForm.Get("method")

	switch method {
	case "buyout":
		err = game.JailBuyout()
		if err == game.ErrNotEnoughMoney {
			http.Error(w, "error: Insufficient funds", http.StatusUnprocessableEntity)
		}

	case "jail_free_card":
		err = game.JailUseCard()
		if err == game.ErrNotEnoughJailCards {
			http.Error(w, "error: Insufficient jail cards", http.StatusUnprocessableEntity)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status": "bad req", "message": "escape jail method does not exist"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success", "message": "Action processed"}`))
}
