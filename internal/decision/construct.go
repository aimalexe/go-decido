package decision

import (
	"errors"
	"strings"
)

var (
	ErrEmptyTitle          = errors.New("title must not be empty")
	ErrEmptyName           = errors.New("name must not be empty")
	ErrNegativeWeight      = errors.New("weight must not be negative")
	ErrInvalidScore        = errors.New("score must be between 1 and 5")
	ErrCriterionNotFound   = errors.New("criterion not found")
	ErrAlternativeNotFound = errors.New("alternative not found")
	ErrNoCriteria          = errors.New("decision has no criteria")
	ErrNoAlternatives      = errors.New("decision has no alternatives")
	ErrZeroWeight          = errors.New("total weight must be greater than zero")
	ErrIncompleteScores    = errors.New("not all alternatives are fully scored")
)

func NewDecision(title string) (Decision, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Decision{}, ErrEmptyTitle
	}
	return Decision{
		Title:             title,
		Criteria:          make([]Criterion, 0),
		Alternatives:      make([]Alternative, 0),
		nextCriterionID:   1,
		nextAlternativeID: 1,
	}, nil
}

func NewCriterion(id int, name string) (Criterion, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Criterion{}, ErrEmptyName
	}
	return Criterion{ID: id, Name: name}, nil
}

func NewAlternative(id int, name string) (Alternative, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Alternative{}, ErrEmptyName
	}
	return Alternative{
		ID:     id,
		Name:   name,
		Scores: make(map[int]int),
	}, nil
}
