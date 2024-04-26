package domain

import "time"

type Competition struct {
	Type          string // LEAGUE | CUP | SUPER_CUP | PLAYOFFS | ...
	Seasons       []Season
	CurrentSeason Season // quizás tiene que ser un método
}

type Season struct {
	Winner Team
	Start  time.Time
	End    time.Time

	Fixture []Match
}
