package game

import (
	"errors"
)

func (ctx *Context) RemovePlayerFromJail(playerID PlayerID) {
	for i, iJV := range ctx.Visitors.InJail {
		if playerID == iJV.visitorID {
			ctx.Visitors.InJail[i] = ctx.Visitors.InJail[len(ctx.Visitors.InJail)-1]
			ctx.Visitors.InJail = ctx.Visitors.InJail[:len(ctx.Visitors.InJail)-1]
		}

	}

	if !ctx.PlayerCanMove(playerID) {
		ctx.Players.CanMove = append(ctx.Players.CanMove, playerID)
	}
}

func (ctx *Context) RemovePlayerFromMoveable(pID PlayerID) {
	for i, cM := range ctx.Players.CanMove {
		if pID == cM {
			ctx.Players.CanMove[i] = ctx.Players.CanMove[len(ctx.Players.CanMove)-1]
			ctx.Players.CanMove = ctx.Players.CanMove[:len(ctx.Players.CanMove)-1]
		}
	}
}

func (ctx *Context) ProcessJail() {
	for _, iJV := range ctx.Visitors.InJail {
		visitorID := iJV.visitorID
		turnsLeft := iJV.TurnsLeft

		if turnsLeft <= 0 {
			ctx.RemovePlayerFromJail(visitorID)
		} else {
			ctx.RemovePlayerFromMoveable(visitorID)
		}

	}
}

var ErrNotEnoughJailCards = errors.New("Cannot use jail card: player does not have enough get out of jail free cards")
var ErrNotEnoughMoney = errors.New("Cannot execute action: player does not have enough money")

func (ctx *Context) JailUseCard() error {
	currID := ctx.Turn.Current
	if ctx.Players.Alive[currID.Index()].GetOutOfJailCards > 0 {
		ctx.RemovePlayerFromJail(currID)
		ctx.Players.Alive[currID.Index()].GetOutOfJailCards -= 1
		return nil

	} else {
		return ErrNotEnoughJailCards
	}
}

func (ctx *Context) JailBuyout() error {
	currID := ctx.Turn.Current
	if ctx.Players.Alive[currID.Index()].Money >= JailBuyoutCost {
		ctx.RemovePlayerFromJail(currID)
		ctx.AdjustPlayerMoney(currID, -JailBuyoutCost)
		return nil
	} else {
		return ErrNotEnoughMoney
	}
}
