package decision

func AddCriterion(d *Decision, name string) error {
	c, err := NewCriterion(d.nextCriterionID, name)
	if err != nil {
		return err
	}
	d.Criteria = append(d.Criteria, c)
	d.nextCriterionID++
	return nil
}

func AddAlternative(d *Decision, name string) error {
	a, err := NewAlternative(d.nextAlternativeID, name)
	if err != nil {
		return err
	}
	d.Alternatives = append(d.Alternatives, a)
	d.nextAlternativeID++
	return nil
}
