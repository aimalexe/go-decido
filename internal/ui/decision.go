package ui

import (
	"fmt"

	"go-decido/internal/decision"
)

func PrintList(decisions []decision.Decision) {
	Section("Your decisions")
	if len(decisions) == 0 {
		Warningf("No decisions yet. Create one first.")
		return
	}

	rows := make([]Row, 0, len(decisions))
	for i, item := range decisions {
		rows = append(rows, Row{i + 1, item.Title, decisionProgress(item)})
	}

	RenderTable(Table{
		Header: Row{"#", "Decision", "Progress"},
		Rows:   rows,
	})
}

func PrintView(item decision.Decision) {
	Section(fmt.Sprintf("%s  %s", IconDecision, item.Title))
	PrintCriteria(item)
	PrintScorecard(item)
}

func PrintCriteria(item decision.Decision) {
	Section("Criteria and weights")
	if len(item.Criteria) == 0 {
		Warningf("No criteria added yet.")
		return
	}

	rows := make([]Row, 0, len(item.Criteria))
	for _, criterion := range item.Criteria {
		rows = append(rows, Row{
			criterion.ID,
			criterion.Name,
			fmt.Sprintf("%.2f", criterion.Weight),
		})
	}

	RenderTable(Table{
		Header: Row{"ID", "Criterion", "Weight"},
		Rows:   rows,
		Footer: Row{"", "Total weight", fmt.Sprintf("%.2f", item.TotalWeight())},
	})
}

func PrintAlternatives(item decision.Decision) {
	Section("Alternatives")
	if len(item.Alternatives) == 0 {
		Warningf("No alternatives added yet.")
		return
	}

	rows := make([]Row, 0, len(item.Alternatives))
	for _, alternative := range item.Alternatives {
		rows = append(rows, Row{alternative.ID, alternative.Name})
	}

	RenderTable(Table{
		Header: Row{"ID", "Alternative"},
		Rows:   rows,
	})
}

func PrintScorecard(item decision.Decision) {
	Section("Alternatives and scores")
	if len(item.Alternatives) == 0 {
		Warningf("No alternatives added yet.")
		return
	}
	if len(item.Criteria) == 0 {
		rows := make([]Row, 0, len(item.Alternatives))
		for _, alternative := range item.Alternatives {
			rows = append(rows, Row{alternative.ID, alternative.Name})
		}
		RenderTable(Table{
			Header: Row{"ID", "Alternative"},
			Rows:   rows,
		})
		return
	}

	header := Row{"Alternative"}
	for _, criterion := range item.Criteria {
		header = append(header, criterion.Name)
	}

	rows := make([]Row, 0, len(item.Alternatives))
	for _, alternative := range item.Alternatives {
		row := Row{alternative.Name}
		for _, criterion := range item.Criteria {
			score, ok := alternative.ScoreFor(criterion.ID)
			if !ok {
				row = append(row, mutedStyle.Sprint("—"))
				continue
			}
			row = append(row, score)
		}
		rows = append(rows, row)
	}

	RenderTable(Table{Header: header, Rows: rows})
}

func PrintResults(results []decision.Result) {
	if len(results) == 0 {
		Warningf("No results to display.")
		return
	}

	Section(fmt.Sprintf("%s  Results", IconTrophy))
	hasTie := decision.HasTie(results)
	rows := make([]Row, 0, len(results))
	for i, result := range results {
		name := result.AlternativeName
		if i == 0 && !hasTie {
			name = successStyle.Sprintf("%s  %s", IconStar, name)
		}
		rows = append(rows, Row{
			i + 1,
			name,
			fmt.Sprintf("%.2f", result.Score),
		})
	}

	RenderTable(Table{
		Header: Row{"Rank", "Alternative", "Weighted score"},
		Rows:   rows,
	})

	if hasTie {
		Warningf("The top alternatives are tied at %.2f.", results[0].Score)
		return
	}
	Successf("Recommended: %s (%.2f)", results[0].AlternativeName, results[0].Score)
}

func PrintScoringScale(scale string) {
	Infof("Scoring scale: %s", scale)
}

func PrintScoringTarget(name string) {
	Section(fmt.Sprintf("%s  Scoring %s", IconScore, name))
}

func decisionProgress(item decision.Decision) string {
	switch {
	case len(item.Criteria) == 0:
		return mutedStyle.Sprint("Add criteria")
	case len(item.Alternatives) == 0:
		return mutedStyle.Sprint("Add alternatives")
	case !item.IsFullyScored():
		return warningStyle.Sprint("Needs scoring")
	default:
		return successStyle.Sprint("Ready")
	}
}
