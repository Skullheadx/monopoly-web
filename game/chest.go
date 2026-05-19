package game

func (ctx *Context) ProcessChest() {
	for _, visitorID := range ctx.Visitors.Chest {
		card := ctx.Random.IntN(len(ChanceCards))

		currentPos := ctx.Players.Alive[visitorID.Index()].CurrentSpaceID

		switch card {
		case 0:
			ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, GetPlayerMoveDistance(currentPos, SpecialSpaces.Go))
		case 1:
			ctx.AdjustPlayerMoney(visitorID, 200)
		case 2:
			ctx.AdjustPlayerMoney(visitorID, 100)
		case 3:
			ctx.AdjustPlayerMoney(visitorID, 100)
		case 4:
			ctx.AdjustPlayerMoney(visitorID, 100)
		case 5:
			ctx.AdjustPlayerMoney(visitorID, 50)
		case 6:
			ctx.AdjustPlayerMoney(visitorID, 20)
		case 7:
			for i := range ctx.Players.Alive {
				pID := PlayerID{id: int32(i)}
				if pID == visitorID {
					ctx.AdjustPlayerMoney(pID, 10*int32(len(ctx.Players.Alive)-1))
				} else {
					ctx.AdjustPlayerMoney(pID, -10)
				}

			}
		case 8:
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
		case 9:
			ctx.AdjustPlayerMoney(visitorID, -100)
		case 10:
			ctx.Turn.MoveQueue = append(ctx.Turn.MoveQueue, -50)
		case 11:
			ctx.AdjustPlayerMoney(visitorID, -50)
		case 12:
			ctx.AdjustPlayerMoney(visitorID, 25)
		case 13:
			ctx.Visitors.InJail = append(ctx.Visitors.InJail, InJailVisitor{visitorID: visitorID, TurnsLeft: JailDefaultTurns})
		case 14:
			ctx.Players.Alive[ctx.Turn.Current.Index()].GetOutOfJailCards++
		case 15:
			ctx.AdjustPlayerMoney(visitorID, 10)
		}
	}
}
