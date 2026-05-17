package game

var Users []User
var DebtEvents []int32

func HasColorMonopoly(playerID int32, targetGroup ColorGroup) bool {
	var ownedCount int32
	for colorID, ownerID := range PropertyOwners {
		if ownerID == playerID && ColorProperties[colorID].GroupID == targetGroup {
			ownedCount++
		}
	}

	return ownedCount == ColorGroupSizes[targetGroup]
}

func HasUtiltyMonopoly(playerID int32) bool {
	var ownedCount int32
	for _, ownerID := range PropertyOwners {
		if ownerID == playerID {
			ownedCount++
		}
	}

	return ownedCount == int32(len(UtilityProperties))
}

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
