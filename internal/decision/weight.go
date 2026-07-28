package decision

func SetWeight(
	decision *Decision,
	criterionID int,
	weight float64,
) {
	criterion := decision.FindCriterion(criterionID)
	if criterion == nil {
		return
	}
	criterion.Weight = weight
}
