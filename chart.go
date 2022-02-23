package chart

// barchart is just a barchart with
// a title to describe it
//  a tick unit (ex: 0.5)
// labels
// values
// char to fill bar
type barChart struct {
	title  string
	ticks  float64
	labels []string
	values []float64
	fill   string
}
