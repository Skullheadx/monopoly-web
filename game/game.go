package game

import (
	"math/rand/v2"
)

var RandSeed = rand.NewPCG(20, 26)
var RandSrc = rand.New(RandSeed)

var Users []User
var DebtEvents []int32
var MoveQueue []int32

var ModifierRailroadRentMultiplier int32 = 1
var ModifierUtilityForceRentMultiplier bool = false

func IsInDebt(playerID int32) (bool, int32) {
	for i, pID := range DebtEvents {
		if pID == playerID {
			return true, int32(i)
		}
	}
	return false, -1
}

func AdjustPlayerMoney(playerID int32, amount int32) {
	Users[playerID].Money += amount

	inDebt, i := IsInDebt(playerID)

	if Users[playerID].Money < 0 {
		if !inDebt {
			DebtEvents = append(DebtEvents, playerID)
		}
	} else { // Money >= 0
		if inDebt { // remove player from DebtEvents table
			DebtEvents[i] = DebtEvents[len(DebtEvents)-1]
			DebtEvents = DebtEvents[:len(DebtEvents)-1]
		}
	}

}
