package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(firstTeam string, fTeamScore int,
	secondTeam string, sTeamScore int) {

	if _, ok := l.Teams[firstTeam]; !ok {
		return
	}

	if _, ok := l.Teams[secondTeam]; !ok {
		return
	}

	if fTeamScore == sTeamScore {
		return
	} else if fTeamScore > sTeamScore {
		l.Wins[firstTeam]++
	} else {
		l.Wins[secondTeam]++
	}
}

func (l *League) Ranking() []string {
	rankedTeams := make([]string, 0, len(l.Teams))
	for _, t := range l.Teams {
		rankedTeams = append(rankedTeams, t.Name)
	}

	sort.Slice(rankedTeams, func(i, j int) bool {
		return l.Wins[rankedTeams[i]] > l.Wins[rankedTeams[j]]
	})

	return rankedTeams
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	results := l.Ranking()
	fmt.Println(results)
}
