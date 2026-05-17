package game

func numUtilities(playerID int32) int32 {
	var ownedCount int32 = 0
	for propID, ownerID := range PropertyOwners {
		if ownerID == playerID && OwnablePropertyType[int32(propID)] == TypeUtility {
			ownedCount++
		}
	}

	return ownedCount
}

func ProcessOwnedUtility() {
	for _, oUV := range OwnedUtilityVisitors {
		visitorID := oUV.visitorID
		ownerID := oUV.ownerID
		// utilityID := oUV.utilityID
		diceRoll := oUV.diceRoll

		var rent int32 = 0

		if !ModifierUtilityForceRentMultiplier {
			rent = UtilityRentMult[numUtilities(ownerID)] * diceRoll
		} else {
			rent = 10 * diceRoll
		}

		AdjustPlayerMoney(visitorID, -rent)
		AdjustPlayerMoney(ownerID, rent)

		ModifierUtilityForceRentMultiplier = false
	}
}
