package chart

import (
	"fmt"
	"strings"
)

// NewStackChart Create a barchart instance
func NewStackChart(title string, tick float64, series []Series) Chart {
	return &stackChart{title, tick, series}
}

// Generate generate horizontal or vertical cli barchart
func (sc stackChart) Generate(chartType int) string {
	return sc.generateHorizontally()
}

func (sc stackChart) generateHorizontally() string {
	var size = len(sc.series)
	var title = White.String() + center(sc.title, " ", size*2) + strings.Repeat("\n", 2)
	var legend, chart string

	for _, s := range sc.series {
		chart += s.GetColor() + strings.Repeat(s.GetFill(), int(s.GetValue()/sc.GetTick()))
	}
	for _, s := range sc.series {
		legend += s.GetColor() + alignLeft(fmt.Sprintf("%s %s : %.1f", s.GetFill(), s.GetLabel(), s.GetValue()), " ", 3)
	}
	return title + strings.Repeat(chart+"\n", 3) + "\n" + legend + White.String()
}
