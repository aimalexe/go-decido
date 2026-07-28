package decision

type Criterion struct {
	ID     int
	Name   string
	Weight float64
}

type Alternative struct {
	ID     int
	Name   string
	Scores map[int]int
}

type Decision struct {
	ID           int
	Title        string
	Criteria     []Criterion
	Alternatives []Alternative
}

type Result struct {
	AlternativeName string
	Score           float64
}
