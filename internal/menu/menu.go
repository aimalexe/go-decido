package menu

import (
	"fmt"

	"go-decido/internal/input"
)

const (
	MainCreate = 1
	MainList   = 2
	MainOpen   = 3
	MainExit   = 4

	WorkView        = 1
	WorkCriterion   = 2
	WorkAlternative = 3
	WorkWeights     = 4
	WorkScore       = 5
	WorkResults     = 6
	WorkBack        = 7
)

func ShowMain(decisionCount int) int {
	fmt.Println()
	fmt.Println("============ MENU ============")
	fmt.Println("--- Decisions ---")
	fmt.Println("  1. Create decision")
	fmt.Println("  2. List decisions")
	fmt.Println("  3. Open decision")
	fmt.Println("--- App ---")
	fmt.Println("  4. Exit")
	fmt.Printf("Status: %d decision(s)\n", decisionCount)
	return input.PromptIntInRange("Enter your choice:", 1, 4)
}

func ShowWorkspace(title string) int {
	fmt.Println()
	fmt.Printf("Working on: %q\n", title)
	fmt.Println("  1. View")
	fmt.Println("  2. Add criterion")
	fmt.Println("  3. Add alternative")
	fmt.Println("  4. Set weights")
	fmt.Println("  5. Score alternatives")
	fmt.Println("  6. Calculate results")
	fmt.Println("  7. Back")
	return input.PromptIntInRange("Enter your choice:", 1, 7)
}
