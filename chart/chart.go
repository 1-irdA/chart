package chart

type Chart interface {
	New(title string, ticks float64, labels []string, values []string, fill string) *Chart
	Generate(chartType int) string
}

type BarChart struct {
	title  string
	ticks  float64
	labels []string
	values []float64
	fill   string
}
