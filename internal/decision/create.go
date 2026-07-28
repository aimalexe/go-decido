package decision

import (
	"go-decido/internal/input"
	"fmt"
)

func Create() Decision {
	fmt.Print("What decision do you want to make? ")

	title := input.ReadString()

	return Decision{
		Title: title,
	}
}
