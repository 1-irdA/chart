package chart

// barchart is just a barchart with
// a title to describe it
// a tick unit (ex: 0.5)
// labels to describe values
// values
// fill bar with custom char
type barChart struct {
	title  string
	ticks  float64
	labels []string
	values []float64
	fill   string
}
