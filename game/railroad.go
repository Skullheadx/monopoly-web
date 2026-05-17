package game

func numRailroadOwned(playerID int32) int32 {
	var ownedCount int32 = 0
	for railroadID, ownerID := range PropertyOwners {
		if ownerID == playerID && OwnablePropertyType[RespPropertyToOwnable[int32(railroadID)]] == TypeRailroad {
			ownedCount++
		}
	}

	return ownedCount
}

func processOwnedRailroad() {
	for _, oRV := range OwnedRailroadVisitors {
		visitorID := oRV.visitorID
		ownerID := oRV.ownerID
		// railroadID := oRV.railroadID

		var rent int32 = RailroadRent[numRailroadOwned(ownerID)]

		AdjustPlayerMoney(visitorID, -rent)
		AdjustPlayerMoney(ownerID, rent)
	}
}
