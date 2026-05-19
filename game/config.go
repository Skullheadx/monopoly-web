package game

const (
	StartingDiceRolls int32 = 1
)

// Board config
const (
	TypeGo PropertyType = iota
	TypeColor
	TypeChest
	TypeTax
	TypeRailroad
	TypeChance
	TypeJail
	TypeUtility
	TypeParking
	TypePolice
)

const (
	GroupBrown ColorGroup = iota
	GroupLightBlue
	GroupPink
	GroupOrange
	GroupRed
	GroupYellow
	GroupGreen
	GroupDarkBlue
)

var ColorGroupSizes = [...]int32{
	GroupBrown:     2,
	GroupLightBlue: 3,
	GroupPink:      3,
	GroupOrange:    3,
	GroupRed:       3,
	GroupYellow:    3,
	GroupGreen:     3,
	GroupDarkBlue:  2,
}

var BoardSpaces = [...]Space{
	Space{PropertyType: TypeGo, SubIndexID: 0},
	Space{PropertyType: TypeColor, SubIndexID: 0},
	Space{PropertyType: TypeChest, SubIndexID: 0},
	Space{PropertyType: TypeColor, SubIndexID: 1},
	Space{PropertyType: TypeTax, SubIndexID: 0},
	Space{PropertyType: TypeRailroad, SubIndexID: 0},
	Space{PropertyType: TypeColor, SubIndexID: 2},
	Space{PropertyType: TypeChance, SubIndexID: 0},
	Space{PropertyType: TypeColor, SubIndexID: 3},
	Space{PropertyType: TypeColor, SubIndexID: 4},
	Space{PropertyType: TypeJail, SubIndexID: 0},
	Space{PropertyType: TypeColor, SubIndexID: 5},
	Space{PropertyType: TypeUtility, SubIndexID: 0},
	Space{PropertyType: TypeColor, SubIndexID: 6},
	Space{PropertyType: TypeColor, SubIndexID: 7},
	Space{PropertyType: TypeRailroad, SubIndexID: 1},
	Space{PropertyType: TypeColor, SubIndexID: 8},
	Space{PropertyType: TypeChest, SubIndexID: 1},
	Space{PropertyType: TypeColor, SubIndexID: 9},
	Space{PropertyType: TypeColor, SubIndexID: 10},
	Space{PropertyType: TypeParking, SubIndexID: 0},
	Space{PropertyType: TypeColor, SubIndexID: 11},
	Space{PropertyType: TypeChance, SubIndexID: 1},
	Space{PropertyType: TypeColor, SubIndexID: 12},
	Space{PropertyType: TypeColor, SubIndexID: 13},
	Space{PropertyType: TypeRailroad, SubIndexID: 2},
	Space{PropertyType: TypeColor, SubIndexID: 14},
	Space{PropertyType: TypeColor, SubIndexID: 15},
	Space{PropertyType: TypeUtility, SubIndexID: 1},
	Space{PropertyType: TypeColor, SubIndexID: 16},
	Space{PropertyType: TypePolice, SubIndexID: 0},
	Space{PropertyType: TypeColor, SubIndexID: 17},
	Space{PropertyType: TypeColor, SubIndexID: 18},
	Space{PropertyType: TypeChest, SubIndexID: 2},
	Space{PropertyType: TypeColor, SubIndexID: 19},
	Space{PropertyType: TypeRailroad, SubIndexID: 3},
	Space{PropertyType: TypeChance, SubIndexID: 2},
	Space{PropertyType: TypeColor, SubIndexID: 20},
	Space{PropertyType: TypeTax, SubIndexID: 1},
	Space{PropertyType: TypeColor, SubIndexID: 21},
}

const BoardSpacesLen = len(BoardSpaces)

// Color Config
var ColorProperties = []ColorProperty{
	{GroupID: GroupBrown, Houses: 0, Name: "Mediterranean Avenue", Price: 60},
	{GroupID: GroupBrown, Houses: 0, Name: "Baltic Avenue", Price: 60},
	{GroupID: GroupLightBlue, Houses: 0, Name: "Oriental Avenue", Price: 100},
	{GroupID: GroupLightBlue, Houses: 0, Name: "Vermont Avenue", Price: 100},
	{GroupID: GroupLightBlue, Houses: 0, Name: "Connecticut Avenue", Price: 120},
	{GroupID: GroupPink, Houses: 0, Name: "St. Charles Place", Price: 140},
	{GroupID: GroupPink, Houses: 0, Name: "States Avenue", Price: 140},
	{GroupID: GroupPink, Houses: 0, Name: "Virginia Avenue", Price: 160},
	{GroupID: GroupOrange, Houses: 0, Name: "St. James Place", Price: 180},
	{GroupID: GroupOrange, Houses: 0, Name: "Tennessee Avenue", Price: 180},
	{GroupID: GroupOrange, Houses: 0, Name: "New York Avenue", Price: 200},
	{GroupID: GroupRed, Houses: 0, Name: "Kentucky Avenue", Price: 220},
	{GroupID: GroupRed, Houses: 0, Name: "Indiana Avenue", Price: 220},
	{GroupID: GroupRed, Houses: 0, Name: "Illinois Avenue", Price: 240},
	{GroupID: GroupYellow, Houses: 0, Name: "Atlantic Avenue", Price: 260},
	{GroupID: GroupYellow, Houses: 0, Name: "Ventnor Avenue", Price: 260},
	{GroupID: GroupYellow, Houses: 0, Name: "Marvin Gardens", Price: 280},
	{GroupID: GroupGreen, Houses: 0, Name: "Pacific Avenue", Price: 300},
	{GroupID: GroupGreen, Houses: 0, Name: "North Carolina Avenue", Price: 300},
	{GroupID: GroupGreen, Houses: 0, Name: "Pennsylvania Avenue", Price: 320},
	{GroupID: GroupDarkBlue, Houses: 0, Name: "Park Place", Price: 350},
	{GroupID: GroupDarkBlue, Houses: 0, Name: "Boardwalk", Price: 400},
}

var ColorPropertyRents = [][]int32{
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

// Railroad Config
var RailroadProperties = []PropertyStatic{
	{Name: "Reading Railroad", Price: 200},
	{Name: "Pennsylvania Railroad", Price: 200},
	{Name: "B.&O. Railroad", Price: 200},
	{Name: "Short Line", Price: 200},
}

const RailroadPrice int32 = 200

var RailroadRent = [...]int32{25, 50, 100, 200}

const RailroadMortgageValue int32 = 100

// Utility Config
var UtilityProperties = []PropertyStatic{
	{Name: "Electric Company", Price: 150},
	{Name: "Waterworks", Price: 150},
}

const UtilityPrice int32 = 150

var UtilityRentMult = [...]int32{4, 10}

const UtilityMortgageValue int32 = 75

// Tax Config
var TaxSpaces = []TaxSpace{
	{Name: "Income Tax", Amount: 200},
	{Name: "Luxury Tax", Amount: 100},
}

var BankPlayerID PlayerID = PlayerID{id: -1}

var SpecialSpaces ChanceSpaceIDs

func init() {
	for i, s := range BoardSpaces {
		spaceID := SpaceID{id: int32(i)}

		propertyType := s.PropertyType

		switch propertyType {
		case TypeColor:
			switch ColorProperties[s.SubIndexID].Name {
			case "St. Charles Place":
				SpecialSpaces.StCharlesPlace = spaceID
			case "Boardwalk":
				SpecialSpaces.Boardwalk = spaceID
			case "Illinois Avenue":
				SpecialSpaces.IllinoisAvenue = spaceID
			}
		case TypeRailroad:
			if RailroadProperties[s.SubIndexID].Name == "Reading Railroad" {
				SpecialSpaces.ReadingRailroad = spaceID
			}
		case TypeGo:
			SpecialSpaces.Go = spaceID
		}
	}
}

const JailDefaultTurns int32 = 3
const JailBuyoutCost int32 = 50

const GoSalary int32 = 200

const ColorMaxHouses = 4

const StartingMoney int32 = 1500
const StartingGetOutOfJailFreeCards int32 = 0

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
