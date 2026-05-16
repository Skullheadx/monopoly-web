package game

import (
	"monopoly-web/types"
)

const MaxPlayers int32 = 8

var Users [MaxPlayers]types.User

var Properties = [22]types.Property{
	{Name: "Mediterranean Avenue", Price: 60, Mortgaged: false, Rent: [6]int32{2, 10, 30, 90, 160, 250}, MortgageValue: 30},
}

var Rent = [22][6]int{
	//b,1h, 2h,3h,4h,hotel
	{2, 10, 30, 90, 160, 250},
}
