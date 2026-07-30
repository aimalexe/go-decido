package decision_test

import (
	"errors"
	"testing"

	"go-decido/internal/decision"
)

func TestNewDecision_EmptyTitle(t *testing.T) {
	_, err := decision.NewDecision("   ")
	if !errors.Is(err, decision.ErrEmptyTitle) {
		t.Fatalf("got %v, want ErrEmptyTitle", err)
	}
}

func TestAddCriterion_EmptyName(t *testing.T) {
	d, err := decision.NewDecision("Laptop")
	if err != nil {
		t.Fatal(err)
	}
	if err := decision.AddCriterion(&d, ""); !errors.Is(err, decision.ErrEmptyName) {
		t.Fatalf("got %v, want ErrEmptyName", err)
	}
}

func TestAddCriterion_StableIDs(t *testing.T) {
	d, _ := decision.NewDecision("Laptop")
	_ = decision.AddCriterion(&d, "Price")
	_ = decision.AddCriterion(&d, "Battery")
	if d.Criteria[0].ID != 1 || d.Criteria[1].ID != 2 {
		t.Fatalf("unexpected IDs: %+v", d.Criteria)
	}
}

func TestSetScore_OutOfRange(t *testing.T) {
	a, _ := decision.NewAlternative(1, "ThinkPad")
	if err := a.SetScore(1, 0); !errors.Is(err, decision.ErrInvalidScore) {
		t.Fatalf("got %v, want ErrInvalidScore", err)
	}
	if err := a.SetScore(1, 6); !errors.Is(err, decision.ErrInvalidScore) {
		t.Fatalf("got %v, want ErrInvalidScore", err)
	}
}

func TestSetWeight_NotFoundAndNegative(t *testing.T) {
	d, _ := decision.NewDecision("Laptop")
	_ = decision.AddCriterion(&d, "Price")

	if err := decision.SetWeight(&d, 99, 1); !errors.Is(err, decision.ErrCriterionNotFound) {
		t.Fatalf("got %v, want ErrCriterionNotFound", err)
	}
	if err := decision.SetWeight(&d, 1, -1); !errors.Is(err, decision.ErrNegativeWeight) {
		t.Fatalf("got %v, want ErrNegativeWeight", err)
	}
}

func TestValidateReady(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() decision.Decision
		wantErr error
	}{
		{
			name: "no criteria",
			setup: func() decision.Decision {
				d, _ := decision.NewDecision("X")
				return d
			},
			wantErr: decision.ErrNoCriteria,
		},
		{
			name: "no alternatives",
			setup: func() decision.Decision {
				d, _ := decision.NewDecision("X")
				_ = decision.AddCriterion(&d, "Price")
				_ = decision.SetWeight(&d, 1, 1)
				return d
			},
			wantErr: decision.ErrNoAlternatives,
		},
		{
			name: "zero weight",
			setup: func() decision.Decision {
				d, _ := decision.NewDecision("X")
				_ = decision.AddCriterion(&d, "Price")
				_ = decision.AddAlternative(&d, "A")
				return d
			},
			wantErr: decision.ErrZeroWeight,
		},
		{
			name: "incomplete scores",
			setup: func() decision.Decision {
				d, _ := decision.NewDecision("X")
				_ = decision.AddCriterion(&d, "Price")
				_ = decision.SetWeight(&d, 1, 1)
				_ = decision.AddAlternative(&d, "A")
				return d
			},
			wantErr: decision.ErrIncompleteScores,
		},
		{
			name: "ready",
			setup: func() decision.Decision {
				d, _ := decision.NewDecision("X")
				_ = decision.AddCriterion(&d, "Price")
				_ = decision.SetWeight(&d, 1, 1)
				_ = decision.AddAlternative(&d, "A")
				_ = d.Alternatives[0].SetScore(1, 4)
				return d
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.setup()
			err := d.ValidateReady()
			if tt.wantErr == nil {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalculateResults_Ranking(t *testing.T) {
	d, _ := decision.NewDecision("Laptop")
	_ = decision.AddCriterion(&d, "Price")
	_ = decision.AddCriterion(&d, "Battery")
	_ = decision.SetWeight(&d, 1, 1)
	_ = decision.SetWeight(&d, 2, 1)
	_ = decision.AddAlternative(&d, "Cheap")
	_ = decision.AddAlternative(&d, "Premium")

	_ = d.Alternatives[0].SetScore(1, 5)
	_ = d.Alternatives[0].SetScore(2, 2)
	_ = d.Alternatives[1].SetScore(1, 3)
	_ = d.Alternatives[1].SetScore(2, 5)

	results := d.CalculateResults()
	if len(results) != 2 {
		t.Fatalf("got %d results", len(results))
	}
	// Cheap: (5+2)/2 = 3.5; Premium: (3+5)/2 = 4.0
	if results[0].AlternativeName != "Premium" {
		t.Fatalf("want Premium first, got %s", results[0].AlternativeName)
	}
	if results[0].Score != 4.0 {
		t.Fatalf("want score 4.0, got %.2f", results[0].Score)
	}
}

func TestHasTie(t *testing.T) {
	tied := []decision.Result{
		{AlternativeName: "A", Score: 3},
		{AlternativeName: "B", Score: 3},
	}
	if !decision.HasTie(tied) {
		t.Fatal("expected tie")
	}
	unique := []decision.Result{
		{AlternativeName: "A", Score: 4},
		{AlternativeName: "B", Score: 3},
	}
	if decision.HasTie(unique) {
		t.Fatal("expected no tie")
	}
}

func TestScoreFor_Ok(t *testing.T) {
	a, _ := decision.NewAlternative(1, "A")
	_ = a.SetScore(2, 4)
	score, ok := a.ScoreFor(2)
	if !ok || score != 4 {
		t.Fatalf("got %d, %v", score, ok)
	}
	if _, ok := a.ScoreFor(99); ok {
		t.Fatal("expected missing score")
	}
}

func TestRenameAndDeleteEntities(t *testing.T) {
	d, _ := decision.NewDecision("Laptop")
	_ = decision.AddCriterion(&d, "Price")
	_ = decision.AddCriterion(&d, "Battery")
	_ = decision.AddAlternative(&d, "ThinkPad")
	_ = decision.AddAlternative(&d, "MacBook")
	_ = d.Alternatives[0].SetScore(1, 4)
	_ = d.Alternatives[0].SetScore(2, 5)
	_ = d.Alternatives[1].SetScore(1, 3)
	_ = d.Alternatives[1].SetScore(2, 4)

	if err := decision.RenameDecision(&d, "  Buy laptop  "); err != nil {
		t.Fatal(err)
	}
	if d.Title != "Buy laptop" {
		t.Fatalf("title = %q", d.Title)
	}

	if err := decision.RenameCriterion(&d, 1, "Cost"); err != nil {
		t.Fatal(err)
	}
	if d.Criteria[0].Name != "Cost" {
		t.Fatalf("criterion name = %q", d.Criteria[0].Name)
	}

	if err := decision.RenameAlternative(&d, 2, "MacBook Pro"); err != nil {
		t.Fatal(err)
	}
	if d.Alternatives[1].Name != "MacBook Pro" {
		t.Fatalf("alternative name = %q", d.Alternatives[1].Name)
	}

	if err := decision.DeleteCriterion(&d, 1); err != nil {
		t.Fatal(err)
	}
	if len(d.Criteria) != 1 || d.Criteria[0].ID != 2 {
		t.Fatalf("unexpected criteria after delete: %+v", d.Criteria)
	}
	if _, ok := d.Alternatives[0].ScoreFor(1); ok {
		t.Fatal("expected orphaned criterion score to be removed")
	}
	if score, ok := d.Alternatives[0].ScoreFor(2); !ok || score != 5 {
		t.Fatalf("expected remaining score 5, got %d, %v", score, ok)
	}

	if err := decision.DeleteAlternative(&d, 1); err != nil {
		t.Fatal(err)
	}
	if len(d.Alternatives) != 1 || d.Alternatives[0].ID != 2 {
		t.Fatalf("unexpected alternatives after delete: %+v", d.Alternatives)
	}

	if err := decision.DeleteCriterion(&d, 99); !errors.Is(err, decision.ErrCriterionNotFound) {
		t.Fatalf("got %v, want ErrCriterionNotFound", err)
	}
	if err := decision.DeleteAlternative(&d, 99); !errors.Is(err, decision.ErrAlternativeNotFound) {
		t.Fatalf("got %v, want ErrAlternativeNotFound", err)
	}
	if err := decision.RenameDecision(&d, " "); !errors.Is(err, decision.ErrEmptyTitle) {
		t.Fatalf("got %v, want ErrEmptyTitle", err)
	}
}
