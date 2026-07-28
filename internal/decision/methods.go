package decision

func (d *Decision) FindCriterion(id int) *Criterion {
	for i := range d.Criteria {
		if d.Criteria[i].ID == id {
			return &d.Criteria[i]
		}
	}
	return nil
}

func (d *Decision) FindAlternative(id int) *Alternative {
	for i := range d.Alternatives {
		if d.Alternatives[i].ID == id {
			return &d.Alternatives[i]
		}
	}
	return nil
}

func (a *Alternative) SetScore(criterionID int, score int) {
	a.Scores[criterionID] = score
}

func (a Alternative) ScoreFor(criterionID int) int {
	return a.Scores[criterionID]
}
