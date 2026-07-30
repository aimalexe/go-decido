package ui

import (
	"bytes"
	"strings"
	"testing"

	"github.com/fatih/color"
)

func TestMessagesUseConsistentIcons(t *testing.T) {
	var buffer bytes.Buffer
	restoreOutput := SetOutput(&buffer)
	defer restoreOutput()

	previousNoColor := color.NoColor
	color.NoColor = true
	defer func() {
		color.NoColor = previousNoColor
	}()

	Successf("Saved %s", "decision")
	Warningf("Check the score")
	Errorf("Could not save")

	got := buffer.String()
	for _, expected := range []string{
		IconSuccess + " Saved decision",
		IconWarning + " Check the score",
		IconError + " Could not save",
	} {
		if !strings.Contains(got, expected) {
			t.Errorf("output %q does not contain %q", got, expected)
		}
	}
}

func TestRenderTableWritesToConfiguredOutput(t *testing.T) {
	var buffer bytes.Buffer
	restoreOutput := SetOutput(&buffer)
	defer restoreOutput()

	RenderTable(Table{
		Header: Row{"#", "Decision"},
		Rows:   []Row{{1, "Choose a database"}},
	})

	got := buffer.String()
	for _, expected := range []string{"DECISION", "Choose a database"} {
		if !strings.Contains(got, expected) {
			t.Errorf("table output %q does not contain %q", got, expected)
		}
	}
}

func TestDecisionCountStatusUsesSingularAndPluralLabels(t *testing.T) {
	if got := DecisionCountStatus(1); !strings.Contains(got, "1 saved decision") {
		t.Errorf("DecisionCountStatus(1) = %q", got)
	}
	if got := DecisionCountStatus(2); !strings.Contains(got, "2 saved decisions") {
		t.Errorf("DecisionCountStatus(2) = %q", got)
	}
}
