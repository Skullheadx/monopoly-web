package game

import (
	"math/rand/v2"
)

type Player struct {
	UUID              string
	Money             int32
	CurrentSpaceID    SpaceID
	GetOutOfJailCards int32
	CanMove           bool
}

type PropertyType int
type Space struct {
	PropertyType PropertyType
	SubIndexID   int32 // FK to resp. OwnableProperty types
}

type PropertyStatic struct {
	Name  string
	Price int32
}

type OwnableProperty struct {
	OwnerID PlayerID
	SpaceID SpaceID
}

type ColorGroup int32

type ColorProperty struct {
	Name    string
	Price   int32
	GroupID ColorGroup
	Houses  int32
}

type TaxSpace struct {
	Name   string
	Amount int32
}

type ChanceSpaceIDs struct {
	StCharlesPlace  SpaceID
	Go              SpaceID
	ReadingRailroad SpaceID
	Boardwalk       SpaceID
	IllinoisAvenue  SpaceID
}

type OwnedColorVisitor struct {
	visitorID PlayerID
	ownerID   PlayerID
	colorID   int32
}

type OwnedRailroadVisitor struct {
	visitorID  PlayerID
	ownerID    PlayerID
	railroadID int32
}

type OwnedUtilityVisitor struct {
	visitorID PlayerID
	ownerID   PlayerID
	utilityID int32
	diceRoll  int32
}

type UnownedPropertyVisitor struct {
	visitorID  PlayerID
	propertyID PropertyID
}

type InJailVisitor struct {
	visitorID PlayerID
	TurnsLeft int32
}

type TaxVisitor struct {
	visitorID PlayerID
	taxID     int32
}

type Visitors struct { // SubIndexID is the PK for each resp table
	Unowned  []UnownedPropertyVisitor
	Color    []OwnedColorVisitor
	Railroad []OwnedRailroadVisitor
	Utility  []OwnedUtilityVisitor
	Go       []PlayerID
	Tax      []TaxVisitor
	Chance   []PlayerID
	Chest    []PlayerID
	InJail   []InJailVisitor
	// Parking  []PlayerID
	// Police   []PlayerID
}

// IDs

type PlayerID struct {
	id int32
}

type PropertyID struct {
	id int32
}

type SpaceID struct {
	id int32
}

func (p *PlayerID) Index() int {
	return int(p.id)
}

func (p *PropertyID) Index() int {
	return int(p.id)
}

func (s *SpaceID) Index() int {
	return int(s.id)
}

func Add(a SpaceID, b int32) SpaceID {
	return SpaceID{id: a.id + b}
}

// Game Context (Ctx)
type Modifiers struct {
	RailroadRentMultiplier     int32
	UtilityForceRentMultiplier bool
}

type Players struct {
	Alive []Player // PlayerID is PK
}

type Turn struct {
	Current            PlayerID
	Ended              bool
	DiceRollsRemaining int32
	NumDiceRolled      int32
	RolledDoubles      bool
	MoveQueue          []int32
	InDebt             bool
	Modifier           Modifiers
}

type Properties struct {
	Owners    []OwnableProperty // PropertyID is PK
	Mortgages []PropertyID
}

type Context struct {
	Random     *rand.Rand
	MoveQueue  []int32
	Players    Players
	Turn       Turn // state reset every turn end
	Visitors   Visitors
	Properties Properties
}
