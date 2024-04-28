package domain

import (
	"fmt"
	"testing"

	lo "github.com/samber/lo"
)

func TestCreateLeague(t *testing.T) {
	type test struct {
		teams          int
		stages         int
		matchesByStage int
	}

	tests := []test{
		{teams: 10, stages: 9, matchesByStage: 5},
		{teams: 5, stages: 5, matchesByStage: 2},
		{teams: 4, stages: 3, matchesByStage: 2},
		{teams: 3, stages: 3, matchesByStage: 1},
	}

	for _, item := range tests {
		builder := NewCompetitionBuilder()
		builder = builder.CreateLeague(item.teams)
		competition := builder.Build()
		fmt.Println(competition.Type)

		stages := lo.GroupBy(competition.CurrentSeason.Fixture, func(m Match) string {
			return m.Stage
		})

		if len(stages) != item.stages {
			t.Errorf("Competition with teams=%v should have %v stages, and it returns %v", item.teams, item.stages, len(stages))
		}

		for stage := range stages {
			if matchesByStage := len(stages[stage]); matchesByStage != item.matchesByStage {
				t.Errorf("Should be equals %v not equals to %v", matchesByStage, item.matchesByStage)
			}
		}

		if competition.Type != "LEAGUE" {
			t.Errorf("Competition should be LEAGUE but it is %s\n", competition.Type)
		}
	}
}

func TestCreateCup(t *testing.T) {
	// hacer un parametric test en lugar de table driven test

	type test struct {
		teams   int
		stages  int
		matches int
	}

	// stages = log2(teams)
	// matches = teams - 1
	tests := []test{
		{teams: 32, stages: 5, matches: 31},
		{teams: 16, stages: 4, matches: 15},
		{teams: 8, stages: 3, matches: 7},
		// probar con impares
		// probar con números más grande
	}

	for _, item := range tests {
		builder := NewCompetitionBuilder()
		builder = builder.CreateCup(item.teams)
		competition := builder.Build()
		fmt.Println(competition.Type)

		if len(competition.CurrentSeason.Fixture) != item.matches {
			t.Errorf("Competition with %v teams should have %v matches", item.teams, item.matches)
		}
		stages := lo.GroupBy(competition.CurrentSeason.Fixture, func(m Match) string {
			return m.Stage
		})

		if len(stages) != item.stages {
			t.Errorf("Competition with %v teams should have %v stages", item.teams, item.stages)
		}

		if competition.Type != "CUP" {
			t.Errorf("Competition should be CUP but it is %s\n", competition.Type)
		}
	}
}
