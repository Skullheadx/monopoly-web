package game

func HasColorMonopoly(playerID int32, targetGroup ColorGroup) bool {
	var ownedCount int32
	for colorID, ownerID := range PropertyOwners {
		if ownerID == playerID && OwnablePropertyType[RespPropertyToOwnable[int32(colorID)]] == TypeColor && ColorProperties[colorID].GroupID == targetGroup {
			ownedCount++
		}
	}

	return ownedCount == ColorGroupSizes[targetGroup]
}

func processOwnedColors() {
	for _, oCV := range OwnedColorVisitors {
		visitorID := oCV.visitorID
		ownerID := oCV.ownerID
		colorID := oCV.colorID

		prop := ColorProperties[colorID]
		prices := ColorPropertyRents[colorID]

		var rent int32 = prices[prop.Houses]

		if prop.Houses == 0 && HasColorMonopoly(ownerID, prop.GroupID) {
			rent *= 2
		}

		AdjustPlayerMoney(visitorID, -rent)
		AdjustPlayerMoney(ownerID, rent)
	}
}
