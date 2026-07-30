package ui

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Row is a UI-owned table row, so callers do not depend on go-pretty.
type Row []any

// Table describes a reusable terminal table.
type Table struct {
	Header Row
	Rows   []Row
	Footer Row
}

// RenderTable renders consistently styled tabular output.
func RenderTable(data Table) {
	writer := table.NewWriter()
	writer.SetOutputMirror(output)
	writer.SetStyle(terminalTableStyle())
	writer.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
	})

	if len(data.Header) > 0 {
		writer.AppendHeader(toTableRow(data.Header))
	}
	for _, row := range data.Rows {
		writer.AppendRow(toTableRow(row))
	}
	if len(data.Footer) > 0 {
		writer.AppendFooter(toTableRow(data.Footer))
	}

	writer.Render()
}

func toTableRow(row Row) table.Row {
	return table.Row(row)
}
