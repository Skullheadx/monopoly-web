package game

const MaxPlayers int32 = 8

var Users [MaxPlayers]User

var BoardSpaceTypes = [40]PropertyType{
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

// var ColorRent = [22][6]int{
// 	//b,1h, 2h,3h,4h,hotel
// 	{2, 10, 30, 90, 160, 250},
// }
// var Colors = [22]Color{
// 	{Name: "Mediterranean Avenue", Price: 60, Mortgaged: false, Rent: [6]int32{2, 10, 30, 90, 160, 250}, MortgageValue: 30},
// }
//
// const UtilityPrice int32 = 150
// const UtilityRentMult = [2]int32{4, 10}
// const UtilityMortgageValue int32 = 75
//
// var Utilities = [2]Utility{
// 	{Name: "Electric Company", Mortgaged: false},
// 	{Name: "Waterworks", Mortgaged: false},
// }
//
// const RailroadPrice int32 = 200
// const RailroadRent = [4]int32{25, 50, 100, 200}
// const RailroadMortgageValue int32 = 100
//
// var RailRoads = [4]RailRoad{
// 	{Name: "Reading Railroad", Mortgaged: false},
// 	{Name: "Pennsylvania Railroad", Mortgaged: false},
// 	{Name: "B.&O. Railroad", Mortgaged: false},
// 	{Name: "Short Line", Mortgaged: false},
// }
//
// var Go = [1]Go{
// 	{Name: "Go", Salary: 200},
// }
//
// var Chest = [3]Chest{
// 	{Name: "Community Chest"},
// 	{Name: "Community Chest"},
// 	{Name: "Community Chest"},
// }
