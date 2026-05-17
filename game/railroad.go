package game

func numRailroadOwned(playerID int32) int32 {
	var ownedCount int32 = 0
	for propID, ownerID := range PropertyOwners {
		if ownerID == playerID && OwnablePropertyType[int32(propID)] == TypeRailroad {
			ownedCount++
		}
	}

	return ownedCount
}

func ProcessOwnedRailroad() {
	for _, oRV := range OwnedRailroadVisitors {
		visitorID := oRV.visitorID
		ownerID := oRV.ownerID
		// railroadID := oRV.railroadID

		var rent int32 = RailroadRent[numRailroadOwned(ownerID)]

		AdjustPlayerMoney(visitorID, -rent)
		AdjustPlayerMoney(ownerID, rent)
		// reset railroad rent mod after payment
		ModifierRailroadRentMultiplier = 1
	}
}
