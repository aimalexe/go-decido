package decision

// func AddCriterion(decision *Decision) {
// 	fmt.Print("Enter criterion name: ")
// 	criterionName := input.ReadString()
// 	criterion := Criterion{
// 		Name: criterionName,
// 	}
// 	decision.Criteria = append(decision.Criteria, criterion)

// 	fmt.Println("Criterion added.")
// }

func AddCriterion(decision *Decision, name string) {
	decision.Criteria = append(
		decision.Criteria,
		Criterion{
			ID:   len(decision.Criteria) + 1,
			Name: name,
		},
	)
}
