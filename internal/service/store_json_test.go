package service_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"go-decido/internal/service"
)

func TestJSONFileStoreRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "data", "decisions.json")

	store, err := service.NewJSONFileStore(path)
	if err != nil {
		t.Fatal(err)
	}
	svc := service.NewDecisionService(store)

	d, err := svc.Create("Buy a laptop")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := svc.AddCriterion(d.ID, "Price"); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.AddAlternative(d.ID, "ThinkPad"); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.SetWeight(d.ID, 1, 0.7); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.SetScore(d.ID, 1, 1, 5); err != nil {
		t.Fatal(err)
	}

	reloadedStore, err := service.NewJSONFileStore(path)
	if err != nil {
		t.Fatal(err)
	}
	reloaded := service.NewDecisionService(reloadedStore)

	got, err := reloaded.Get(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "Buy a laptop" || len(got.Criteria) != 1 || len(got.Alternatives) != 1 {
		t.Fatalf("unexpected reloaded decision: %+v", got)
	}
	if score, ok := got.Alternatives[0].ScoreFor(got.Criteria[0].ID); !ok || score != 5 {
		t.Fatalf("unexpected reloaded score: %d, %v", score, ok)
	}

	// IDs must continue from loaded data, including the unexported domain counters.
	got, err = reloaded.AddCriterion(d.ID, "Battery")
	if err != nil {
		t.Fatal(err)
	}
	if got.Criteria[1].ID != 2 {
		t.Fatalf("want criterion ID 2, got %d", got.Criteria[1].ID)
	}

	second, err := reloaded.Create("Choose a phone")
	if err != nil {
		t.Fatal(err)
	}
	if second.ID != 2 {
		t.Fatalf("want decision ID 2, got %d", second.ID)
	}

	if err := reloaded.Delete(d.ID); err != nil {
		t.Fatal(err)
	}

	afterDelete, err := service.NewJSONFileStore(path)
	if err != nil {
		t.Fatal(err)
	}
	remaining := afterDelete.List()
	if len(remaining) != 1 || remaining[0].ID != second.ID {
		t.Fatalf("unexpected decisions after delete reload: %+v", remaining)
	}
}

func TestJSONFileStoreMissingFileStartsEmpty(t *testing.T) {
	path := filepath.Join(t.TempDir(), "decisions.json")
	store, err := service.NewJSONFileStore(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := store.List(); len(got) != 0 {
		t.Fatalf("want empty store, got %d decisions", len(got))
	}
}

func TestJSONFileStoreRejectsInvalidJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "decisions.json")
	if err := os.WriteFile(path, []byte("{not-json"), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := service.NewJSONFileStore(path)
	if err == nil {
		t.Fatal("expected invalid JSON error")
	}
}

func TestJSONFileStoreUnknownDecision(t *testing.T) {
	store, err := service.NewJSONFileStore(filepath.Join(t.TempDir(), "decisions.json"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.Get(99)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("got %v, want ErrNotFound", err)
	}
}
