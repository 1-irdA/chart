package main

import (
	"fmt"

	chart "github.com/garrou/chart/lib"
)

func main() {
	label := []string{"go", "cpp", "js", "java"}
	values := []float64{17, 14.95, 35.05, 54.65}
	colors := []chart.AnsiColor{chart.Yellow, chart.Red, chart.Green, chart.Blue}

	barChart := chart.NewBarChart("Languages", 5, label, values, colors, "=")
	result := barChart.Generate(chart.Horizontal)

	fmt.Println(result)
}
