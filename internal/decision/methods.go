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

func (a *Alternative) SetScore(criterionID int, score int) error {
	if score < MinScore || score > MaxScore {
		return ErrInvalidScore
	}
	if a.Scores == nil {
		a.Scores = make(map[int]int)
	}
	a.Scores[criterionID] = score
	return nil
}

func (a Alternative) ScoreFor(criterionID int) (int, bool) {
	score, ok := a.Scores[criterionID]
	return score, ok
}

func (d Decision) TotalWeight() float64 {
	var total float64
	for _, c := range d.Criteria {
		total += c.Weight
	}
	return total
}

func (d Decision) IsFullyScored() bool {
	if len(d.Criteria) == 0 || len(d.Alternatives) == 0 {
		return false
	}
	for _, alt := range d.Alternatives {
		for _, c := range d.Criteria {
			if _, ok := alt.ScoreFor(c.ID); !ok {
				return false
			}
		}
	}
	return true
}
