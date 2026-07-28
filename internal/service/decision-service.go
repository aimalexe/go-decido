package service

import "go-decido/internal/decision"

type DecisionService struct {
	decisions []decision.Decision
}

func NewDecisionService() *DecisionService {
	return &DecisionService{
		decisions: make([]decision.Decision, 0),
	}
}

func (s *DecisionService) AddDecision(d decision.Decision) {
	s.decisions = append(s.decisions, d)
}

func (s *DecisionService) Decisions() []decision.Decision {
	return s.decisions
}

func (s *DecisionService) GetDecision(index int) (decision.Decision, bool) {
	if index < 0 || index >= len(s.decisions) {
		return decision.Decision{}, false
	}

	return s.decisions[index], true
}

func (s *DecisionService) UpdateDecision(index int, d decision.Decision) bool {
	if index < 0 || index >= len(s.decisions) {
		return false
	}

	s.decisions[index] = d
	return true
}
