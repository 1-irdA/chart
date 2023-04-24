package chart

import (
	"fmt"
	"log"
	"strings"
)

// NewStackChart Create a barchart instance
func NewStackChart(title string, ticks float64, labels []string, values []float64, fills []string) Chart {
	lenLabels := len(labels)
	lenValues := len(values)
	lenFills := len(fills)

	if lenLabels != lenValues || lenLabels != lenFills {
		log.Fatalf("Needs same number of values : %d, labels : %d, fills : %d", lenValues, lenLabels, lenFills)
	}
	return &stackChart{title: title, ticks: ticks, labels: labels, values: values, fills: fills, colors: make([]AnsiColor, lenFills)}
}

// NewColoredStackChart Create a barchart instance
func NewColoredStackChart(title string, ticks float64, labels []string, values []float64, fills []string, colors []AnsiColor) Chart {
	lenLabels := len(labels)
	lenValues := len(values)
	lenFills := len(fills)
	lenColors := len(colors)

	if lenLabels != lenValues || lenLabels != lenFills || lenLabels != lenColors {
		log.Fatalf("Needs same number of values : %d, labels : %d, fills : %d, colors : %d", lenValues, lenLabels, lenFills, lenColors)
	}
	return &stackChart{title: title, ticks: ticks, labels: labels, values: values, fills: fills, colors: colors}
}

// Generate generate horizontal or vertical cli barchart
func (s stackChart) Generate(chartType int) string {
	return s.generateHorizontally()
}

func (s stackChart) generateHorizontally() string {
	var size int = len(s.labels)
	var title = center(s.title, " ", size*2) + strings.Repeat("\n", 2)
	var chart string
	var legend = strings.Repeat("\n", 2)

	for i := 0; i < size; i++ {
		chart += s.colors[i].String() + strings.Repeat(s.fills[i], int(s.values[i]/s.ticks))
	}
	for i := 0; i < size; i++ {
		legend += s.colors[i].String() + alignLeft(fmt.Sprintf("%s %s : %.1f", s.fills[i], s.labels[i], s.values[i]), " ", 3)
	}
	return title + duplicateHorizontal(chart, 2) + legend + White.String()
}
