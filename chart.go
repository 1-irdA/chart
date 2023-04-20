package chart

const (
	Vertical   = 1
	Horizontal = 2
)

// barChart is just a barchart with
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
	colors []AnsiColor
	fill   string
}

// stackChart is just a stackChart with
// a title to describe it
// labels to describe values
// values
// fills to fill stack
type stackChart struct {
	title  string
	ticks  float64
	labels []string
	values []float64
	fills  []string
	colors []AnsiColor
}

type Chart interface {
	Generate(int) string
}
