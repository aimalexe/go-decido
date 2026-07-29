package decision

type Criterion struct {
	ID     int
	Name   string
	Weight float64
}

type Alternative struct {
	ID     int
	Name   string
	Scores map[int]int // criterion ID -> score (1-5)
}

type Decision struct {
	ID           int
	Title        string
	Criteria     []Criterion
	Alternatives []Alternative

	nextCriterionID   int
	nextAlternativeID int
}

type Result struct {
	AlternativeID   int
	AlternativeName string
	Score           float64
}

const (
	MinScore = 1
	MaxScore = 5
)
