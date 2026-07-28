package decision

import (
	"fmt"
)

func List(decisions []Decision) {
	if len(decisions) == 0 {
		fmt.Println("No decisions available.")
		return
	}

	fmt.Println()
	fmt.Println("Your Decisions:")

	for i, decision := range decisions {
		fmt.Printf("\t%d. %s\n", i+1, decision.Title)
	}
}
