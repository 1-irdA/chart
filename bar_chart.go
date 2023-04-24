package chart

import (
	"fmt"
	"math"
	"strings"
)

// NewBarChart Create a barchart instance
func NewBarChart(title string, tick float64, series []Serie) Chart {
	return &barChart{title, tick, series}
}

// Generate generate horizontal or vertical cli barchart
func (b barChart) Generate(chartType int) string {
	if chartType == Vertical {
		return b.generateVertically()
	}
	return b.generateHorizontally()
}

func (b barChart) generateVertically() string {
	var title = center(b.title, " ", len(b.series)*2) + strings.Repeat("\n", 2)
	var chart, row string
	var max = findMax(b.series)

	for i := (math.Round(max/b.GetTick()) * b.GetTick()); i > 0.0; i -= b.GetTick() {

		// add ticks
		row = White.String() + ajustRight(fmt.Sprintf("%.2f| ", i), " ", 8)

		// fill for each
		for _, serie := range b.series {
			if serie.GetValue() >= i {
				row += serie.GetColor() + alignLeft(serie.GetFill(), " ", 2)
			} else {
				row += alignLeft("", " ", 3)
			}
		}
		chart += row + "\n"
	}

	row = White.String() + alignLeft("-", "---", len(b.series))
	chart += alignRight(row, " ", 6)
	var longestName = findLonguest(b.series)

	for i := 0; i < longestName; i++ {
		row = alignLeft("\n", " ", 8)

		// add label vertically
		for _, serie := range b.series {
			name := serie.GetLabel()

			if i < len(name) {
				row += serie.GetColor() + alignLeft(string(name[i]), " ", 2)
			} else {
				row += alignLeft("", " ", 3)
			}
		}
		chart += row
	}
	return title + chart + White.String()
}

func (b barChart) generateHorizontally() string {
	var title = center(b.title, " ", len(b.series)*2) + strings.Repeat("\n", 2)
	var longestName = findLonguest(b.series)
	var max = findMax(b.series)
	var chart, row string

	for _, serie := range b.series {
		row = strings.Repeat(serie.GetFill(), int(serie.GetValue()/b.GetTick()))
		chart += serie.GetColor() + ajustLeft(serie.GetLabel(), " ", longestName) + " | "
		chart += ajustLeft(row, " ", longestName+2+int(max/b.GetTick())) + fmt.Sprintf("%.2f\n", serie.GetValue())
	}
	return title + chart + White.String()
}
