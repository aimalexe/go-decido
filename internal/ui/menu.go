package ui

import "fmt"

// MenuItem is one numbered action in a menu.
type MenuItem struct {
	Number int
	Icon   string
	Label  string
}

// MenuSection groups related actions.
type MenuSection struct {
	Title string
	Items []MenuItem
}

// PrintMenu renders a menu with consistent spacing, colors, and icons.
func PrintMenu(title, context string, sections []MenuSection, status string) {
	Blank()
	primaryStyle.Fprintf(output, "%s  %s\n", IconMenu, title)
	if context != "" {
		highlightStyle.Fprintf(output, "   %s\n", context)
	}

	for _, section := range sections {
		mutedStyle.Fprintf(output, "\n   %s\n", section.Title)
		for _, item := range section.Items {
			Printf("   ")
			infoStyle.Fprintf(output, "%d", item.Number)
			Printf("  %s  %s\n", item.Icon, item.Label)
		}
	}

	if status != "" {
		mutedStyle.Fprintf(output, "\n   %s\n", status)
	}
	Blank()
}

func DecisionCountStatus(count int) string {
	label := "decisions"
	if count == 1 {
		label = "decision"
	}
	return fmt.Sprintf("%s  %d saved %s", IconInfo, count, label)
}
