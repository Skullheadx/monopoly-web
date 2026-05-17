package game

const BoardSpaces = 40

var BoardSpaceTypes = [BoardSpaces]PropertyType{
	TypeGo,
	TypeColor,
	TypeChest,
	TypeColor,
	TypeTax,
	TypeRailroad,
	TypeColor,
	TypeChance,
	TypeColor,
	TypeColor,
	TypeJail,
	TypeColor,
	TypeUtility,
	TypeColor,
	TypeColor,
	TypeRailroad,
	TypeColor,
	TypeChest,
	TypeColor,
	TypeColor,
	TypeParking,
	TypeColor,
	TypeChance,
	TypeColor,
	TypeColor,
	TypeRailroad,
	TypeColor,
	TypeColor,
	TypeUtility,
	TypeColor,
	TypeJail,
	TypeColor,
	TypeColor,
	TypeChest,
	TypeColor,
	TypeRailroad,
	TypeChance,
	TypeColor,
	TypeTax,
	TypeColor,
}

type PropertyStatic struct {
	Name  string
	Price int32
}

var ColorProperties = []PropertyStatic{
	{Name: "Mediterranean Avenue", Price: 60},
	{Name: "Baltic Avenue", Price: 60},
	{Name: "Oriental Avenue", Price: 100},
	{Name: "Vermont Avenue", Price: 100},
	{Name: "Connecticut Avenue", Price: 120},
	{Name: "St. Charles Place", Price: 140},
	{Name: "States Avenue", Price: 160},
	{Name: "Virginia Avenue", Price: 140},
	{Name: "St. James Place", Price: 180},
	{Name: "Tennessee Avenue", Price: 200},
	{Name: "Kentucky Avenue", Price: 220},
	{Name: "Indiana Avenue", Price: 220},
	{Name: "Illinois Avenue", Price: 240},
	{Name: "Atlantic Avenue", Price: 260},
	{Name: "Ventnor Avenue", Price: 260},
	{Name: "Marvin Gardens", Price: 280},
	{Name: "Pacific Avenue", Price: 300},
	{Name: "North Carolina Avenue", Price: 300},
	{Name: "Pennsylvania Avenue", Price: 320},
	{Name: "Park Place", Price: 350},
	{Name: "Boardwalk", Price: 400},
}

var RailroadProperties = []PropertyStatic{
	{Name: "Reading Railroad", Price: 200},
	{Name: "Pennsylvania Railroad", Price: 200},
	{Name: "B.&O. Railroad", Price: 200},
	{Name: "Short Line", Price: 200},
}

var UtilityProperties = []PropertyStatic{
	{Name: "Electric Company", Price: 150},
	{Name: "Waterworks", Price: 150},
}

type TaxSpace struct {
	Name   string
	Amount int32
}

var TaxSpaces = []TaxSpace{
	{Name: "Income Tax", Amount: 200},
	{Name: "Luxury Tax", Amount: 100},
}

var PropertyOwners = []int32{}

var SpaceToOwnableProperty = make(map[int]int)
var SpaceToTaxSpace = make(map[int]int)

func init() {
	colorIndex := 0
	railroadIndex := 0
	utilityIndex := 0
	taxIndex := 0
	for i, propertyType := range BoardSpaceTypes {
		switch propertyType {
		case TypeColor:
			SpaceToOwnableProperty[i] = colorIndex
			colorIndex++
		case TypeRailroad:
			SpaceToOwnableProperty[i] = railroadIndex
			railroadIndex++
		case TypeUtility:
			SpaceToOwnableProperty[i] = utilityIndex
			utilityIndex++
		case TypeTax:
			SpaceToTaxSpace[i] = taxIndex
			taxIndex++
		}

	}

}

type ColorVisitor struct {
	playerID int32
	spaceID  int32
}

var (
	GoVisitors       []int32
	ColorVisitors    []ColorVisitor
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
		ColorVisitors = append(ColorVisitors, ColorVisitor{playerID: playerID, spaceID: nextPos})
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
