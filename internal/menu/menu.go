package menu

import (
	"go-decido/internal/input"
	"fmt"
)

var options = []string{
	"Create new decision",
	"List decisions",
	"View decision",
	"Add criterion",
	"Add alternative",
	"Score alternatives",
	"Set criterion weights",
	"Calculate Results",
	"Exit",
}

func Show() int {
	fmt.Println("============ MENU ============")
	fmt.Println("> Please choose an option:")
	for i, option := range options {
		fmt.Printf("\t%d. %s\n", i+1, option)
	}

	fmt.Print("> Enter your choice: ")
	return input.ReadInt()
}
