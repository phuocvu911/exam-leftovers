package main

import (
	"fmt"
	"math"
	"sort"
)

type PlayerRating struct {
	PlayerID string
	Rating   int
}

type Match struct {
	PlayerA string
	PlayerB string
	Winner  string
}

func matchMaker(players []string, matches []Match) []PlayerRating {
	// Initialize ratings map
	ratings := make(map[string]float64)
	for _, p := range players {
		ratings[p] = 1000.0 // Default rating
	}

	kFactor := 32.0

	for _, m := range matches { // loop through all matches
		ra := ratings[m.PlayerA]
		rb := ratings[m.PlayerB]

		// Calculate Expected Scores
		// Ea = 1 / (1 + 10^((Rb - Ra) / 400))
		ea := 1.0 / (1.0 + math.Pow(10.0, (rb-ra)/400.0))
		eb := 1.0 / (1.0 + math.Pow(10.0, (ra-rb)/400.0))

		// Determine Actual Scores
		sa := 0.0
		sb := 0.0
		if m.Winner == m.PlayerA {
			sa = 1.0
			sb = 0.0
		} else if m.Winner == m.PlayerB {
			sa = 0.0
			sb = 1.0
		} else {
			sa = 0.5
			sb = 0.5
		}

		// Update Ratings
		ratings[m.PlayerA] = ra + kFactor*(sa-ea)
		ratings[m.PlayerB] = rb + kFactor*(sb-eb)
	}

	// Convert back to slice
	var result []PlayerRating
	for key, value := range ratings {
		subArr := PlayerRating{key, int(math.Round(value))}
		result = append(result, subArr)
	}

	// Sort by rating descending, then by name ascending
	sort.Slice(result, func(i, j int) bool {
		if result[i].Rating != result[j].Rating {
			return result[i].Rating > result[j].Rating
		}
		return result[i].PlayerID < result[j].PlayerID
	})

	return result
}

func main() {
	players := []string{"Alice", "Bob", "Charlie"}
	matches := []Match{
		{PlayerA: "Alice", PlayerB: "Bob", Winner: "Alice"},     // Alice beats Bob
		{PlayerA: "Bob", PlayerB: "Charlie", Winner: "Bob"},     // Bob beats Charlie
		{PlayerA: "Alice", PlayerB: "Charlie", Winner: "Alice"}, // Alice beats Charlie
	}

	results := matchMaker(players, matches)
	fmt.Println(results)
}
