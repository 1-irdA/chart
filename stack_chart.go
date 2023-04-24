package chart

import (
	"fmt"
	"strings"
)

// NewStackChart Create a barchart instance
func NewStackChart(title string, tick float64, series []Serie) Chart {
	return &stackChart{title, tick, series}
}

// Generate generate horizontal or vertical cli barchart
func (s stackChart) Generate(chartType int) string {
	return s.generateHorizontally()
}

func (s stackChart) generateHorizontally() string {
	var size int = len(s.series)
	var title = White.String() + center(s.title, " ", size*2) + strings.Repeat("\n", 2)
	var legend, chart string

	for _, serie := range s.series {
		chart += serie.GetColor() + strings.Repeat(serie.GetFill(), int(serie.GetValue()/s.GetTick()))
	}
	for _, serie := range s.series {
		legend += serie.GetColor() + alignLeft(fmt.Sprintf("%s %s : %.1f", serie.GetFill(), serie.GetLabel(), serie.GetValue()), " ", 3)
	}
	return title + strings.Repeat(chart+"\n", 3) + "\n" + legend + White.String()
}
