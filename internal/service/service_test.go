package service_test

import (
	"errors"
	"testing"

	"go-decido/internal/decision"
	"go-decido/internal/service"
)

func TestCreateAndList(t *testing.T) {
	svc := service.NewDefaultService()
	d, err := svc.Create("Buy laptop")
	if err != nil {
		t.Fatal(err)
	}
	if d.ID != 1 {
		t.Fatalf("want ID 1, got %d", d.ID)
	}
	list := svc.List()
	if len(list) != 1 || list[0].Title != "Buy laptop" {
		t.Fatalf("unexpected list: %+v", list)
	}
}

func TestListReturnsCopies(t *testing.T) {
	svc := service.NewDefaultService()
	_, _ = svc.Create("A")
	list := svc.List()
	list[0].Title = "mutated"
	got, _ := svc.Get(1)
	if got.Title != "A" {
		t.Fatalf("store was mutated via List: %q", got.Title)
	}
}

func TestAddCriterionAndAlternative(t *testing.T) {
	svc := service.NewDefaultService()
	d, _ := svc.Create("Laptop")
	d, err := svc.AddCriterion(d.ID, "Price")
	if err != nil {
		t.Fatal(err)
	}
	d, err = svc.AddAlternative(d.ID, "ThinkPad")
	if err != nil {
		t.Fatal(err)
	}
	if len(d.Criteria) != 1 || len(d.Alternatives) != 1 {
		t.Fatalf("unexpected state: %+v", d)
	}
}

func TestUnknownDecisionID(t *testing.T) {
	svc := service.NewDefaultService()
	_, err := svc.Get(99)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("got %v, want ErrNotFound", err)
	}
}

func TestResultsFlow(t *testing.T) {
	svc := service.NewDefaultService()
	d, _ := svc.Create("Laptop")
	_, _ = svc.AddCriterion(d.ID, "Price")
	_, _ = svc.AddCriterion(d.ID, "Battery")
	_, _ = svc.SetWeight(d.ID, 1, 0.4)
	_, _ = svc.SetWeight(d.ID, 2, 0.6)
	_, _ = svc.AddAlternative(d.ID, "ThinkPad")
	_, _ = svc.AddAlternative(d.ID, "MacBook")

	_, err := svc.Results(d.ID)
	if !errors.Is(err, decision.ErrIncompleteScores) {
		t.Fatalf("got %v, want incomplete scores", err)
	}

	_, _ = svc.SetScore(d.ID, 1, 1, 4)
	_, _ = svc.SetScore(d.ID, 1, 2, 5)
	_, _ = svc.SetScore(d.ID, 2, 1, 3)
	_, _ = svc.SetScore(d.ID, 2, 2, 4)

	results, err := svc.Results(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 2 {
		t.Fatalf("got %d results", len(results))
	}
	if results[0].AlternativeName != "ThinkPad" {
		t.Fatalf("want ThinkPad first, got %s (%.2f)", results[0].AlternativeName, results[0].Score)
	}
}

func TestSetScore_UnknownAlternative(t *testing.T) {
	svc := service.NewDefaultService()
	d, _ := svc.Create("X")
	_, _ = svc.AddCriterion(d.ID, "Price")
	_, err := svc.SetScore(d.ID, 99, 1, 3)
	if !errors.Is(err, decision.ErrAlternativeNotFound) {
		t.Fatalf("got %v, want ErrAlternativeNotFound", err)
	}
}

func TestCreate_EmptyTitle(t *testing.T) {
	svc := service.NewDefaultService()
	_, err := svc.Create("")
	if !errors.Is(err, decision.ErrEmptyTitle) {
		t.Fatalf("got %v, want ErrEmptyTitle", err)
	}
}

func TestRenameAndDeleteDecision(t *testing.T) {
	svc := service.NewDefaultService()
	d, _ := svc.Create("Laptop")
	_, _ = svc.Create("Phone")

	renamed, err := svc.Rename(d.ID, "Buy laptop")
	if err != nil {
		t.Fatal(err)
	}
	if renamed.Title != "Buy laptop" {
		t.Fatalf("title = %q", renamed.Title)
	}

	if err := svc.Delete(d.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.Get(d.ID); !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("got %v, want ErrNotFound", err)
	}
	if svc.Count() != 1 {
		t.Fatalf("count = %d, want 1", svc.Count())
	}
}

func TestRenameAndDeleteCriterionAndAlternative(t *testing.T) {
	svc := service.NewDefaultService()
	d, _ := svc.Create("Laptop")
	_, _ = svc.AddCriterion(d.ID, "Price")
	_, _ = svc.AddCriterion(d.ID, "Battery")
	_, _ = svc.AddAlternative(d.ID, "ThinkPad")
	_, _ = svc.AddAlternative(d.ID, "MacBook")
	_, _ = svc.SetScore(d.ID, 1, 1, 4)
	_, _ = svc.SetScore(d.ID, 1, 2, 5)

	updated, err := svc.RenameCriterion(d.ID, 1, "Cost")
	if err != nil {
		t.Fatal(err)
	}
	if updated.Criteria[0].Name != "Cost" {
		t.Fatalf("criterion = %q", updated.Criteria[0].Name)
	}

	updated, err = svc.RenameAlternative(d.ID, 2, "MacBook Pro")
	if err != nil {
		t.Fatal(err)
	}
	if updated.FindAlternative(2).Name != "MacBook Pro" {
		t.Fatalf("alternative = %q", updated.FindAlternative(2).Name)
	}

	updated, err = svc.DeleteCriterion(d.ID, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(updated.Criteria) != 1 || updated.Criteria[0].ID != 2 {
		t.Fatalf("unexpected criteria: %+v", updated.Criteria)
	}
	if _, ok := updated.FindAlternative(1).ScoreFor(1); ok {
		t.Fatal("expected score for deleted criterion to be removed")
	}

	updated, err = svc.DeleteAlternative(d.ID, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(updated.Alternatives) != 1 || updated.Alternatives[0].ID != 2 {
		t.Fatalf("unexpected alternatives: %+v", updated.Alternatives)
	}
}
