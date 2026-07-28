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
		println("Please enter a valid number:")
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
