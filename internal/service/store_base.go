package service

import (
	"fmt"
	"go-decido/internal/decision"
)

type baseStore struct {
	byID   map[int]decision.Decision
	order  []int
	nextID int
}

func (s *baseStore) add(d decision.Decision) (decision.Decision, error) {
	d.ID = s.nextID
	s.nextID++

	s.byID[d.ID] = cloneDecision(d)
	s.order = append(s.order, d.ID)

	return cloneDecision(d), nil
}

func (s *baseStore) list() []decision.Decision {
	out := make([]decision.Decision, 0, len(s.order))

	for _, id := range s.order {
		if d, ok := s.byID[id]; ok {
			out = append(out, cloneDecision(d))
		}
	}

	return out
}

func (s *baseStore) get(id int) (decision.Decision, error) {
	d, ok := s.byID[id]
	if !ok {
		return decision.Decision{}, fmt.Errorf("%w: id %d", ErrNotFound, id)
	}

	return cloneDecision(d), nil
}

func (s *baseStore) update(d decision.Decision) error {
	if _, ok := s.byID[d.ID]; !ok {
		return fmt.Errorf("%w: id %d", ErrNotFound, d.ID)
	}

	s.byID[d.ID] = cloneDecision(d)
	return nil
}

func (s *baseStore) delete(id int) error {
	if _, ok := s.byID[id]; !ok {
		return fmt.Errorf("%w: id %d", ErrNotFound, id)
	}

	delete(s.byID, id)
	for i, existingID := range s.order {
		if existingID == id {
			s.order = append(s.order[:i], s.order[i+1:]...)
			break
		}
	}
	return nil
}

func cloneDecision(d decision.Decision) decision.Decision {
	cp := d

	if d.Criteria != nil {
		cp.Criteria = make([]decision.Criterion, len(d.Criteria))
		copy(cp.Criteria, d.Criteria)
	}

	if d.Alternatives != nil {
		cp.Alternatives = make([]decision.Alternative, len(d.Alternatives))

		for i, a := range d.Alternatives {
			cp.Alternatives[i] = a

			if a.Scores != nil {
				cp.Alternatives[i].Scores = make(map[int]int, len(a.Scores))

				for k, v := range a.Scores {
					cp.Alternatives[i].Scores[k] = v
				}
			}
		}
	}

	return cp
}
