package decision

type Criterion struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

type Alternative struct {
	ID     int         `json:"id"`
	Name   string      `json:"name"`
	Scores map[int]int `json:"scores"` // criterion ID -> score (1-5)
}

type Decision struct {
	ID           int           `json:"id"`
	Title        string        `json:"title"`
	Criteria     []Criterion   `json:"criteria"`
	Alternatives []Alternative `json:"alternatives"`

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
