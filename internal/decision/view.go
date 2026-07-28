package decision

import (
	"fmt"
)

func View(decision Decision) {
	fmt.Println()
	fmt.Println("Decision:")
	fmt.Println(decision.Title)

	fmt.Println("\nAlternatives:")

	if len(decision.Alternatives) == 0 {
		fmt.Println("  None")
	} else {

		for _, alternative := range decision.Alternatives {

			fmt.Println()
			fmt.Println(alternative.Name)

			if len(alternative.Scores) == 0 {
				fmt.Println("  No scores")
				continue
			}

			for _, criterion := range decision.Criteria {
				fmt.Printf(
					"%d. %-15s Weight: %.2f\n",
					criterion.ID,
					criterion.Name,
					criterion.Weight,
				)
			}
		}
	}
}
