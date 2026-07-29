package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"go-decido/internal/decision"
)

type JSONFileStore struct {
	baseStore
	path string
}

var _ Store = (*JSONFileStore)(nil)

type jsonData struct {
	NextID    int                 `json:"next_id"`
	Order     []int               `json:"order"`
	Decisions []decision.Decision `json:"decisions"`
}

func NewJSONFileStore(path string) (*JSONFileStore, error) {
	store := &JSONFileStore{
		baseStore: baseStore{
			byID:   make(map[int]decision.Decision),
			order:  []int{},
			nextID: 1,
		},
		path: path,
	}

	if err := store.Load(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *JSONFileStore) marshal() ([]byte, error) {
	file := jsonData{
		NextID: s.nextID,
		Order:  append([]int(nil), s.order...),
	}

	for _, id := range s.order {
		if d, ok := s.byID[id]; ok {
			file.Decisions = append(file.Decisions, cloneDecision(d))
		}
	}

	return json.MarshalIndent(file, "", "  ")
}

func (s *JSONFileStore) Add(d decision.Decision) (decision.Decision, error) {
	result, err := s.add(d)
	if err != nil {
		return result, err
	}

	if err := s.Save(); err != nil {
		return decision.Decision{}, err
	}

	return result, nil
}

func (s *JSONFileStore) List() []decision.Decision {
	return s.list()
}

func (s *JSONFileStore) Get(id int) (decision.Decision, error) {
	return s.get(id)
}

func (s *JSONFileStore) Update(d decision.Decision) error {
	if err := s.update(d); err != nil {
		return err
	}

	return s.Save()
}

func (s *JSONFileStore) Save() error {
	data, err := s.marshal()
	if err != nil {
		return fmt.Errorf("encode decisions: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(s.path), 0755); err != nil {
		return fmt.Errorf("create data directory: %w", err)
	}

	if err := os.WriteFile(s.path, data, 0644); err != nil {
		return fmt.Errorf("write decisions file: %w", err)
	}

	return nil
}

func (s *JSONFileStore) Load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("read decisions file: %w", err)
	}

	if len(data) == 0 {
		return nil
	}

	var file jsonData

	if err := json.Unmarshal(data, &file); err != nil {
		return fmt.Errorf("decode decisions file: %w", err)
	}

	s.byID = make(map[int]decision.Decision, len(file.Decisions))
	s.order = make([]int, 0, len(file.Decisions))

	seen := make(map[int]bool, len(file.Decisions))
	maxID := 0

	for _, d := range file.Decisions {
		if d.ID <= 0 {
			return fmt.Errorf("decode decisions file: invalid decision id %d", d.ID)
		}

		if seen[d.ID] {
			return fmt.Errorf("decode decisions file: duplicate decision id %d", d.ID)
		}

		d.RebuildCounters()

		s.byID[d.ID] = cloneDecision(d)
		seen[d.ID] = true

		if d.ID > maxID {
			maxID = d.ID
		}
	}

	for _, id := range file.Order {
		if seen[id] {
			s.order = append(s.order, id)
			seen[id] = false
		}
	}

	for _, d := range file.Decisions {
		if seen[d.ID] {
			s.order = append(s.order, d.ID)
		}
	}

	s.nextID = maxID + 1

	if file.NextID > s.nextID {
		s.nextID = file.NextID
	}

	return nil
}
