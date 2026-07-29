package ui

import (
	"fmt"

	"go-decido/internal/decision"
	"go-decido/internal/input"
)

func PrintBanner() {
	fmt.Println("==============================")
	fmt.Println("         GO-DECIDO")
	fmt.Println("==============================")
	fmt.Println("Weigh options. Decide with clarity.")
}

func PrintSuccess(msg string) {
	fmt.Println(msg)
}

func PrintError(err error) {
	fmt.Printf("Error: %v\n", err)
}

func PrintList(decisions []decision.Decision) {
	if len(decisions) == 0 {
		fmt.Println("No decisions yet. Create one first.")
		return
	}
	fmt.Println()
	fmt.Println("Decisions:")
	for i, d := range decisions {
		fmt.Printf("  %d. %s\n", i+1, d.Title)
	}
}

// SelectDecision asks the user to pick from the list (1-based index).
// Returns the decision ID and true, or 0 and false if cancelled/empty.
func SelectDecision(decisions []decision.Decision) (id int, ok bool) {
	if len(decisions) == 0 {
		fmt.Println("No decisions yet. Create one first.")
		return 0, false
	}
	PrintList(decisions)
	choice := input.PromptChoice("Enter decision number (0 = cancel):", 0, len(decisions))
	if choice == 0 {
		fmt.Println("Cancelled.")
		return 0, false
	}
	return decisions[choice-1].ID, true
}

func PrintView(d decision.Decision) {
	fmt.Println()
	fmt.Printf("Decision: %s\n", d.Title)
	fmt.Println()

	fmt.Println("Criteria:")
	if len(d.Criteria) == 0 {
		fmt.Println("  (none)")
	} else {
		fmt.Printf("  %-4s %-18s %s\n", "ID", "Name", "Weight")
		for _, c := range d.Criteria {
			fmt.Printf("  %-4d %-18s %.2f\n", c.ID, c.Name, c.Weight)
		}
		fmt.Printf("  Total weight: %.2f\n", d.TotalWeight())
	}

	fmt.Println()
	fmt.Println("Alternatives / scores:")
	if len(d.Alternatives) == 0 {
		fmt.Println("  (none)")
		return
	}

	if len(d.Criteria) == 0 {
		for _, a := range d.Alternatives {
			fmt.Printf("  - %s\n", a.Name)
		}
		return
	}

	fmt.Print("  ")
	fmt.Printf("%-14s", "Name")
	for _, c := range d.Criteria {
		name := c.Name
		if len(name) > 10 {
			name = name[:10]
		}
		fmt.Printf(" %10s", name)
	}
	fmt.Println()

	for _, a := range d.Alternatives {
		fmt.Printf("  %-14s", a.Name)
		for _, c := range d.Criteria {
			if score, ok := a.ScoreFor(c.ID); ok {
				fmt.Printf(" %10d", score)
			} else {
				fmt.Printf(" %10s", "-")
			}
		}
		fmt.Println()
	}
}

func PrintCriteria(d decision.Decision) {
	fmt.Println()
	fmt.Println("Criteria:")
	for _, c := range d.Criteria {
		fmt.Printf("  %d. %-18s Weight: %.2f\n", c.ID, c.Name, c.Weight)
	}
	fmt.Printf("Total weight: %.2f\n", d.TotalWeight())
}

func PrintAlternatives(d decision.Decision) {
	fmt.Println()
	fmt.Println("Alternatives:")
	for i, a := range d.Alternatives {
		fmt.Printf("  %d. %s\n", i+1, a.Name)
	}
}

func PrintResults(results []decision.Result) {
	fmt.Println()
	fmt.Println("Results (weighted average, 1–5 scale):")
	fmt.Printf("  %-4s %-20s %s\n", "Rank", "Alternative", "Score")
	for i, r := range results {
		fmt.Printf("  %-4d %-20s %.2f\n", i+1, r.AlternativeName, r.Score)
	}
	fmt.Println()
	if len(results) == 0 {
		return
	}
	if decision.HasTie(results) {
		fmt.Printf("Tie for top place at %.2f — consider refining weights or scores.\n", results[0].Score)
		return
	}
	fmt.Printf("Recommended: %s (%.2f)\n", results[0].AlternativeName, results[0].Score)
}
