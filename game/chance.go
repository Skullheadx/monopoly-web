package game

func (ctx *Context) getOwnerID(sID SpaceID) PlayerID {
	for _, oP := range ctx.Properties.Owners {
		if oP.SpaceID == sID {
			return oP.OwnerID
		}
	}
	return BankPlayerID
}

func (ctx *Context) ProcessChance() {
	for _, visitorID := range ctx.Visitors.Chance {
		card := ctx.Random.IntN(len(ChanceCards))

		currentPos := ctx.Players.Alive[visitorID.Index()].CurrentSpaceID

		switch card {
		case 0:
			ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, GetPlayerMoveDistance(currentPos, SpecialSpaces.StCharlesPlace))
		case 1, 2:
			for i := range BoardSpaces {
				offset := int32(i)
				next_pos := Add(currentPos, offset)
				prop := BoardSpaces[next_pos.Index()]
				if prop.PropertyType == TypeRailroad {
					distance := GetPlayerMoveDistance(currentPos, next_pos)

					if ctx.getOwnerID(CalculateNextPos(currentPos, distance)) != visitorID {
						ctx.Turn.Modifier.RailroadRentMultiplier = 2
					}
					ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, distance)
					break
				}
			}
		case 3:
			for i := range BoardSpaces {
				offset := int32(i)
				next_pos := Add(currentPos, offset)
				prop := BoardSpaces[next_pos.Index()]
				if prop.PropertyType == TypeUtility {
					distance := GetPlayerMoveDistance(currentPos, next_pos)

					if ctx.getOwnerID(CalculateNextPos(currentPos, distance)) != visitorID {
						ctx.Turn.Modifier.UtilityForceRentMultiplier = true
					}
					ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, distance)
					break
				}
			}

		case 4:
			ctx.Visitors.InJail = append(ctx.Visitors.InJail, InJailVisitor{visitorID: visitorID, TurnsLeft: JailDefaultTurns})
		case 5:
			ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, GetPlayerMoveDistance(currentPos, SpecialSpaces.ReadingRailroad))
		case 6:
			ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, GetPlayerMoveDistance(currentPos, SpecialSpaces.Go))
		case 7:
			ctx.Players.Alive[ctx.Turn.Current.Index()].GetOutOfJailCards++
		case 8:
			ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, GetPlayerMoveDistance(currentPos, SpecialSpaces.Boardwalk))
		case 9:
			ctx.AdjustPlayerMoney(visitorID, 150)
		case 10:
			ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, -3)
		case 11:
			ctx.AdjustPlayerMoney(visitorID, -15)
		case 12:
			ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, GetPlayerMoveDistance(currentPos, SpecialSpaces.IllinoisAvenue))
		case 13:
			var repairCost int32 = 0
			for _, prop := range ctx.Properties.Owners {
				ownerID := prop.OwnerID
				spaceID := prop.SpaceID
				space := BoardSpaces[spaceID.Index()]

				if ownerID == visitorID && space.PropertyType == TypeColor {
					colorID := space.SubIndexID
					colorProp := ColorProperties[colorID]
					if colorProp.Houses > ColorMaxHouses { // its a hotel
						repairCost += 100
					} else {
						repairCost += colorProp.Houses * 25
					}
				}
			}

			ctx.AdjustPlayerMoney(visitorID, -repairCost)
		case 14:
			for i := range ctx.Players.Alive {
				pID := PlayerID{id: int32(i)}
				if pID == visitorID {
					ctx.AdjustPlayerMoney(pID, -50*int32(len(ctx.Players.Alive)-1))
				} else {
					ctx.AdjustPlayerMoney(pID, 50)
				}

			}
		case 15:
			ctx.AdjustPlayerMoney(visitorID, 50)
		}
	}
}
