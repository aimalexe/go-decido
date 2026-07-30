package menu

import (
	"go-decido/internal/input"
	"go-decido/internal/ui"
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
	ui.PrintMenu("MAIN MENU", "", []ui.MenuSection{
		{
			Title: "DECISIONS",
			Items: []ui.MenuItem{
				{Number: MainCreate, Icon: ui.IconDecision, Label: "Create decision"},
				{Number: MainList, Icon: ui.IconFolder, Label: "List decisions"},
				{Number: MainOpen, Icon: ui.IconArrow, Label: "Open decision"},
			},
		},
		{
			Title: "APP",
			Items: []ui.MenuItem{
				{Number: MainExit, Icon: ui.IconExit, Label: "Exit"},
			},
		},
	}, ui.DecisionCountStatus(decisionCount))
	return input.PromptIntInRange("Enter your choice:", 1, 4)
}

func ShowWorkspace(title string) int {
	ui.PrintMenu("DECISION WORKSPACE", "Working on: "+title, []ui.MenuSection{
		{
			Title: "BUILD",
			Items: []ui.MenuItem{
				{Number: WorkView, Icon: ui.IconDecision, Label: "View overview"},
				{Number: WorkCriterion, Icon: ui.IconCriterion, Label: "Add criterion"},
				{Number: WorkAlternative, Icon: ui.IconAlternative, Label: "Add alternative"},
				{Number: WorkWeights, Icon: ui.IconStar, Label: "Set weights"},
			},
		},
		{
			Title: "EVALUATE",
			Items: []ui.MenuItem{
				{Number: WorkScore, Icon: ui.IconScore, Label: "Score alternatives"},
				{Number: WorkResults, Icon: ui.IconTrophy, Label: "Calculate results"},
				{Number: WorkBack, Icon: ui.IconBack, Label: "Back"},
			},
		},
	}, "")
	return input.PromptIntInRange("Enter your choice:", 1, 7)
}
