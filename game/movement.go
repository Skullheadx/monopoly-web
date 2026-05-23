package game

func GetPlayerMoveDistance(start SpaceID, dest SpaceID) int32 {
	distance := dest.id - start.id
	if distance < 0 {
		distance += int32(len(BoardSpaces))
	}
	return distance
}

func (ctx *Context) ProcessMovement() {
	cID := ctx.Turn.Current
	if ctx.PlayerCanMove(cID) {
		dist := ctx.Turn.MoveQueue[0]
		ctx.Turn.MoveQueue = ctx.Turn.MoveQueue[1:]
		ctx.AdvancePlayer(cID, ctx.Players.Alive[cID.Index()].CurrentSpaceID, dist)

		ctx.ProcessLanding()
	}

}

func CalculateNextPos(currentPosition SpaceID, distance int32) SpaceID {
	nextPos := Add(currentPosition, distance)
	nextPos.id %= int32(len(BoardSpaces))

	return nextPos
}

func (ctx *Context) AdvancePlayer(playerID PlayerID, currentPosition SpaceID, diceRoll int32) {
	nextPos := CalculateNextPos(currentPosition, diceRoll)

	numGoPasses := Add(currentPosition, diceRoll).id / int32(len(BoardSpaces))

	if numGoPasses > 0 {
		if BoardSpaces[nextPos.Index()].PropertyType == TypeGo {
			numGoPasses--
		}
		for range numGoPasses {
			ctx.Visitors.Go = append(ctx.Visitors.Go, playerID)
		}
	}

	prop := BoardSpaces[nextPos.Index()]

	switch prop.PropertyType {
	case TypeGo:
		ctx.Visitors.Go = append(ctx.Visitors.Go, playerID)
	case TypeChest:
		ctx.Visitors.Chest = append(ctx.Visitors.Chest, playerID)
	case TypeChance:
		ctx.Visitors.Chance = append(ctx.Visitors.Chance, playerID)
	case TypeTax:
		ctx.Visitors.Tax = append(ctx.Visitors.Tax, TaxVisitor{visitorID: playerID, taxID: prop.SubIndexID})
	case TypePolice: // hardcoding to send straight to jail
		ctx.Visitors.InJail = append(ctx.Visitors.InJail, InJailVisitor{visitorID: playerID, TurnsLeft: JailDefaultTurns})
	case TypeJail:
		ctx.Visitors.InJail = append(ctx.Visitors.InJail, InJailVisitor{visitorID: playerID, TurnsLeft: JailDefaultTurns})
	case TypeColor:
		propID := ctx.getPropID(nextPos)
		ownerID := ctx.Properties.Owners[propID.Index()].OwnerID
		if ownerID != BankPlayerID { // property owned?
			if ownerID != playerID && !ctx.IsMortgaged(propID) { // not by you
				ctx.Visitors.Color = append(ctx.Visitors.Color, OwnedColorVisitor{visitorID: playerID, ownerID: ownerID, colorID: prop.SubIndexID})
			}
		} else {
			ctx.Visitors.Unowned = append(ctx.Visitors.Unowned, UnownedPropertyVisitor{visitorID: playerID, propertyID: propID})
		}
	case TypeRailroad:
		propID := ctx.getPropID(nextPos)
		ownerID := ctx.Properties.Owners[propID.Index()].OwnerID
		if ownerID != BankPlayerID { // property owned?
			if ownerID != playerID && !ctx.IsMortgaged(propID) { // not by you
				ctx.Visitors.Railroad = append(ctx.Visitors.Railroad, OwnedRailroadVisitor{visitorID: playerID, ownerID: ownerID, railroadID: prop.SubIndexID})
			}
		} else {
			ctx.Visitors.Unowned = append(ctx.Visitors.Unowned, UnownedPropertyVisitor{visitorID: playerID, propertyID: propID})
		}
	case TypeUtility:
		propID := ctx.getPropID(nextPos)
		ownerID := ctx.Properties.Owners[propID.Index()].OwnerID
		if ownerID != BankPlayerID { // property owned?
			if ownerID != playerID && !ctx.IsMortgaged(propID) { // not by you
				ctx.Visitors.Utility = append(ctx.Visitors.Utility, OwnedUtilityVisitor{visitorID: playerID, ownerID: ownerID, utilityID: prop.SubIndexID, diceRoll: diceRoll})
			}
		} else {
			ctx.Visitors.Unowned = append(ctx.Visitors.Unowned, UnownedPropertyVisitor{visitorID: playerID, propertyID: propID})
		}
	}
}
