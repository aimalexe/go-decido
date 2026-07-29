package service

import "go-decido/internal/decision"

// MemoryStore is an in-memory Store suitable for the CLI session.
type MemoryStore struct {
	baseStore
}

var _ Store = (*MemoryStore)(nil)

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		baseStore: baseStore{
			byID:   make(map[int]decision.Decision),
			order:  []int{},
			nextID: 1,
		},
	}
}

func (m *MemoryStore) Add(d decision.Decision) (decision.Decision, error) {
	return m.add(d)
}

func (m *MemoryStore) List() []decision.Decision {
	return m.list()
}

func (m *MemoryStore) Get(id int) (decision.Decision, error) {
	return m.get(id)
}

func (m *MemoryStore) Update(d decision.Decision) error {
	return m.update(d)
}
