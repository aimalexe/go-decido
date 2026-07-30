package ui

func PrintBanner() {
	Blank()
	primaryStyle.Fprintln(output, "  ╔══════════════════════════════════╗")
	primaryStyle.Fprintln(output, "  ║          GO · DECIDO             ║")
	primaryStyle.Fprintln(output, "  ╚══════════════════════════════════╝")
	mutedStyle.Fprintln(output, "       Weigh options. Decide clearly.")
}

func PrintGoodbye() {
	Blank()
	Successf("Thank you for using go-decido. Goodbye!")
}
