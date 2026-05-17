package game

func numUtilities(playerID int32) int32 {
	var ownedCount int32 = 0
	for utilityID, ownerID := range PropertyOwners {
		if ownerID == playerID && OwnablePropertyType[RespPropertyToOwnable[int32(utilityID)]] == TypeUtility {
			ownedCount++
		}
	}

	return ownedCount
}

func processOwnedUtility() {
	for _, oUV := range OwnedUtilityVisitors {
		visitorID := oUV.visitorID
		ownerID := oUV.ownerID
		// utilityID := oUV.utilityID
		diceRoll := oUV.diceRoll

		var rent int32 = UtilityRentMult[numUtilities(ownerID)] * diceRoll

		AdjustPlayerMoney(visitorID, -rent)
		AdjustPlayerMoney(ownerID, rent)
	}
}
