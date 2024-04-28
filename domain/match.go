package domain

// FINAL | THIRD_PLACE | SEMI_FINALS | QUARTER_FINALS | LAST_16 | LAST_32 | LAST_64 | ROUND_4 | ROUND_3 | ROUND_2
// | ROUND_1 | GROUP_STAGE | PRELIMINARY_ROUND | QUALIFICATION | QUALIFICATION_ROUND_1 | QUALIFICATION_ROUND_2
// | QUALIFICATION_ROUND_3 | PLAYOFF_ROUND_1 | PLAYOFF_ROUND_2 | PLAYOFFS | REGULAR_SEASON | CLAUSURA | APERTURA
// | CHAMPIONSHIP | RELEGATION | RELEGATION_ROUND
// https://builtin.com/software-engineering-perspectives/golang-enum
// type Stage string

// CANCELLED | POSTPONED | FINISH | INPROGRESS | SUSPENDED ... [SCHEDULED | LIVE | IN_PLAY | PAUSED | FINISHED | POSTPONED | SUSPENDED | CANCELLED]
type MatchStatus string

type Match struct {
	Competition Competition
	Season      Season
	Stage       string
	StageOrder  int
	Number      int
	// Venue       string // podría ser un struct

	HomeTeam MatchTeam
	AwayTeam MatchTeam

	Score Score

	Statistics MatchStats

	Referees []Referee
	Status   MatchStatus
}

type Referee struct {
	Person Player
	Type   string // ASSISTANT_REFEREE_N1 | ASSISTANT_REFEREE_N2 | FOURTH_OFFICIAL | VIDEO_ASSISTANT_REFEREE_N1 | VIDEO_ASSISTANT_REFEREE_N2
}

type MatchTeam struct {
	Formation string
	LineUp    []MatchTeamPlayer
	Bench     []MatchTeamPlayer
	Coach     Coach
	Staff     []Staff
	Team      Team
}

type MatchTeamPlayer struct {
	Player      Player
	Name        string
	Position    string
	ShirtNumber int
}

type Score struct {
	Winner string // HOME | AWAY
	Result ScoreGoals
}

// should be embbeded
type ScoreGoals struct {
	Home int
	Away int
}

type MatchEventer interface {
	GetMinute() int
}

type MatchStats struct {
	// una cosa es dominio y otra modelo
	Events []MatchEventer
}

type MatchEvent struct {
	Minute int
}

func (m *MatchEvent) GetMinute() int {
	return m.Minute
}
