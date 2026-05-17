package game

var ChestCards = [...]string{
	0:  "Advance to Go. (Collect $200)",
	1:  "Bank error in your favor. Collect $200.",
	2:  "Holiday fund matures. Receive $100.",
	3:  "Life insurance matures. Collect $100.",
	4:  "You inherit $100.",
	5:  "From sale of stock you get $50.",
	6:  "Income tax refund. Collect $20.",
	7:  "It is your birthday. Collect $10 from every player.",
	8:  "You are assessed for street repair. $40 per house. $115 per hotel.",
	9:  "Pay hospital fees of $100.",
	10: "Pay school fees of $50.",
	11: "Doctor’s fee. Pay $50.",
	12: "Receive $25 consultancy fee.",
	13: "Get Out of Jail Free.",
	14: "Go to Jail. Go directly to jail, do not pass Go, do not collect $200.",
	15: "You have won second prize in a beauty contest. Collect $10.",
}

func ProcessChest() {
	for _, visitorID := range ChestVisitors {
		card := RandSrc.IntN(len(ChanceCards))

		currentPos := Users[visitorID].CurrentSpaceID

		switch card {
		case 0:
			MoveQueue = append(MoveQueue, GetPlayerMoveDistance(currentPos, GoSpaceID))
		case 1:
			AdjustPlayerMoney(visitorID, 200)
		case 2:
			AdjustPlayerMoney(visitorID, 100)
		case 3:
			AdjustPlayerMoney(visitorID, 100)
		case 4:
			AdjustPlayerMoney(visitorID, 100)
		case 5:
			AdjustPlayerMoney(visitorID, 50)
		case 6:
			AdjustPlayerMoney(visitorID, 20)
		case 7:
			for i := range Users {
				pID := int32(i)
				if pID == visitorID {
					AdjustPlayerMoney(pID, 10*int32(len(Users)-1))
				} else {
					AdjustPlayerMoney(pID, -10)
				}

			}
		case 8:
			var repairCost int32 = 0
			for propID, ownerID := range PropertyOwners {
				if ownerID == visitorID && OwnablePropertyType[propID] == TypeColor {
					colorID := OwnableToRespProperty[int32(propID)]
					colorProp := ColorProperties[colorID]
					if colorProp.Houses > MAX_PROP_HOUSES { // its a hotel
						repairCost += 100
					} else {
						repairCost += colorProp.Houses * 25
					}
				}
			}

			AdjustPlayerMoney(visitorID, -repairCost)
		case 9:
			AdjustPlayerMoney(visitorID, -100)
		case 10:
			MoveQueue = append(MoveQueue, -50)
		case 11:
			AdjustPlayerMoney(visitorID, -50)
		case 12:
			AdjustPlayerMoney(visitorID, 25)
		case 13:
			InJailVisitors = append(InJailVisitors, InJailVisitor{visitorID: visitorID, turns: DEFAULT_JAIL_TURNS})
		case 14:
			// TODO: GET OUT OF JAIL FREE
		case 15:
			AdjustPlayerMoney(visitorID, 10)
		}
	}
}
