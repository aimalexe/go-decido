package ui

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func terminalTableStyle() table.Style {
	style := table.StyleLight
	style.Options.DrawBorder = true
	style.Options.SeparateColumns = true
	style.Options.SeparateHeader = true
	style.Options.SeparateRows = false
	style.Color.Border = text.Colors{text.FgHiBlack}
	style.Color.Header = text.Colors{text.FgHiCyan, text.Bold}
	style.Color.Row = text.Colors{text.FgWhite}
	style.Color.RowAlternate = text.Colors{text.FgHiWhite}
	style.Color.Footer = text.Colors{text.FgHiCyan, text.Bold}
	return style
}
