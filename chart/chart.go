package chart

type Chart interface {
	New(title string, tick float64, labels []string, values []string, fill string) *Chart
	Generate() string
}

type BarChart struct {
	title  string
	tick   float64
	labels []string
	values []float64
	fill   string
}
