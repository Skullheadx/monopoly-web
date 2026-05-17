package game

const DEFAULT_JAIL_TURNS int32 = 3

func RemovePlayerFromJail(playerID int32) {
	for i, iJV := range InJailVisitors {
		if playerID == iJV.visitorID {
			InJailVisitors[i] = InJailVisitors[len(InJailVisitors)-1]
			InJailVisitors = InJailVisitors[:len(InJailVisitors)-1]

		}

	}
}

func RemovePlayerFromMoveable(pID int32) {
	for i, playerID := range MoveablePlayers {
		if pID == playerID {
			MoveablePlayers[i] = MoveablePlayers[len(MoveablePlayers)-1]
			MoveablePlayers = MoveablePlayers[:len(MoveablePlayers)-1]
		}
	}
}

func ProcessJail() {
	for _, iJV := range InJailVisitors {
		visitorID := iJV.visitorID
		turns := iJV.turns

		if turns <= 0 {
			RemovePlayerFromJail(visitorID)
			MoveablePlayers = append(MoveablePlayers, visitorID)
		} else {
			RemovePlayerFromMoveable(visitorID)
		}

	}
}
