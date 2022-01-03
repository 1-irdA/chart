package chart

import (
	"fmt"
	"os"
)

func New(title string, tick float64, labels []string, values []float64, fill string) *BarChart {
	if len(labels) != len(values) {
		fmt.Println(fmt.Sprintf("Needs same number of values : %d, labels : %d", len(values), len(labels)))
		os.Exit(1)
	}
	return &BarChart{title: title, tick: tick, labels: labels, values: values, fill: fill}
}

func (b *BarChart) Generate() string {
	var title string = center(b.title, " ", len(b.labels)*2) + "\n"
	var chart string
	var row string

	for i := 100.0; i > 0; i -= b.tick {

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
