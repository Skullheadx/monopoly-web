package game

import (
	"math/rand/v2"
)

func initTurn(pID PlayerID) Turn {
	return Turn{
		Current:            pID,
		Ended:              false,
		DiceRollsRemaining: StartingDiceRolls,
		NumDiceRolled:      0,
		RolledDoubles:      false,
		MoveQueue:          []int32{},
		InDebt:             false,
		Modifier: Modifiers{
			RailroadRentMultiplier:     1,
			UtilityForceRentMultiplier: false,
		},
	}
}

func InitCtx(randSeed rand.Source, players []Player) *Context {
	startingPlayerID := PlayerID{id: 0}
	ownableProps := []OwnableProperty{}
	for i, s := range BoardSpaces {
		spaceID := SpaceID{id: int32(i)}

		propertyType := s.PropertyType
		if propertyType == TypeColor || propertyType == TypeRailroad || propertyType == TypeUtility {
			ownableProps = append(ownableProps, OwnableProperty{
				OwnerID: BankPlayerID,
				SpaceID: spaceID,
			})
		}
	}

	return &Context{
		Random: rand.New(randSeed),
		Players: Players{
			Alive: players,
		},
		Turn: initTurn(startingPlayerID),
		Visitors: Visitors{
			Unowned:  []UnownedPropertyVisitor{},
			Color:    []OwnedColorVisitor{},
			Railroad: []OwnedRailroadVisitor{},
			Utility:  []OwnedUtilityVisitor{},
			Go:       []PlayerID{},
			Tax:      []TaxVisitor{},
			Chance:   []PlayerID{},
			Chest:    []PlayerID{},
			InJail:   []InJailVisitor{},
			// Parking:  []PlayerID{},
			// Police:   []PlayerID{},
		},
		Properties: Properties{
			Owners:    ownableProps,
			Mortgages: []PropertyID{},
		},
	}
}

func InitPlayer() Player {
	return Player{
		UUID:              "abc", // TODO: Generate proper UUID
		Money:             StartingMoney,
		CurrentSpaceID:    SpecialSpaces.Go,
		GetOutOfJailCards: StartingGetOutOfJailFreeCards,
		CanMove:           true,
	}
}

func (ctx *Context) GetCurrentTurnPlayer() *Player {
	return &ctx.Players.Alive[ctx.Turn.Current.Index()]
}

func (ctx *Context) ValidateIsTurn(UUID string) bool {
	if ctx.GetCurrentTurnPlayer().UUID == UUID {
		return true
	}
	return false
}

func (ctx *Context) ValidateCanRoll(UUID string) bool {
	if ctx.ValidateIsTurn(UUID) && ctx.Turn.DiceRollsRemaining > 0 {
		return true
	}
	return false
}

func (ctx *Context) ValidateCanEndTurn(UUID string) bool {
	if !(ctx.ValidateIsTurn(UUID) || ctx.Turn.DiceRollsRemaining > 0 || ctx.GetCurrentTurnPlayer().Money < 0) {
		return true
	}
	return false
}

func (ctx *Context) ValidateCanExitJail(UUID string) bool {
	for _, iJV := range ctx.Visitors.InJail {
		if ctx.Players.Alive[iJV.visitorID.Index()].UUID == UUID {
			return true
		}
	}
	return false
}

func (ctx *Context) ProcessLanding() {
	ctx.ProcessGo()
	ctx.ProcessTax()

	ctx.ProcessOwnedColors()
	ctx.ProcessOwnedUtility()
	ctx.ProcessOwnedRailroad()
	ctx.ProcessUnowned()

	ctx.ProcessChance()
	ctx.ProcessChest()
	// ProcessPolice()

	ctx.ProcessJail()
}

func (ctx *Context) RollDice() {
	// Roll Dice
	diceRoll1 := ctx.Random.Int32N(6) + 1
	diceRoll2 := ctx.Random.Int32N(6) + 1

	ctx.Turn.NumDiceRolled++
	ctx.Turn.DiceRollsRemaining--

	if diceRoll1 == diceRoll2 {
		ctx.Turn.RolledDoubles = true
		ctx.Turn.DiceRollsRemaining++
		ctx.RemovePlayerFromJail(ctx.Turn.Current)
	}

	if ctx.Turn.NumDiceRolled >= 3 {
		ctx.Visitors.InJail = append(ctx.Visitors.InJail, InJailVisitor{visitorID: ctx.Turn.Current, TurnsLeft: JailDefaultTurns})
	}

}

func (ctx *Context) EndTurn() {
	nextTurnPlayerID := PlayerID{id: (ctx.Turn.Current.id + 1) % int32(len(ctx.Players.Alive))}
	ctx.Turn = initTurn(nextTurnPlayerID)
}

func (ctx *Context) IsMortgaged(propID PropertyID) bool {
	for _, oPID := range ctx.Properties.Mortgages {
		if oPID == propID {
			return true
		}
	}
	return false
}

func (ctx *Context) IsOwned(spaceID SpaceID) bool {
	for _, prop := range ctx.Properties.Owners {
		if spaceID == prop.SpaceID {
			if prop.OwnerID == BankPlayerID {
				return false
			} else {
				return true
			}
		}
	}
	panic("Space is not an ownable property")
}

func (ctx *Context) getPropID(spaceID SpaceID) PropertyID {
	for i, prop := range ctx.Properties.Owners {
		if spaceID == prop.SpaceID {
			return PropertyID{id: int32(i)}
		}
	}
	panic("Space is not an ownable property")
}
