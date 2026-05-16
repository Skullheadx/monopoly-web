package game

type User struct {
	UUID  string
	Money int32
}

type PropertyType int

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

type Color struct {
	Name          string
	Price         int32
	Mortgaged     bool
	Rent          [6]int32
	MortgageValue int32
}
