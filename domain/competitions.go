package domain

// https://es.wikipedia.org/wiki/Categor%C3%ADa:Sistemas_de_competici%C3%B3n
import (
	"fmt"
	"math"
	"time"

	"gonum.org/v1/gonum/stat/combin"
)

// TODO: estrategia que calcula los puntos según una lista de partidos
type Competition struct {
	Type          string // LEAGUE | CUP | SUPER_CUP | PLAYOFFS | ...
	Sport         string // FOOTBALL | PADDLE | ...
	Seasons       []Season
	CurrentSeason *Season // quizás tiene que ser un método
}

type Season struct {
	Winner Team
	Start  time.Time
	End    *time.Time

	Fixture []Match
}

type CompetitionBuilder interface {
	CreateLeague(teams int) CompetitionBuilder
	CreateCup(teams int) CompetitionBuilder

	Build() *Competition
}

type competitionBuilder struct {
	Competition *Competition
}

func NewCompetitionBuilder() CompetitionBuilder {
	return &competitionBuilder{}
}

func (b *competitionBuilder) CreateLeague(teams int) CompetitionBuilder {
	// algoritmo
	// cantidad de partidos por fecha: equipos / 2
	// cantidad de partidos: combinatoria(equipos, tomados de a 2)
	// cantidad de fechas: cantidad de partidos / cantidad de fechas
	// estrategia para puntos: https://futbol.fandom.com/es/wiki/Sistema_de_todos_contra_todos

	competition := &Competition{
		Type:    "LEAGUE",
		Seasons: []Season{},
	}

	competition.CurrentSeason = newSeasonLeague(teams)

	return &competitionBuilder{Competition: competition}
}

func (b *competitionBuilder) CreateCup(teams int) CompetitionBuilder {
	// algoritmo
	// por cada etapa hay 2**(etapa-1) partidos
	// la cantidad de etapas está dada por log2(equipos)
	// en caso de impares hay un equipo flotante

	competition := &Competition{
		Type:    "CUP",
		Seasons: []Season{},
	}

	competition.CurrentSeason = newSeasonCup(teams)

	return &competitionBuilder{Competition: competition}
}

func (b *competitionBuilder) Build() *Competition {
	// TODO: Validate if it is correct created
	return b.Competition
}

func newSeasonLeague(teams int) *Season {
	// no tiene en cuenta partidos de empate
	// esos partidos deberían agregarse manualmente
	season := Season{}

	season.Start = time.Now()
	season.End = nil

	matches := len(combin.Combinations(teams, 2))
	matchesByStage := int(math.Floor(float64(teams) / 2))
	stages := matches / matchesByStage
	// fmt.Printf("Teams: %v - Fechas: %v - Partidos(por fecha): %v\n", teams, stages, matchesByStage)

	for stage := 0; stage < stages; stage++ {
		for match := 0; match < matchesByStage; match++ {
			season.Fixture = append(season.Fixture, Match{
				Stage:      fmt.Sprintf("STAGE_%v", stage),
				StageOrder: stage + 1,
				Number:     match,
				Status:     "CREATED",
			})
		}
	}

	return &season
}

func newSeasonCup(teams int) *Season {
	// no tiene en cuenta los partidos "consuelo", como los del 3er puesto
	// es posible agregarlos a mano, o con algún método especial
	season := Season{}

	season.Start = time.Now()
	season.End = nil

	stages := int(math.Floor(math.Log2(float64(teams))))
	// fmt.Printf("Teams: %v - Fechas: %v - Partidos(por fecha): %v\n", teams, stages, matchesByStage)

	labels := []string{
		"FINAL",
		"SEMI_FINAL",
		"QUARTER_FINAL",
		"LAST_16",
		"LAST_32",
		"LAST_64",
	}

	for stage := stages; stage > 0; stage-- {
		matchesByStage := int(math.Pow(2, float64(stage-1)))
		for match := 1; match <= matchesByStage; match++ {
			season.Fixture = append(season.Fixture, Match{
				Stage:      labels[stage-1],
				StageOrder: 10 * stage, // pongo este número en caso de querer colocar partidos intermedios
				Number:     match,
				Status:     "CREATED",
			})
		}
	}

	return &season
}
