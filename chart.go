package chart

// type chart for struct who implement theses methods
type Chart interface {
	New(title string, ticks float64, labels []string, values []string, fill string) *Chart
	Generate(chartType int) string
}

// barchart is just a barchart with
// a title to describe it
//  a tick unit (ex: 0.5)
// labels
// values
// char to fill bar
type BarChart struct {
	title  string
	ticks  float64
	labels []string
	values []float64
	fill   string
}
