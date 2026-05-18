package game

var ChanceCards = [...]string{
	0:  "Advance to St. Charles Place. If you pass Go, collect $200.",
	1:  "Advance to the nearest Railroad. If unowned, you may buy it from the Bank. If owned, pay owner twice the rental to which they are otherwise entitled.",
	2:  "Advance to the nearest Railroad. If unowned, you may buy it from the Bank. If owned, pay owner twice the rental to which they are otherwise entitled.",
	3:  "Advance token to nearest Utility. If unowned, you may buy it from the Bank. If owned, throw dice and pay owner a total ten times amount thrown.",
	4:  "Go to Jail. Go directly to Jail, do not pass Go, do not collect $200.",
	5:  "Take a trip to Reading Railroad. If you pass Go, collect $200.",
	6:  "Advance to Go (Collect $200)",
	7:  "Get Out of Jail Free.",
	8:  "Advance to Boardwalk.",
	9:  "Your building loan matures. Collect $150.",
	10: "Go Back 3 Spaces.",
	11: "Speeding fine $15.",
	12: "Advance to Illinois Avenue. If you pass Go, collect $200.",
	13: "Make general repairs on all your property. For each house pay $25. For each hotel pay $100.",
	14: "You have been elected Chairman of the Board. Pay each player $50.",
	15: "Bank pays you dividend of $50.",
}

func GetPlayerMoveDistance(start int32, dest int32) int32 {
	distance := dest - start
	if distance < 0 {
		distance += int32(len(BoardSpaces))
	}
	return distance
}

func ProcessChance() {
	for _, visitorID := range ChanceVisitors {
		card := RandSrc.IntN(len(ChanceCards))

		currentPos := Users[visitorID].CurrentSpaceID

		switch card {
		case 0:
			MoveQueue = append(MoveQueue, GetPlayerMoveDistance(currentPos, StCharlesPlaceSpaceID))
		case 1, 2:
			for i := range BoardSpaces {
				offset := int32(i)
				next_pos := currentPos + offset
				propertyType := BoardSpaces[next_pos]
				if propertyType == TypeRailroad {
					distance := GetPlayerMoveDistance(currentPos, next_pos)
					if PropertyOwners[SpaceToOwnableProperty[CalculateNextPos(currentPos, distance)]] != visitorID {
						ModifierRailroadRentMultiplier = 2
					}
					MoveQueue = append(MoveQueue, distance)
				}
			}
		case 3:
			for i := range BoardSpaces {
				offset := int32(i)
				next_pos := currentPos + offset
				propertyType := BoardSpaces[next_pos]
				if propertyType == TypeUtility {
					distance := GetPlayerMoveDistance(currentPos, next_pos)
					if PropertyOwners[SpaceToOwnableProperty[CalculateNextPos(currentPos, distance)]] != visitorID {
						ModifierUtilityForceRentMultiplier = true
					}
					MoveQueue = append(MoveQueue, distance)
				}
			}

		case 4:
			InJailVisitors = append(InJailVisitors, InJailVisitor{visitorID: visitorID, turns: DEFAULT_JAIL_TURNS})
		case 5:
			MoveQueue = append(MoveQueue, GetPlayerMoveDistance(currentPos, ReadingRailroadSpaceID))
		case 6:
			MoveQueue = append(MoveQueue, GetPlayerMoveDistance(currentPos, GoSpaceID))
		case 7:
			Users[visitorID].GetOutOfJailCards++
		case 8:
			MoveQueue = append(MoveQueue, GetPlayerMoveDistance(currentPos, BoardwalkSpaceID))
		case 9:
			AdjustPlayerMoney(visitorID, 150)
		case 10:
			MoveQueue = append(MoveQueue, -3)
		case 11:
			AdjustPlayerMoney(visitorID, -15)
		case 12:
			MoveQueue = append(MoveQueue, GetPlayerMoveDistance(currentPos, IllinoisAvenueSpaceID))
		case 13:
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
		case 14:
			for i := range Users {
				pID := int32(i)
				if pID == visitorID {
					AdjustPlayerMoney(pID, -50*int32(len(Users)-1))
				} else {
					AdjustPlayerMoney(pID, 50)
				}

			}
		case 15:
			AdjustPlayerMoney(visitorID, 50)
		}
	}
}
