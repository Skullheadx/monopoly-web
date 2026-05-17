package game

func IsInDebt(playerID int32) (bool, int32) {
	for i, pID := range DebtEvents {
		if pID == playerID {
			return true, int32(i)
		}
	}
	return false, -1
}

func AdjustPlayerMoney(playerID int32, amount int32) {
	Users[playerID].Money += amount

	inDebt, i := IsInDebt(playerID)

	if Users[playerID].Money < 0 {
		if !inDebt {
			DebtEvents = append(DebtEvents, playerID)
		}
	} else { // Money >= 0
		if inDebt { // remove player from DebtEvents table
			DebtEvents[i] = DebtEvents[len(DebtEvents)-1]
			DebtEvents = DebtEvents[:len(DebtEvents)-1]
		}
	}

}
