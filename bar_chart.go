package chart

import (
	"fmt"
	"log"
	"strings"
)

// NewBarChart Create a barchart instance
func NewBarChart(title string, ticks float64, labels []string, values []float64, fill string) Chart {
	lenLabels := len(labels)
	lenValues := len(values)

	if lenLabels != lenValues {
		log.Fatalf("Needs same number of values : %d, labels : %d", lenValues, lenLabels)
	}
	return &barChart{title: title, ticks: ticks, labels: labels, values: values, fill: fill, colors: make([]AnsiColor, lenLabels)}
}

// NewColoredBarChart Create a barchart instance
func NewColoredBarChart(title string, ticks float64, labels []string, values []float64, colors []AnsiColor, fill string) Chart {
	lenLabels := len(labels)
	lenValues := len(values)
	lenColors := len(colors)

	if lenLabels != lenValues || lenLabels != lenColors {
		log.Fatalf("Needs same number of values : %d, labels : %d, colors : %d", lenValues, lenLabels, lenColors)
	}
	return &barChart{title: title, ticks: ticks, labels: labels, values: values, fill: fill, colors: colors}
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
	var chart, row string
	var max = findMax(b.values)

	for i := max; i > 0.0; i -= b.ticks {

		// add ticks
		row = White.String() + ajustRight(fmt.Sprintf("%.2f| ", i), " ", 8)

		// fill for each
		for index, value := range b.values {
			if value >= i {
				row += b.colors[index].String() + alignLeft(b.fill, " ", 2)
			} else {
				row += alignLeft("", " ", 3)
			}
		}
		chart += row + "\n"
	}

	row = White.String() + alignLeft("-", "---", len(b.labels))
	chart += alignRight(row, " ", 6)
	var longestName = findLonguest(b.labels)

	for i := 0; i < longestName; i++ {
		row = alignLeft("\n", " ", 8)

		// add label vertically
		for index, name := range b.labels {
			if i < len(name) {
				row += b.colors[index].String() + alignLeft(string(name[i]), " ", 2)
			} else {
				row += alignLeft("", " ", 3)
			}
		}
		chart += row
	}
	return title + chart + White.String()
}

func (b barChart) generateHorizontally() string {
	var title = center(b.title, " ", len(b.labels)*2) + "\n\n"
	var longestName = findLonguest(b.labels)
	var max = findMax(b.values)
	var chart, row string

	for i := 0; i < len(b.labels); i++ {
		row = strings.Repeat(b.fill, int(b.values[i]/b.ticks))
		chart += b.colors[i].String() + ajustLeft(b.labels[i], " ", longestName) + " | "
		chart += ajustLeft(row, " ", longestName+2+int(max/b.ticks)) + fmt.Sprintf("%.2f\n", b.values[i])
	}
	return title + chart + White.String()
}
