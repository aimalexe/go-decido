package decision

// func AddAlternative(decision *Decision) {
// 	fmt.Print("Alternative name: ")

// 	name := input.ReadString()

// 	alternative := Alternative{
// 		Name:   name,
// 		Scores: make(map[string]int),
// 	}

// 	decision.Alternatives = append(
// 		decision.Alternatives,
// 		alternative,
// 	)

// 	fmt.Println("Alternative added.")
// }

func AddAlternative(decision *Decision, name string) {
	alternative := Alternative{
		ID:     len(decision.Alternatives) + 1,
		Name:   name,
		Scores: make(map[int]int),
	}

	decision.Alternatives = append(
		decision.Alternatives,
		alternative,
	)
}
