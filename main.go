package main

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"monopoly-web/game"
	"net/http"
)

type Room struct {
	gameCtx *game.Context
}

func initRoom() Room {
	randSeed := rand.NewPCG(20, 26)
	players := []game.Player{
		game.InitPlayer(),
	}

	return Room{
		gameCtx: game.InitCtx(randSeed, players),
	}
}

func main() {
	fmt.Println("monopoly-web backend")

	room := initRoom()

	// register routes
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/v1/roll", room.rollDiceHandler)
	http.HandleFunc("POST /api/v1/turn", room.endTurnHandler)
	http.HandleFunc("POST /api/v1/exit-jail", room.exitJailHandler)

	// listen and serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Status: healthy\n")
}

const UUID = "abc" // TODO: UUID in cookie

func (r *Room) rollDiceHandler(w http.ResponseWriter, req *http.Request) {
	if r.gameCtx.ValidateCanRoll(UUID) {
		r.gameCtx.RollDice()
		r.gameCtx.ProcessMovement()
	}
}

func (r *Room) endTurnHandler(w http.ResponseWriter, req *http.Request) {
	if r.gameCtx.ValidateCanEndTurn(UUID) {
		r.gameCtx.EndTurn()
	}
}

func (r *Room) exitJailHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request: Failed to parse form data", http.StatusBadRequest)
		return
	}

	if !r.gameCtx.ValidateCanExitJail(UUID) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"status": "forbidden", "message": "Not your turn or not in jail"}`))
	}

	method := req.PostForm.Get("method")

	switch method {
	case "buyout":
		err = r.gameCtx.JailBuyout()
		if err == game.ErrNotEnoughMoney {
			http.Error(w, "error: Insufficient funds", http.StatusUnprocessableEntity)
		}

	case "jail_free_card":
		err = r.gameCtx.JailUseCard()
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
