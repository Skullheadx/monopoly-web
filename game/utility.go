package game

func (ctx *Context) NumUtilities(playerID PlayerID) int32 {
	var ownedCount int32 = 0
	for _, prop := range ctx.Properties.Owners {
		ownerID := prop.OwnerID

		if ownerID != playerID {
			continue
		}

		spaceID := prop.SpaceID
		space := BoardSpaces[spaceID.Index()]

		propType := space.PropertyType
		if propType != TypeUtility {
			continue
		}

		ownedCount++

	}

	return ownedCount
}

func (ctx *Context) ProcessOwnedUtility() {
	for _, oUV := range ctx.Visitors.Utility {
		visitorID := oUV.visitorID
		ownerID := oUV.ownerID
		// utilityID := oUV.utilityID
		diceRoll := oUV.diceRoll

		var rent int32 = 0

		if !ctx.Turn.Modifier.UtilityForceRentMultiplier {
			rent = UtilityRentMult[ctx.NumUtilities(ownerID)] * diceRoll
		} else {
			rent = 10 * diceRoll
		}

		ctx.AdjustPlayerMoney(visitorID, -rent)
		ctx.AdjustPlayerMoney(ownerID, rent)

		ctx.Turn.Modifier.UtilityForceRentMultiplier = false
	}
}
