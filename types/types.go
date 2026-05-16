package types

type User struct {
	UUID  string
	Money int32
}

type Square struct {
	Name  string
	Class string
}

type Property struct {
	Name          string
	Price         int32
	Mortgaged     bool
	Rent          [6]int32
	MortgageValue int32
}
