package chart

import (
	"fmt"
	"math"
	"strings"
)

// NewBarChart Create a barchart instance
func NewBarChart(title string, tick float64, series []Series) Chart {
	return &barChart{title, tick, series}
}

// Generate generate horizontal or vertical cli barchart
func (bc barChart) Generate(chartType int) string {
	if chartType == Vertical {
		return bc.generateVertically()
	}
	return bc.generateHorizontally()
}

func (bc barChart) generateVertically() string {
	var title = center(bc.title, " ", len(bc.series)*2) + strings.Repeat("\n", 2)
	var chart, row string
	var max = findMax(bc.series)

	for i := math.Round(max/bc.GetTick()) * bc.GetTick(); i > 0.0; i -= bc.GetTick() {

		// add ticks
		row = White.String() + alignRight(fmt.Sprintf("%.2f| ", i), " ", 8)

		// fill for each
		for _, s := range bc.series {
			if s.GetValue() >= i {
				row += s.GetColor() + alignLeft(s.GetFill(), " ", 2)
			} else {
				row += alignLeft("", " ", 3)
			}
		}
		chart += row + "\n"
	}

	row = White.String() + alignLeft("-", "---", len(bc.series))
	chart += alignRight(row, " ", 6)
	var longestName = findLongest(bc.series)

	for i := 0; i < longestName; i++ {
		row = alignLeft("\n", " ", 8)

		// add label vertically
		for _, s := range bc.series {
			name := s.GetLabel()

			if i < len(name) {
				row += s.GetColor() + alignLeft(string(name[i]), " ", 2)
			} else {
				row += alignLeft("", " ", 3)
			}
		}
		chart += row
	}
	return title + chart + White.String()
}

func (bc barChart) generateHorizontally() string {
	var title = center(bc.title, " ", len(bc.series)*2) + strings.Repeat("\n", 2)
	var longestName = findLongest(bc.series)
	var max = findMax(bc.series)
	var chart, row string

	for _, s := range bc.series {
		row = strings.Repeat(s.GetFill(), int(s.GetValue()/bc.GetTick()))
		chart += s.GetColor() + alignLeft(s.GetLabel(), " ", longestName) + " | "
		chart += alignLeft(row, " ", longestName+2+int(max/bc.GetTick())) + fmt.Sprintf("%.2f\n", s.GetValue())
	}
	return title + chart + White.String()
}
