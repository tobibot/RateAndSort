package rateAndSort

type stock struct {
	Symbol string
	Name   string
	Value  int
	Type   StockType
}

type StockType string

var (
	stockTypes = []StockType{hyperTech, techIndustry, value, highRisk}
)

const (
	hyperTech    StockType = "HyperTech"
	techIndustry StockType = "TechIndustry"
	value        StockType = "Value"
	highRisk     StockType = "HighRisk"
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
