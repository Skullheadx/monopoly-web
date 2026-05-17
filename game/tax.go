package game

func ProcessTax() {
	for _, playerID := range TaxVisitors {
		taxID := SpaceToTaxSpace[Users[playerID].CurrentSpaceID]
		AdjustPlayerMoney(playerID, -TaxSpaces[taxID].Amount)
	}
}
