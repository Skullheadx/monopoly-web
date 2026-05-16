package game

const BoardSpaces = 40

var BoardSpaceTypes = [BoardSpaces]PropertyType{
	0:  TypeGo,
	1:  TypeColor,
	2:  TypeChest,
	3:  TypeColor,
	4:  TypeTax,
	5:  TypeRailroad,
	6:  TypeColor,
	7:  TypeChance,
	8:  TypeColor,
	9:  TypeColor,
	10: TypeJail,
	11: TypeColor,
	12: TypeUtility,
	13: TypeColor,
	14: TypeColor,
	15: TypeRailroad,
	16: TypeColor,
	17: TypeChest,
	18: TypeColor,
	19: TypeColor,
	20: TypeParking,
	21: TypeColor,
	22: TypeChance,
	23: TypeColor,
	24: TypeColor,
	25: TypeRailroad,
	26: TypeColor,
	27: TypeColor,
	28: TypeUtility,
	29: TypeColor,
	30: TypeJail,
	31: TypeColor,
	32: TypeColor,
	33: TypeChest,
	34: TypeColor,
	35: TypeRailroad,
	36: TypeChance,
	37: TypeColor,
	38: TypeTax,
	39: TypeColor,
}

var (
	GoVisitors       []int32
	ColorVisitors    []int32
	ChestVisitors    []int32
	TaxVisitors      []int32
	RailroadVisitors []int32
	ChanceVisitors   []int32
	JailVisitors     []int32
	UtilityVisitors  []int32
	ParkingVisitors  []int32
	PoliceVisitors   []int32
)

func AdvancePlayer(playerID int32, currentPosition int32, diceRoll int32) {
	nextPos := (currentPosition + diceRoll)
	if nextPos > BoardSpaces-1 { // Passed Go, but did not land on Go
		GoVisitors = append(GoVisitors, playerID)
	}
	nextPos %= BoardSpaces

	propType := BoardSpaceTypes[nextPos]

	switch propType {
	case TypeGo:
		GoVisitors = append(GoVisitors, playerID)
	case TypeColor:
		ColorVisitors = append(ColorVisitors, playerID)
	case TypeChest:
		ChestVisitors = append(ChestVisitors, playerID)
	case TypeTax:
		TaxVisitors = append(TaxVisitors, playerID)
	case TypeRailroad:
		RailroadVisitors = append(RailroadVisitors, playerID)
	case TypeChance:
		ChanceVisitors = append(ChanceVisitors, playerID)
	case TypeJail:
		JailVisitors = append(JailVisitors, playerID)
	case TypeUtility:
		UtilityVisitors = append(UtilityVisitors, playerID)
	case TypeParking:
		ParkingVisitors = append(ParkingVisitors, playerID)
	case TypePolice:
		PoliceVisitors = append(PoliceVisitors, playerID)
	}
}
