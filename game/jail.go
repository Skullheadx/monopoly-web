package game

import (
	"errors"
)

const DEFAULT_JAIL_TURNS int32 = 3
const JAIL_BUYOUT_COST int32 = 50

func RemovePlayerFromJail(playerID int32) {
	for i, iJV := range InJailVisitors {
		if playerID == iJV.visitorID {
			InJailVisitors[i] = InJailVisitors[len(InJailVisitors)-1]
			InJailVisitors = InJailVisitors[:len(InJailVisitors)-1]

		}

	}
	MoveablePlayers = append(MoveablePlayers, playerID)
}

func RemovePlayerFromMoveable(pID int32) {
	for i, playerID := range MoveablePlayers {
		if pID == playerID {
			MoveablePlayers[i] = MoveablePlayers[len(MoveablePlayers)-1]
			MoveablePlayers = MoveablePlayers[:len(MoveablePlayers)-1]
		}
	}
}

func ProcessJail() {
	for _, iJV := range InJailVisitors {
		visitorID := iJV.visitorID
		turns := iJV.turns

		if turns <= 0 {
			RemovePlayerFromJail(visitorID)
		} else {
			RemovePlayerFromMoveable(visitorID)
		}

	}
}

var ErrNotEnoughJailCards = errors.New("Cannot use jail card: player does not have enough get out of jail free cards")
var ErrNotEnoughMoney = errors.New("Cannot execute action: player does not have enough money")

func JailUseCard() error {
	if Users[TurnPlayerID].GetOutOfJailCards > 0 {
		RemovePlayerFromJail(TurnPlayerID)
		Users[TurnPlayerID].GetOutOfJailCards -= 1
		return nil

	} else {
		return ErrNotEnoughJailCards
	}
}

func JailBuyout() error {
	if Users[TurnPlayerID].Money >= JAIL_BUYOUT_COST {
		RemovePlayerFromJail(TurnPlayerID)
		AdjustPlayerMoney(TurnPlayerID, -JAIL_BUYOUT_COST)
		return nil
	} else {
		return ErrNotEnoughMoney
	}
}
