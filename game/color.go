package game

func (ctx *Context) HasColorMonopoly(playerID PlayerID, targetGroup ColorGroup) bool {
	var ownedCount int32
	for _, prop := range ctx.Properties.Owners {
		ownerID := prop.OwnerID

		if ownerID != playerID {
			continue
		}

		spaceID := prop.SpaceID
		space := BoardSpaces[spaceID.Index()]

		propType := space.PropertyType
		if propType != TypeColor {
			continue
		}

		colorID := space.SubIndexID
		if ColorProperties[colorID].GroupID != targetGroup {
			continue
		}

		// the property belongs to player, is a color property and has the right color
		ownedCount++
	}

	return ownedCount == ColorGroupSizes[targetGroup]
}

func (ctx *Context) ProcessOwnedColors() {
	for _, oCV := range ctx.Visitors.Color {
		visitorID := oCV.visitorID
		ownerID := oCV.ownerID
		colorID := oCV.colorID

		prop := ColorProperties[colorID]
		prices := ColorPropertyRents[colorID]

		var rent int32 = prices[prop.Houses]

		if prop.Houses == 0 && ctx.HasColorMonopoly(ownerID, prop.GroupID) {
			rent *= 2
		}

		ctx.AdjustPlayerMoney(visitorID, -rent)
		ctx.AdjustPlayerMoney(ownerID, rent)
	}
}
