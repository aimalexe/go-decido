package decision

func SetWeight(d *Decision, criterionID int, weight float64) error {
	if weight < 0 {
		return ErrNegativeWeight
	}
	criterion := d.FindCriterion(criterionID)
	if criterion == nil {
		return ErrCriterionNotFound
	}
	criterion.Weight = weight
	return nil
}
