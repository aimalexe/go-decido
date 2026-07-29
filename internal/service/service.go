package service

import (
	"errors"
	"fmt"

	"go-decido/internal/decision"
)

var ErrNotFound = errors.New("decision not found")

// Store persists decisions. Implementations must return copies from List/Get
// so callers cannot mutate the backing store without Update.
type Store interface {
	Add(d decision.Decision) decision.Decision
	List() []decision.Decision
	Get(id int) (decision.Decision, error)
	Update(d decision.Decision) error
}

type DecisionService struct {
	store  Store
	nextID int
}

func NewDecisionService(store Store) *DecisionService {
	if store == nil {
		store = NewMemoryStore()
	}
	return &DecisionService{
		store:  store,
		nextID: 1,
	}
}

func NewDefaultService() *DecisionService {
	return NewDecisionService(NewMemoryStore())
}

func (s *DecisionService) Create(title string) (decision.Decision, error) {
	d, err := decision.NewDecision(title)
	if err != nil {
		return decision.Decision{}, err
	}
	d.ID = s.nextID
	s.nextID++
	return s.store.Add(d), nil
}

func (s *DecisionService) List() []decision.Decision {
	return s.store.List()
}

func (s *DecisionService) Get(id int) (decision.Decision, error) {
	return s.store.Get(id)
}

func (s *DecisionService) Count() int {
	return len(s.store.List())
}

func (s *DecisionService) AddCriterion(decisionID int, name string) (decision.Decision, error) {
	d, err := s.store.Get(decisionID)
	if err != nil {
		return decision.Decision{}, err
	}
	if err := decision.AddCriterion(&d, name); err != nil {
		return decision.Decision{}, err
	}
	if err := s.store.Update(d); err != nil {
		return decision.Decision{}, err
	}
	return d, nil
}

func (s *DecisionService) AddAlternative(decisionID int, name string) (decision.Decision, error) {
	d, err := s.store.Get(decisionID)
	if err != nil {
		return decision.Decision{}, err
	}
	if err := decision.AddAlternative(&d, name); err != nil {
		return decision.Decision{}, err
	}
	if err := s.store.Update(d); err != nil {
		return decision.Decision{}, err
	}
	return d, nil
}

func (s *DecisionService) SetWeight(decisionID, criterionID int, weight float64) (decision.Decision, error) {
	d, err := s.store.Get(decisionID)
	if err != nil {
		return decision.Decision{}, err
	}
	if err := decision.SetWeight(&d, criterionID, weight); err != nil {
		return decision.Decision{}, err
	}
	if err := s.store.Update(d); err != nil {
		return decision.Decision{}, err
	}
	return d, nil
}

func (s *DecisionService) SetScore(decisionID, alternativeID, criterionID, score int) (decision.Decision, error) {
	d, err := s.store.Get(decisionID)
	if err != nil {
		return decision.Decision{}, err
	}
	alt := d.FindAlternative(alternativeID)
	if alt == nil {
		return decision.Decision{}, decision.ErrAlternativeNotFound
	}
	if d.FindCriterion(criterionID) == nil {
		return decision.Decision{}, decision.ErrCriterionNotFound
	}
	if err := alt.SetScore(criterionID, score); err != nil {
		return decision.Decision{}, err
	}
	if err := s.store.Update(d); err != nil {
		return decision.Decision{}, err
	}
	return d, nil
}

func (s *DecisionService) Results(decisionID int) ([]decision.Result, error) {
	d, err := s.store.Get(decisionID)
	if err != nil {
		return nil, err
	}
	if err := d.ValidateReady(); err != nil {
		return nil, fmt.Errorf("cannot calculate results: %w", err)
	}
	return d.CalculateResults(), nil
}
