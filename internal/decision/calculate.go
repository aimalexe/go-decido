package decision

import (
	"fmt"
	"sort"
)

func (d Decision) ValidateReady() error {
	if len(d.Criteria) == 0 {
		return ErrNoCriteria
	}
	if len(d.Alternatives) == 0 {
		return ErrNoAlternatives
	}
	if d.TotalWeight() == 0 {
		return ErrZeroWeight
	}
	if !d.IsFullyScored() {
		return ErrIncompleteScores
	}
	return nil
}

func (d Decision) CalculateScore(a Alternative) float64 {
	var totalWeight float64
	var weightedScore float64

	for _, criterion := range d.Criteria {
		score, ok := a.ScoreFor(criterion.ID)
		if !ok {
			score = 0
		}
		weightedScore += float64(score) * criterion.Weight
		totalWeight += criterion.Weight
	}

	if totalWeight == 0 {
		return 0
	}

	return weightedScore / totalWeight
}

func (d Decision) CalculateResults() []Result {
	results := make([]Result, 0, len(d.Alternatives))
	for _, alternative := range d.Alternatives {
		results = append(results, Result{
			AlternativeID:   alternative.ID,
			AlternativeName: alternative.Name,
			Score:           d.CalculateScore(alternative),
		})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return results[i].AlternativeName < results[j].AlternativeName
		}
		return results[i].Score > results[j].Score
	})

	return results
}

// HasTie returns true when two or more top-ranked results share the best score.
func HasTie(results []Result) bool {
	if len(results) < 2 {
		return false
	}
	return results[0].Score == results[1].Score
}

func FormatScoreScale() string {
	return fmt.Sprintf("%d = poor … %d = excellent", MinScore, MaxScore)
}
