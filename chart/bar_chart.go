package chart

import (
	"fmt"
	"os"
)

const (
	Vertical   = 1
	Horizontal = 2
)

func New(title string, ticks float64, labels []string, values []float64, fill string) *BarChart {
	if len(labels) != len(values) {
		fmt.Println(fmt.Sprintf("Needs same number of values : %d, labels : %d", len(values), len(labels)))
		os.Exit(1)
	}
	return &BarChart{title: title, ticks: ticks, labels: labels, values: values, fill: fill}
}

func (b BarChart) Generate(chartType int) string {
	if chartType == 1 {
		return b.generateVertically()
	}
	return b.generateHorizontally()
}

func (b BarChart) generateVertically() string {
	var title string = center(b.title, " ", len(b.labels)*2) + "\n"
	var chart string
	var row string

	for i := 100.0; i > 0; i -= b.ticks {

		// add tickss
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
	var longuestName int = findLonguest(b.labels)

	for i := 0; i < longuestName; i++ {
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

func (b BarChart) generateHorizontally() string {
	var title string = center(b.title, " ", len(b.labels)*2) + "\n\n"
	var longuestName int = findLonguest(b.labels)
	var max float64 = findMax(b.values)
	var chart string

	for i := 0; i < len(b.labels); i++ {
		row := ""
		chart += ajustRight(fmt.Sprintf("%s| ", b.labels[i]), " ", longuestName+2)

		for j := 100.0; j > 0.0; j -= b.ticks {

			if b.values[i] >= j {
				row += alignLeft(b.fill, " ", 2)
			}
		}
		chart += ajustLeft(row, " ", longuestName+2+int(max/b.ticks)*3) + fmt.Sprintf("%.2f", b.values[i]) + " %\n"
	}
	return title + chart
}
