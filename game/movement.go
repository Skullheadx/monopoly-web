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

type OwnableProperty struct {
	Name    string
	Price   int32
	GroupID int32
}

var ColorProperties = []OwnableProperty{
	{GroupID: 0, Name: "Mediterranean Avenue", Price: 60},
	{GroupID: 0, Name: "Baltic Avenue", Price: 60},
	{GroupID: 1, Name: "Oriental Avenue", Price: 100},
	{GroupID: 1, Name: "Vermont Avenue", Price: 100},
	{GroupID: 1, Name: "Connecticut Avenue", Price: 120},
	{GroupID: 2, Name: "St. Charles Place", Price: 140},
	{GroupID: 2, Name: "States Avenue", Price: 140},
	{GroupID: 2, Name: "Virginia Avenue", Price: 160},
	{GroupID: 3, Name: "St. James Place", Price: 180},
	{GroupID: 3, Name: "Tennessee Avenue", Price: 180},
	{GroupID: 3, Name: "New York Avenue", Price: 200},
	{GroupID: 4, Name: "Kentucky Avenue", Price: 220},
	{GroupID: 4, Name: "Indiana Avenue", Price: 220},
	{GroupID: 4, Name: "Illinois Avenue", Price: 240},
	{GroupID: 5, Name: "Atlantic Avenue", Price: 260},
	{GroupID: 5, Name: "Ventnor Avenue", Price: 260},
	{GroupID: 5, Name: "Marvin Gardens", Price: 280},
	{GroupID: 6, Name: "Pacific Avenue", Price: 300},
	{GroupID: 6, Name: "North Carolina Avenue", Price: 300},
	{GroupID: 6, Name: "Pennsylvania Avenue", Price: 320},
	{GroupID: 7, Name: "Park Place", Price: 350},
	{GroupID: 7, Name: "Boardwalk", Price: 400},
}

var ColorPropertyRents = [][6]int32{
	{2, 10, 30, 90, 160, 250},
	{4, 20, 60, 180, 320, 450},
	{6, 30, 90, 270, 400, 550},
	{6, 30, 90, 270, 400, 550},
	{8, 40, 100, 300, 450, 600},
	{10, 50, 150, 450, 625, 750},
	{10, 50, 150, 450, 625, 750},
	{12, 60, 180, 500, 700, 900},
	{14, 70, 200, 550, 750, 950},
	{14, 70, 200, 550, 750, 950},
	{16, 80, 220, 600, 800, 1000},
	{18, 90, 250, 700, 875, 1050},
	{18, 90, 250, 700, 875, 1050},
	{20, 100, 300, 750, 925, 1100},
	{22, 110, 330, 800, 975, 1150},
	{22, 110, 330, 800, 975, 1150},
	{24, 120, 360, 850, 1025, 1200},
	{26, 130, 390, 900, 1100, 1275},
	{26, 130, 390, 900, 1100, 1275},
	{28, 150, 450, 1000, 1200, 1400},
	{35, 175, 500, 1100, 1300, 1500},
	{50, 200, 600, 1400, 1700, 2000},
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
