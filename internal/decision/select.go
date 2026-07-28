package decision

import (
	"go-decido/internal/input"
	"fmt"
)

func Select(decisions []Decision) (int, bool) {
	if len(decisions) == 0 {
		fmt.Println("No decisions found.")
		return 0, false
	}

	List(decisions)

	fmt.Println()
	fmt.Print("> Enter decision number: ")

	choice := input.ReadInt()

	if choice < 1 || choice > len(decisions) {
		fmt.Println("Invalid decision number.")
		return 0, false
	}

	return choice - 1, true
}
