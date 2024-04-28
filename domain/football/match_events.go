package domain

import (
	. "competition-manager/domain"
	"fmt"
)

type Goal struct {
	MatchEvent

	Type   string // REGULAR | PENALTY | AUTOGOL
	Team   Team
	Scorer Player
	Assist Player
}

type Foul struct {
	MatchEvent

	From Player
	To   Player
}

type Booking struct {
	MatchEvent

	Team   Team
	Player Player
	Card   string // YELLOW | RED
}

type Substitution struct {
	MatchEvent

	PlayerIn  Player
	PlayerOut Player

	Team Team
}

type Penalty struct {
	MatchEvent

	Player Player
	Team   Team
	Scored bool
}

func prueba() {
	var events = []MatchEventer{
		&Goal{}, &Penalty{},
	}

	// var a interfaces.IMatchEvent = &Goal{}
	var v = MatchStats{
		Events: events,
	}

	fmt.Println(v)
}
