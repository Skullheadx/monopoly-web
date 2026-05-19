package game

func (ctx *Context) numRailroadOwned(playerID PlayerID) int32 {
	var ownedCount int32 = 0
	for _, prop := range ctx.Properties.Owners {
		ownerID := prop.OwnerID

		if ownerID != playerID {
			continue
		}

		spaceID := prop.SpaceID
		space := BoardSpaces[spaceID.Index()]

		propType := space.PropertyType
		if propType != TypeRailroad {
			continue
		}

		ownedCount++

	}

	return ownedCount
}

func (ctx *Context) ProcessOwnedRailroad() {
	for _, oRV := range ctx.Visitors.Railroad {
		visitorID := oRV.visitorID
		ownerID := oRV.ownerID
		// railroadID := oRV.railroadID

		var rent int32 = RailroadRent[ctx.numRailroadOwned(ownerID)]

		ctx.AdjustPlayerMoney(visitorID, -rent)
		ctx.AdjustPlayerMoney(ownerID, rent)

		// reset railroad rent mod after payment
		ctx.Turn.Modifier.RailroadRentMultiplier = 1
	}
}
