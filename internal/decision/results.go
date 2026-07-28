package decision

import "fmt"

func ShowResults(results []Result) {
	if len(results) == 0 {
		fmt.Println("No results available.")
		return
	}

	fmt.Println()
	fmt.Println("========== Results ==========")

	for i, result := range results {
		fmt.Printf(
			"%d. %-20s %.2f\n",
			i+1,
			result.AlternativeName,
			result.Score,
		)
	}

	fmt.Println()
	fmt.Println("Recommended:")
	fmt.Println(results[0].AlternativeName)
}
