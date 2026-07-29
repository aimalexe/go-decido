package service

import (
	"fmt"

	"go-decido/internal/decision"
)

// MemoryStore is an in-memory Store suitable for the CLI session.
type MemoryStore struct {
	byID  map[int]decision.Decision
	order []int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		byID:  make(map[int]decision.Decision),
		order: make([]int, 0),
	}
}

func (m *MemoryStore) Add(d decision.Decision) decision.Decision {
	m.byID[d.ID] = cloneDecision(d)
	m.order = append(m.order, d.ID)
	return cloneDecision(d)
}

func (m *MemoryStore) List() []decision.Decision {
	out := make([]decision.Decision, 0, len(m.order))
	for _, id := range m.order {
		if d, ok := m.byID[id]; ok {
			out = append(out, cloneDecision(d))
		}
	}
	return out
}

func (m *MemoryStore) Get(id int) (decision.Decision, error) {
	d, ok := m.byID[id]
	if !ok {
		return decision.Decision{}, fmt.Errorf("%w: id %d", ErrNotFound, id)
	}
	return cloneDecision(d), nil
}

func (m *MemoryStore) Update(d decision.Decision) error {
	if _, ok := m.byID[d.ID]; !ok {
		return fmt.Errorf("%w: id %d", ErrNotFound, d.ID)
	}
	m.byID[d.ID] = cloneDecision(d)
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
