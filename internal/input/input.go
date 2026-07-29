package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		fmt.Println("Please enter a valid number.")
	}
}

func ReadFloat() float64 {
	for {
		text := ReadString()
		number, err := strconv.ParseFloat(text, 64)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid number.")
	}
}

func PromptNonEmpty(label string) string {
	for {
		fmt.Print(label + " ")
		text := ReadString()
		if text != "" {
			return text
		}
		fmt.Println("Please enter a non-empty value.")
	}
}

func PromptIntInRange(label string, min, max int) int {
	for {
		fmt.Print(label + " ")
		n := ReadInt()
		if n >= min && n <= max {
			return n
		}
		fmt.Printf("Please enter a number between %d and %d.\n", min, max)
	}
}

// PromptChoice prompts for an int in [min, max]. Use min=0 for cancel.
func PromptChoice(label string, min, max int) int {
	return PromptIntInRange(label, min, max)
}

func PromptFloat(label string) float64 {
	fmt.Print(label + " ")
	return ReadFloat()
}

func PromptFloatNonNegative(label string) float64 {
	for {
		fmt.Print(label + " ")
		n := ReadFloat()
		if n >= 0 {
			return n
		}
		fmt.Println("Please enter a non-negative number.")
	}
}

func PromptYesNo(label string) bool {
	for {
		fmt.Print(label + " ")
		text := strings.ToLower(ReadString())
		switch text {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Please enter y or n.")
		}
	}
}
