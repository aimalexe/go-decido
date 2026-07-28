package app

import (
	"fmt"

	"go-decido/internal/decision"
	"go-decido/internal/input"
	"go-decido/internal/menu"
	"go-decido/internal/service"
)

func Run() {
	decisionService := service.NewDecisionService()

	fmt.Println("==============================")
	fmt.Println("     GO-DECIDO")
	fmt.Println("==============================")
	fmt.Println()

	for {
		choice := menu.Show()

		switch choice {

		case 1:
			newDecision := decision.Create()
			decisionService.AddDecision(newDecision)

		case 2:
			decision.List(decisionService.Decisions())

		case 3:
			index, ok := decision.Select(decisionService.Decisions())
			if !ok {
				break
			}
			selectedDecision, _ := decisionService.GetDecision(index)
			decision.View(selectedDecision)

		case 4:
			index, ok := decision.Select(decisionService.Decisions())
			if !ok {
				break
			}
			selectedDecision, _ := decisionService.GetDecision(index)
			fmt.Println("Criterion name: ")
			name := input.ReadString()
			decision.AddCriterion(&selectedDecision, name)
			decisionService.UpdateDecision(index, selectedDecision)

		case 5:
			index, ok := decision.Select(decisionService.Decisions())
			if !ok {
				break
			}
			selectedDecision, _ := decisionService.GetDecision(index)
			fmt.Println("Alternative name: ")
			alternativeName := input.ReadString()
			decision.AddAlternative(&selectedDecision, alternativeName)
			decisionService.UpdateDecision(index, selectedDecision)

		case 6:
			index, ok := decision.Select(decisionService.Decisions())
			if !ok {
				break
			}

			selectedDecision, _ := decisionService.GetDecision(index)
			if len(selectedDecision.Criteria) == 0 {
				fmt.Println("Please add criteria first.")
				break
			}

			if len(selectedDecision.Alternatives) == 0 {
				fmt.Println("Please add alternatives first.")
				break
			}

			fmt.Println()
			fmt.Println("Select Alternative")

			for i, alt := range selectedDecision.Alternatives {
				fmt.Printf("%d. %s\n", i+1, alt.Name)
			}

			fmt.Print("Choice: ")
			altIndex := input.ReadInt()
			if altIndex < 1 || altIndex > len(selectedDecision.Alternatives) {
				fmt.Println("Invalid choice.")
				break
			}

			decision.ScoreAlternative(
				&selectedDecision.Alternatives[altIndex-1],
				selectedDecision.Criteria,
			)
			decisionService.UpdateDecision(index, selectedDecision)

		case 7:
			index, ok := decision.Select(decisionService.Decisions())
			if !ok {
				break
			}
			selectedDecision, _ := decisionService.GetDecision(index)
			if len(selectedDecision.Criteria) == 0 {
				fmt.Println("No criteria available.")
				break
			}
			fmt.Println()
			for _, criterion := range selectedDecision.Criteria {
				fmt.Printf("%d. %s\n",
					criterion.ID,
					criterion.Name,
				)
			}
			fmt.Print("Criterion: ")
			criterionID := input.ReadInt()
			fmt.Print("Weight: ")
			weight := input.ReadFloat()
			decision.SetWeight(
				&selectedDecision,
				criterionID,
				weight,
			)
			decisionService.UpdateDecision(index, selectedDecision)

		case 8:
			index, ok := decision.Select(
				decisionService.Decisions(),
			)
			if !ok {
				break
			}
			selectedDecision, _ := decisionService.GetDecision(index)
			results := selectedDecision.CalculateResults()
			decision.ShowResults(results)

		case 9:
			fmt.Println("Thank you for using go-decido. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}

		fmt.Println()
	}
}
