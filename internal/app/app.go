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
		ui.Errorf("Could not load saved decisions: %v", err)
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
		case menu.MainRename:
			renameDecision(svc)
		case menu.MainDelete:
			deleteDecision(svc)
		case menu.MainExit:
			if input.PromptYesNo("Exit go-decido? (y/n):") {
				ui.PrintGoodbye()
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
	ui.Successf("Created decision %q. Open it to add criteria and alternatives.", d.Title)
}

func renameDecision(svc *service.DecisionService) {
	id, ok := selectDecision(svc.List())
	if !ok {
		return
	}
	renameDecisionByID(svc, id)
}

func deleteDecision(svc *service.DecisionService) {
	id, ok := selectDecision(svc.List())
	if !ok {
		return
	}
	deleteDecisionByID(svc, id)
}

func openWorkspace(svc *service.DecisionService) {
	id, ok := selectDecision(svc.List())
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
		case menu.WorkCriteria:
			manageCriteria(svc, id)
		case menu.WorkAlternatives:
			manageAlternatives(svc, id)
		case menu.WorkWeights:
			setWeights(svc, id)
		case menu.WorkScore:
			scoreAlternatives(svc, id)
		case menu.WorkResults:
			showResults(svc, id)
		case menu.WorkRename:
			if !renameDecisionByID(svc, id) {
				return
			}
		case menu.WorkDelete:
			if deleteDecisionByID(svc, id) {
				return
			}
		case menu.WorkBack:
			return
		}
	}
}

func manageCriteria(svc *service.DecisionService, decisionID int) {
	for {
		switch menu.ShowCriteria() {
		case menu.ItemAdd:
			addCriterion(svc, decisionID)
		case menu.ItemRename:
			renameCriterion(svc, decisionID)
		case menu.ItemDelete:
			deleteCriterion(svc, decisionID)
		case menu.ItemBack:
			return
		}
	}
}

func manageAlternatives(svc *service.DecisionService, decisionID int) {
	for {
		switch menu.ShowAlternatives() {
		case menu.ItemAdd:
			addAlternative(svc, decisionID)
		case menu.ItemRename:
			renameAlternative(svc, decisionID)
		case menu.ItemDelete:
			deleteAlternative(svc, decisionID)
		case menu.ItemBack:
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
	ui.Successf("Added criterion %q to %q.", name, d.Title)
}

func renameCriterion(svc *service.DecisionService, decisionID int) {
	d, err := svc.Get(decisionID)
	if err != nil {
		ui.PrintError(err)
		return
	}
	if len(d.Criteria) == 0 {
		ui.Warningf("Please add criteria first.")
		return
	}

	criterionID, ok := selectCriterionID(d)
	if !ok {
		return
	}

	name := input.PromptNonEmpty("New criterion name:")
	updated, err := svc.RenameCriterion(decisionID, criterionID, name)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.Successf("Renamed criterion to %q in %q.", name, updated.Title)
}

func deleteCriterion(svc *service.DecisionService, decisionID int) {
	d, err := svc.Get(decisionID)
	if err != nil {
		ui.PrintError(err)
		return
	}
	if len(d.Criteria) == 0 {
		ui.Warningf("Please add criteria first.")
		return
	}

	criterionID, ok := selectCriterionID(d)
	if !ok {
		return
	}

	criterion := d.FindCriterion(criterionID)
	if criterion == nil {
		ui.Warningf("No criterion with ID %d.", criterionID)
		return
	}
	if !input.PromptYesNo(fmt.Sprintf("Delete criterion %q? (y/n):", criterion.Name)) {
		ui.Infof("Deletion cancelled.")
		return
	}

	updated, err := svc.DeleteCriterion(decisionID, criterionID)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.Successf("Deleted criterion %q from %q.", criterion.Name, updated.Title)
}

func addAlternative(svc *service.DecisionService, decisionID int) {
	name := input.PromptNonEmpty("Alternative name:")
	d, err := svc.AddAlternative(decisionID, name)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.Successf("Added alternative %q to %q.", name, d.Title)
}

func renameAlternative(svc *service.DecisionService, decisionID int) {
	d, err := svc.Get(decisionID)
	if err != nil {
		ui.PrintError(err)
		return
	}
	if len(d.Alternatives) == 0 {
		ui.Warningf("Please add alternatives first.")
		return
	}

	alternativeID, ok := selectAlternativeID(d)
	if !ok {
		return
	}

	name := input.PromptNonEmpty("New alternative name:")
	updated, err := svc.RenameAlternative(decisionID, alternativeID, name)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.Successf("Renamed alternative to %q in %q.", name, updated.Title)
}

func deleteAlternative(svc *service.DecisionService, decisionID int) {
	d, err := svc.Get(decisionID)
	if err != nil {
		ui.PrintError(err)
		return
	}
	if len(d.Alternatives) == 0 {
		ui.Warningf("Please add alternatives first.")
		return
	}

	alternativeID, ok := selectAlternativeID(d)
	if !ok {
		return
	}

	alternative := d.FindAlternative(alternativeID)
	if alternative == nil {
		ui.Warningf("No alternative with ID %d.", alternativeID)
		return
	}
	if !input.PromptYesNo(fmt.Sprintf("Delete alternative %q? (y/n):", alternative.Name)) {
		ui.Infof("Deletion cancelled.")
		return
	}

	updated, err := svc.DeleteAlternative(decisionID, alternativeID)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.Successf("Deleted alternative %q from %q.", alternative.Name, updated.Title)
}

func renameDecisionByID(svc *service.DecisionService, id int) bool {
	d, err := svc.Get(id)
	if err != nil {
		ui.PrintError(err)
		return false
	}

	title := input.PromptNonEmpty(fmt.Sprintf("New title for %q:", d.Title))
	updated, err := svc.Rename(id, title)
	if err != nil {
		ui.PrintError(err)
		return false
	}
	ui.Successf("Renamed decision to %q.", updated.Title)
	return true
}

func deleteDecisionByID(svc *service.DecisionService, id int) bool {
	d, err := svc.Get(id)
	if err != nil {
		ui.PrintError(err)
		return false
	}
	if !input.PromptYesNo(fmt.Sprintf("Delete decision %q? This cannot be undone. (y/n):", d.Title)) {
		ui.Infof("Deletion cancelled.")
		return false
	}
	if err := svc.Delete(id); err != nil {
		ui.PrintError(err)
		return false
	}
	ui.Successf("Deleted decision %q.", d.Title)
	return true
}

func setWeights(svc *service.DecisionService, decisionID int) {
	for {
		d, err := svc.Get(decisionID)
		if err != nil {
			ui.PrintError(err)
			return
		}
		if len(d.Criteria) == 0 {
			ui.Warningf("Please add criteria first.")
			return
		}

		criterionID, ok := selectCriterionID(d)
		if !ok {
			return
		}

		weight := input.PromptFloatNonNegative("Weight:")
		updated, err := svc.SetWeight(decisionID, criterionID, weight)
		if err != nil {
			ui.PrintError(err)
			continue
		}
		ui.Successf("Updated weight. Total weight: %.2f", updated.TotalWeight())

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
		ui.Warningf("Please add criteria first.")
		return
	}
	if len(d.Alternatives) == 0 {
		ui.Warningf("Please add alternatives first.")
		return
	}

	ui.PrintScoringScale(decision.FormatScoreScale())
	ui.PrintAlternatives(d)

	ids := alternativeIDs(d)
	ui.Promptf("Score which? (enter ID, or 0 = all):")
	choice := input.ReadInt()

	var targets []decision.Alternative
	if choice == 0 {
		targets = d.Alternatives
	} else {
		if _, exists := ids[choice]; !exists {
			ui.Warningf("No alternative with ID %d.", choice)
			return
		}
		targets = []decision.Alternative{*d.FindAlternative(choice)}
	}

	for _, alt := range targets {
		ui.PrintScoringTarget(alt.Name)
		for _, c := range d.Criteria {
			label := fmt.Sprintf("  %s (%d-%d):", c.Name, decision.MinScore, decision.MaxScore)
			score := input.PromptIntInRange(label, decision.MinScore, decision.MaxScore)
			if _, err := svc.SetScore(decisionID, alt.ID, c.ID, score); err != nil {
				ui.PrintError(err)
				return
			}
		}
		ui.Successf("Saved scores for %q.", alt.Name)
	}
}

func selectDecision(decisions []decision.Decision) (id int, ok bool) {
	if len(decisions) == 0 {
		ui.Warningf("No decisions yet. Create one first.")
		return 0, false
	}

	ui.PrintList(decisions)
	choice := input.PromptChoice("Enter decision number (0 = cancel):", 0, len(decisions))
	if choice == 0 {
		ui.Infof("Selection cancelled.")
		return 0, false
	}
	return decisions[choice-1].ID, true
}

func selectCriterionID(d decision.Decision) (id int, ok bool) {
	ui.PrintCriteria(d)
	ids := criterionIDs(d)
	ui.Promptf("Criterion ID (0 = cancel):")
	criterionID := input.ReadInt()
	if criterionID == 0 {
		ui.Infof("Selection cancelled.")
		return 0, false
	}
	if _, exists := ids[criterionID]; !exists {
		ui.Warningf("No criterion with ID %d.", criterionID)
		return 0, false
	}
	return criterionID, true
}

func selectAlternativeID(d decision.Decision) (id int, ok bool) {
	ui.PrintAlternatives(d)
	ids := alternativeIDs(d)
	ui.Promptf("Alternative ID (0 = cancel):")
	alternativeID := input.ReadInt()
	if alternativeID == 0 {
		ui.Infof("Selection cancelled.")
		return 0, false
	}
	if _, exists := ids[alternativeID]; !exists {
		ui.Warningf("No alternative with ID %d.", alternativeID)
		return 0, false
	}
	return alternativeID, true
}

func showResults(svc *service.DecisionService, decisionID int) {
	results, err := svc.Results(decisionID)
	if err != nil {
		ui.PrintError(err)
		return
	}
	ui.PrintResults(results)
}

func criterionIDs(d decision.Decision) map[int]struct{} {
	ids := make(map[int]struct{}, len(d.Criteria))
	for _, c := range d.Criteria {
		ids[c.ID] = struct{}{}
	}
	return ids
}

func alternativeIDs(d decision.Decision) map[int]struct{} {
	ids := make(map[int]struct{}, len(d.Alternatives))
	for _, a := range d.Alternatives {
		ids[a.ID] = struct{}{}
	}
	return ids
}
