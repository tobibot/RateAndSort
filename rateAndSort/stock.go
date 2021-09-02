package rateAndSort

type stock struct {
	Symbol string
	Name   string
	Value  int
	Type   StockType
}

type StockType string

var (
	stockTypes = []StockType{hypertech, techindustry, value}
)

const (
	hypertech    StockType = "HyperTech"
	techindustry StockType = "TechIndustry"
	value        StockType = "Value"
)

func (s *stock) decreaseBy(b int) {
	s.Value -= b
	if s.Value < 0 {
		s.Value = 0
	}
}

func (s *stock) increaseBy(b int) {
	s.Value += b
	if s.Value > 100 {
		s.Value = 100
	}
}
