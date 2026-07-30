package decision

import "strings"

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

func RenameDecision(d *Decision, title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return ErrEmptyTitle
	}
	d.Title = title
	return nil
}

func RenameCriterion(d *Decision, criterionID int, name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrEmptyName
	}
	criterion := d.FindCriterion(criterionID)
	if criterion == nil {
		return ErrCriterionNotFound
	}
	criterion.Name = name
	return nil
}

func RenameAlternative(d *Decision, alternativeID int, name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrEmptyName
	}
	alternative := d.FindAlternative(alternativeID)
	if alternative == nil {
		return ErrAlternativeNotFound
	}
	alternative.Name = name
	return nil
}

func DeleteCriterion(d *Decision, criterionID int) error {
	index := -1
	for i, criterion := range d.Criteria {
		if criterion.ID == criterionID {
			index = i
			break
		}
	}
	if index < 0 {
		return ErrCriterionNotFound
	}

	d.Criteria = append(d.Criteria[:index], d.Criteria[index+1:]...)
	for i := range d.Alternatives {
		delete(d.Alternatives[i].Scores, criterionID)
	}
	return nil
}

func DeleteAlternative(d *Decision, alternativeID int) error {
	index := -1
	for i, alternative := range d.Alternatives {
		if alternative.ID == alternativeID {
			index = i
			break
		}
	}
	if index < 0 {
		return ErrAlternativeNotFound
	}

	d.Alternatives = append(d.Alternatives[:index], d.Alternatives[index+1:]...)
	return nil
}
