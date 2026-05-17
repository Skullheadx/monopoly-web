package game

const MAX_PROP_HOUSES = 4

func HasColorMonopoly(playerID int32, targetGroup ColorGroup) bool {
	var ownedCount int32
	for propID, ownerID := range PropertyOwners {
		if ownerID == playerID && OwnablePropertyType[int32(propID)] == TypeColor && ColorProperties[OwnableToRespProperty[int32(propID)]].GroupID == targetGroup {
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
