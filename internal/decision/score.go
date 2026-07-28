package decision

import (
	"fmt"

	"go-decido/internal/input"
)

func ScoreAlternative(alternative *Alternative, criteria []Criterion) {

	fmt.Println()
	fmt.Println("Scoring:", alternative.Name)

	for _, criterion := range criteria {

		fmt.Printf("%s (1-5): ", criterion.Name)

		var score int
		for {
			score = input.ReadInt()

			if score >= 1 && score <= 5 {
				break
			}

			fmt.Println("Please enter a score between 1 and 5.")
		}

		alternative.SetScore(criterion.ID, score)
	}
}
