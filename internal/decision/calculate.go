package decision

import "sort"

func (d Decision) CalculateScore(a Alternative) float64 {
	var totalWeight float64
	var weightedScore float64

	for _, criterion := range d.Criteria {
		score := float64(a.ScoreFor(criterion.ID))
		weightedScore += score * criterion.Weight
		totalWeight += criterion.Weight
	}

	if totalWeight == 0 {
		return 0
	}

	return weightedScore / totalWeight
}

func (d Decision) CalculateResults() []Result {
	results := make([]Result, 0)
	for _, alternative := range d.Alternatives {
		result := Result{
			AlternativeName: alternative.Name,
			Score:           d.CalculateScore(alternative),
		}
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results
}
