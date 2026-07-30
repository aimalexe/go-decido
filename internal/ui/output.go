package ui

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

var output io.Writer = color.Output

// SetOutput redirects all UI output and returns a function that restores the
// previous writer. It is useful for tests and alternate terminal front ends.
func SetOutput(w io.Writer) (restore func()) {
	previous := output
	output = w
	return func() {
		output = previous
	}
}

func Printf(format string, args ...any) {
	fmt.Fprintf(output, format, args...)
}

func Println(args ...any) {
	fmt.Fprintln(output, args...)
}

func Blank() {
	fmt.Fprintln(output)
}

func Heading(format string, args ...any) {
	headerStyle.Fprintf(output, format+"\n", args...)
}

func Section(title string) {
	headerStyle.Fprintf(output, "\n%s %s\n", IconArrow, title)
}

func Successf(format string, args ...any) {
	printMessage(successStyle, IconSuccess, format, args...)
}

func Errorf(format string, args ...any) {
	printMessage(errorStyle, IconError, format, args...)
}

func Warningf(format string, args ...any) {
	printMessage(warningStyle, IconWarning, format, args...)
}

func Infof(format string, args ...any) {
	printMessage(infoStyle, IconInfo, format, args...)
}

func Mutedf(format string, args ...any) {
	mutedStyle.Fprintf(output, format+"\n", args...)
}

func Promptf(format string, args ...any) {
	promptStyle.Fprintf(output, "%s "+format+" ", append([]any{IconArrow}, args...)...)
}

func Highlightf(format string, args ...any) {
	highlightStyle.Fprintf(output, format+"\n", args...)
}

func PrintSuccess(message string) {
	Successf("%s", message)
}

func PrintError(err error) {
	Errorf("%v", err)
}

func printMessage(style *color.Color, icon, format string, args ...any) {
	style.Fprintf(output, "%s "+format+"\n", append([]any{icon}, args...)...)
}
