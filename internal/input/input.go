package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"go-decido/internal/ui"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ReadInt() int {
	for {
		text := ReadString()
		number, err := strconv.Atoi(text)
		if err == nil {
			return number
		}
		ui.Warningf("Please enter a valid number.")
	}
}

func ReadFloat() float64 {
	for {
		text := ReadString()
		number, err := strconv.ParseFloat(text, 64)
		if err == nil {
			return number
		}
		ui.Warningf("Please enter a valid number.")
	}
}

func PromptNonEmpty(label string) string {
	for {
		ui.Promptf("%s", label)
		text := ReadString()
		if text != "" {
			return text
		}
		ui.Warningf("Please enter a non-empty value.")
	}
}

func PromptIntInRange(label string, min, max int) int {
	for {
		ui.Promptf("%s", label)
		n := ReadInt()
		if n >= min && n <= max {
			return n
		}
		ui.Warningf("Please enter a number between %d and %d.", min, max)
	}
}

// PromptChoice prompts for an int in [min, max]. Use min=0 for cancel.
func PromptChoice(label string, min, max int) int {
	return PromptIntInRange(label, min, max)
}

func PromptFloat(label string) float64 {
	ui.Promptf("%s", label)
	return ReadFloat()
}

func PromptFloatNonNegative(label string) float64 {
	for {
		ui.Promptf("%s", label)
		n := ReadFloat()
		if n >= 0 {
			return n
		}
		ui.Warningf("Please enter a non-negative number.")
	}
}

func PromptYesNo(label string) bool {
	for {
		ui.Promptf("%s", label)
		text := strings.ToLower(ReadString())
		switch text {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			ui.Warningf("Please enter y or n.")
		}
	}
}
