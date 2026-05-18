package game

import (
	"math/rand/v2"
)

var RandSeed = rand.NewPCG(20, 26)
var RandSrc = rand.New(RandSeed)

var Users []User
var DebtEvents []int32
var MoveablePlayers []int32
var MoveQueue []int32

var ModifierRailroadRentMultiplier int32 = 1
var ModifierUtilityForceRentMultiplier bool = false

var TurnPlayerID int32 = 0
var TurnEndedSignal bool = false
var DiceRollsRemaining int32 = 1
var numDiceRolled int32 = 0
var RolledDoubles bool = false

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

func ValidateCanExitJail(UUID string) bool {
	for _, iJV := range InJailVisitors {
		player := Users[iJV.visitorID]
		if Users[TurnPlayerID].UUID == UUID && player.UUID == UUID {
			return true
		}
	}
	return false
}

func ProcessLanding() {
	ProcessGo()
	ProcessTax()

	ProcessOwnedColors()
	ProcessOwnedUtility()
	ProcessOwnedRailroad()

	ProcessChance()
	ProcessChest()
	// ProcessPolice()

	ProcessJail()
}

func RollDice() {
	// Roll Dice
	diceRoll1 := RandSrc.Int32N(6) + 1
	diceRoll2 := RandSrc.Int32N(6) + 1

	numDiceRolled++

	if diceRoll1 == diceRoll2 {
		RolledDoubles = true
		DiceRollsRemaining++
		RemovePlayerFromJail(TurnPlayerID)
	}

	if numDiceRolled >= 3 {
		InJailVisitors = append(InJailVisitors, InJailVisitor{visitorID: TurnPlayerID, turns: DEFAULT_JAIL_TURNS})
	}

}

func EndTurn() {
	// next player's turn
	TurnPlayerID = (TurnPlayerID + 1) % int32(len(Users))
	TurnEndedSignal = false

	// reset dice
	DiceRollsRemaining = 1
	numDiceRolled = 0
	RolledDoubles = false

	// let player move next turn
	MoveQueue = MoveQueue[:0]
}
