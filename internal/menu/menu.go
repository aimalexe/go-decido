package menu

import (
	"go-decido/internal/input"
	"go-decido/internal/ui"
)

const (
	MainCreate = 1
	MainList   = 2
	MainOpen   = 3
	MainRename = 4
	MainDelete = 5
	MainExit   = 6

	WorkView         = 1
	WorkCriteria     = 2
	WorkAlternatives = 3
	WorkWeights      = 4
	WorkScore        = 5
	WorkResults      = 6
	WorkRename       = 7
	WorkDelete       = 8
	WorkBack         = 9

	ItemAdd    = 1
	ItemRename = 2
	ItemDelete = 3
	ItemBack   = 4
)

func ShowMain(decisionCount int) int {
	ui.PrintMenu("MAIN MENU", "", []ui.MenuSection{
		{
			Title: "DECISIONS",
			Items: []ui.MenuItem{
				{Number: MainCreate, Icon: ui.IconDecision, Label: "Create decision"},
				{Number: MainList, Icon: ui.IconFolder, Label: "List decisions"},
				{Number: MainOpen, Icon: ui.IconArrow, Label: "Open decision"},
				{Number: MainRename, Icon: ui.IconEdit, Label: "Rename decision"},
				{Number: MainDelete, Icon: ui.IconDelete, Label: "Delete decision"},
			},
		},
		{
			Title: "APP",
			Items: []ui.MenuItem{
				{Number: MainExit, Icon: ui.IconExit, Label: "Exit"},
			},
		},
	}, ui.DecisionCountStatus(decisionCount))
	return input.PromptIntInRange("Enter your choice:", 1, 6)
}

func ShowWorkspace(title string) int {
	ui.PrintMenu("DECISION WORKSPACE", "Working on: "+title, []ui.MenuSection{
		{
			Title: "BUILD",
			Items: []ui.MenuItem{
				{Number: WorkView, Icon: ui.IconDecision, Label: "View overview"},
				{Number: WorkCriteria, Icon: ui.IconCriterion, Label: "Manage criteria"},
				{Number: WorkAlternatives, Icon: ui.IconAlternative, Label: "Manage alternatives"},
				{Number: WorkWeights, Icon: ui.IconStar, Label: "Set weights"},
			},
		},
		{
			Title: "EVALUATE",
			Items: []ui.MenuItem{
				{Number: WorkScore, Icon: ui.IconScore, Label: "Score alternatives"},
				{Number: WorkResults, Icon: ui.IconTrophy, Label: "Calculate results"},
			},
		},
		{
			Title: "MANAGE",
			Items: []ui.MenuItem{
				{Number: WorkRename, Icon: ui.IconEdit, Label: "Rename this decision"},
				{Number: WorkDelete, Icon: ui.IconDelete, Label: "Delete this decision"},
				{Number: WorkBack, Icon: ui.IconBack, Label: "Back"},
			},
		},
	}, "")
	return input.PromptIntInRange("Enter your choice:", 1, 9)
}

func ShowCriteria() int {
	ui.PrintMenu("CRITERIA", "", []ui.MenuSection{
		{
			Title: "ACTIONS",
			Items: []ui.MenuItem{
				{Number: ItemAdd, Icon: ui.IconCriterion, Label: "Add criterion"},
				{Number: ItemRename, Icon: ui.IconEdit, Label: "Rename criterion"},
				{Number: ItemDelete, Icon: ui.IconDelete, Label: "Delete criterion"},
				{Number: ItemBack, Icon: ui.IconBack, Label: "Back"},
			},
		},
	}, "")
	return input.PromptIntInRange("Enter your choice:", 1, 4)
}

func ShowAlternatives() int {
	ui.PrintMenu("ALTERNATIVES", "", []ui.MenuSection{
		{
			Title: "ACTIONS",
			Items: []ui.MenuItem{
				{Number: ItemAdd, Icon: ui.IconAlternative, Label: "Add alternative"},
				{Number: ItemRename, Icon: ui.IconEdit, Label: "Rename alternative"},
				{Number: ItemDelete, Icon: ui.IconDelete, Label: "Delete alternative"},
				{Number: ItemBack, Icon: ui.IconBack, Label: "Back"},
			},
		},
	}, "")
	return input.PromptIntInRange("Enter your choice:", 1, 4)
}
