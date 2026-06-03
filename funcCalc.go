package main

type Searchable interface {
	ShowAnswers(n1, n2 float64) float64
}

type Sum struct {
	Operation
}

func (s Sum) ShowAnswers(n1, n2 float64) float64 {
	return n1 + n2
}