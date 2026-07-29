package app

import (
	"fmt"
	"path/filepath"

	"go-decido/internal/decision"
	"go-decido/internal/input"
	"go-decido/internal/menu"
	"go-decido/internal/service"
	"go-decido/internal/ui"
)

func Run() {
	store, err := service.NewJSONFileStore(filepath.Join("data", "decisions.json"))
	if err != nil {
		ui.PrintError(fmt.Errorf("load saved decisions: %w", err))
		return
	}
	svc := service.NewDecisionService(store)
	ui.PrintBanner()

	for {
		choice := menu.ShowMain(svc.Count())
		switch choice {
		case menu.MainCreate:
			createDecision(svc)
		case menu.MainList:
			ui.PrintList(svc.List())
		case menu.MainOpen:
			openWorkspace(svc)
		case menu.MainExit:
			if input.PromptYesNo("Exit go-decido? (y/n):") {
				fmt.Println("Thank you for using go-decido. Goodbye!")
				return
			}
		}
	}
}

func createDecision(svc *service.DecisionService) {
	title := input.PromptNonEmpty("What decision do you want to make?")
	d, err := svc.Create(title)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.PrintSuccess(fmt.Sprintf("Created decision %q. Open it to add criteria and alternatives.", d.Title))
}

func openWorkspace(svc *service.DecisionService) {
	id, ok := ui.SelectDecision(svc.List())
	if !ok {
		return
	}

	for {
		d, err := svc.Get(id)
		if err != nil {
			ui.PrintError(err)
			return
		}

		choice := menu.ShowWorkspace(d.Title)
		switch choice {
		case menu.WorkView:
			ui.PrintView(d)
		case menu.WorkCriterion:
			addCriterion(svc, id)
		case menu.WorkAlternative:
			addAlternative(svc, id)
		case menu.WorkWeights:
			setWeights(svc, id)
		case menu.WorkScore:
			scoreAlternatives(svc, id)
		case menu.WorkResults:
			showResults(svc, id)
		case menu.WorkBack:
			return
		}
	}
}

func addCriterion(svc *service.DecisionService, decisionID int) {
	name := input.PromptNonEmpty("Criterion name:")
	d, err := svc.AddCriterion(decisionID, name)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.PrintSuccess(fmt.Sprintf("Added criterion %q to %q.", name, d.Title))
}

func addAlternative(svc *service.DecisionService, decisionID int) {
	name := input.PromptNonEmpty("Alternative name:")
	d, err := svc.AddAlternative(decisionID, name)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.PrintSuccess(fmt.Sprintf("Added alternative %q to %q.", name, d.Title))
}

func setWeights(svc *service.DecisionService, decisionID int) {
	for {
		d, err := svc.Get(decisionID)
		if err != nil {
			ui.PrintError(err)
			return
		}
		if len(d.Criteria) == 0 {
			fmt.Println("Please add criteria first.")
			return
		}

		ui.PrintCriteria(d)
		ids := make(map[int]struct{}, len(d.Criteria))
		for _, c := range d.Criteria {
			ids[c.ID] = struct{}{}
		}

		fmt.Print("Criterion ID (0 = done): ")
		criterionID := input.ReadInt()
		if criterionID == 0 {
			return
		}
		if _, exists := ids[criterionID]; !exists {
			fmt.Printf("No criterion with ID %d.\n", criterionID)
			continue
		}

		weight := input.PromptFloatNonNegative("Weight:")
		updated, err := svc.SetWeight(decisionID, criterionID, weight)
		if err != nil {
			ui.PrintError(err)
			continue
		}
		ui.PrintSuccess(fmt.Sprintf("Updated weight. Total weight: %.2f", updated.TotalWeight()))

		if !input.PromptYesNo("Set another weight? (y/n):") {
			return
		}
	}
}

func scoreAlternatives(svc *service.DecisionService, decisionID int) {
	d, err := svc.Get(decisionID)
	if err != nil {
		ui.PrintError(err)
		return
	}
	if len(d.Criteria) == 0 {
		fmt.Println("Please add criteria first.")
		return
	}
	if len(d.Alternatives) == 0 {
		fmt.Println("Please add alternatives first.")
		return
	}

	fmt.Printf("Scoring scale: %s\n", decision.FormatScoreScale())
	ui.PrintAlternatives(d)

	choice := input.PromptChoice(
		fmt.Sprintf("Score which? (1-%d, or 0 = all):", len(d.Alternatives)),
		0,
		len(d.Alternatives),
	)

	var targets []decision.Alternative
	if choice == 0 {
		targets = d.Alternatives
	} else {
		targets = []decision.Alternative{d.Alternatives[choice-1]}
	}

	for _, alt := range targets {
		fmt.Printf("\nScoring: %s\n", alt.Name)
		for _, c := range d.Criteria {
			label := fmt.Sprintf("  %s (%d-%d):", c.Name, decision.MinScore, decision.MaxScore)
			score := input.PromptIntInRange(label, decision.MinScore, decision.MaxScore)
			if _, err := svc.SetScore(decisionID, alt.ID, c.ID, score); err != nil {
				ui.PrintError(err)
				return
			}
		}
		ui.PrintSuccess(fmt.Sprintf("Saved scores for %q.", alt.Name))
	}
}

func showResults(svc *service.DecisionService, decisionID int) {
	results, err := svc.Results(decisionID)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.PrintResults(results)
}
