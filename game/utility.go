package game

func processOwnedUtility() {
	for _, oUV := range OwnedUtilityVisitors {
		visitorID := oUV.visitorID
		ownerID := oUV.ownerID
		// utilityID := oUV.utilityID
		diceRoll := oUV.diceRoll

		var rent int32 = UtilityRentMult[HasUtiltyMonopoly(ownerID)] * diceRoll

		AdjustPlayerMoney(visitorID, -rent)
		AdjustPlayerMoney(ownerID, rent)
	}
}
