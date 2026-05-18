package game

type User struct {
	UUID              string
	Money             int32
	CurrentSpaceID    int32
	GetOutOfJailCards int32
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
