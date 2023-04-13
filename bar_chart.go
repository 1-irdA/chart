package chart

import (
	"fmt"
	"os"
)

const (
	Vertical   = 1
	Horizontal = 2
)

type BarChart interface {
	Generate(int) string
}

// New Create a barchart instance
func New(title string, ticks float64, labels []string, values []float64, fill string) BarChart {
	if len(labels) != len(values) {
		fmt.Printf("%s", fmt.Sprintf("Needs same number of values : %d, labels : %d", len(values), len(labels)))
		os.Exit(1)
	}
	return &barChart{title: title, ticks: ticks, labels: labels, values: values, fill: fill}
}

// Generate generate horizontal or vertical cli barchart
func (b barChart) Generate(chartType int) string {
	if chartType == Vertical {
		return b.generateVertically()
	}
	return b.generateHorizontally()
}

func (b barChart) generateVertically() string {
	var title = center(b.title, " ", len(b.labels)*2) + "\n"
	var chart string
	var row string

	for i := 100.0; i > 0; i -= b.ticks {

		// add ticks
		row = ajustRight(fmt.Sprintf("%.2f| ", i), " ", 8)

		// fill for each
		for _, percent := range b.values {
			if percent >= i {
				row += alignLeft(b.fill, " ", 2)
			} else {
				row += alignLeft("", " ", 3)
			}
		}
		chart += row + "\n"
	}

	row = alignLeft("-", "---", len(b.labels))
	chart += alignRight(row, " ", 6)
	var longestName = findLonguest(b.labels)

	for i := 0; i < longestName; i++ {
		row = alignLeft("\n", " ", 8)

		// add label vertically
		for _, name := range b.labels {
			if i < len(name) {
				row += alignLeft(string(name[i]), " ", 2)
			} else {
				row += alignLeft("", " ", 3)
			}
		}
		chart += row
	}
	return title + chart
}

func (b barChart) generateHorizontally() string {
	var title = center(b.title, " ", len(b.labels)*2) + "\n\n"
	var longestName = findLonguest(b.labels)
	var max = findMax(b.values)
	var chart string

	for i := 0; i < len(b.labels); i++ {
		row := ""
		chart += ajustLeft(fmt.Sprintf("%s", b.labels[i]), " ", longestName) + " | "

		for j := 100.0; j > 0.0; j -= b.ticks {

			if b.values[i] >= j {
				row += b.fill
			}
		}
		chart += ajustLeft(row, " ", longestName+2+int(max/b.ticks)) + fmt.Sprintf("%.2f", b.values[i]) + " %\n"
	}
	return title + chart
}
