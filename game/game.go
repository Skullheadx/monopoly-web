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

var TurnPlayerID int32 = 0
var TurnEndedSignal bool = false
var DiceRollsRemaining int32 = 1
var numDiceRolled int32 = 0
var Doubles bool = false

func ValidateCanRoll(UUID string) bool {
	if Users[TurnPlayerID].UUID == UUID && DiceRollsRemaining > 0 {
		return true
	}

	return false
}

func ValidateCanEndTurn(UUID string) bool {
	if Users[TurnPlayerID].UUID != UUID || DiceRollsRemaining > 0 || Users[TurnPlayerID].Money < 0 {
		return false
	}
	return true
}

func RollDice() {
	// Roll Dice
	diceRoll1 := RandSrc.Int32N(6) + 1
	diceRoll2 := RandSrc.Int32N(6) + 1

	numDiceRolled++

	if diceRoll1 == diceRoll2 {
		Doubles = true
		DiceRollsRemaining++
	}

	if numDiceRolled >= 3 {
		// TODO: GO TO JAIL
	}

	// Movement
	for {
		// condition to stop moving
		if len(MoveQueue) == 0 {
			break
		}

		// does one movement
		// TODO: prevent movement while in JAIL
		ProcessMovement()

		ProcessGo()
		ProcessTax()

		ProcessOwnedColors()
		ProcessOwnedUtility()
		ProcessOwnedRailroad()

		ProcessChance()
		ProcessChest()

		// ProcessPolice()
		// ProcessJail()
	}

}

func EndTurn() {
	TurnPlayerID = (TurnPlayerID + 1) % int32(len(Users))
	TurnEndedSignal = false
	DiceRollsRemaining = 1
	numDiceRolled = 0
	Doubles = false
}
