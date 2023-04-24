package chart

const (
	Vertical   = 1
	Horizontal = 2
)

type Chart interface {
	GetTick() float64
	GetSeries() []Serie
	Generate(int) string
}

type Serie interface {
	GetLabel() string
	GetValue() float64
	GetColor() string
	GetFill() string
}

// barChart is just a barchart with
// a title to describe it
// a tick unit (ex: 0.5)
// labels to describe values
// values
// fill bar with custom char
type barChart struct {
	title  string
	tick   float64
	series []Serie
}

func (b barChart) GetSeries() []Serie {
	return b.series
}

func (b barChart) GetTick() float64 {
	return b.tick
}

// stackChart is just a stackChart with
// a title to describe it
// labels to describe values
// values
// fills to fill stack
type stackChart struct {
	title  string
	tick   float64
	series []Serie
}

func (s stackChart) GetSeries() []Serie {
	return s.series
}

func (s stackChart) GetTick() float64 {
	return s.tick
}

// barSerie is just a chart series with
// a label to describe it
// a value
// a color
type serie struct {
	label, fill string
	value       float64
	color       AnsiColor
}

// NewSerie create a new bar serie
func NewSerie(label, fill string, value float64) *serie {
	return &serie{label, fill, value, White}
}

// NewColoredSerie create a new bar serie
func NewColoredSerie(label, fill string, value float64, color AnsiColor) *serie {
	return &serie{label, fill, value, color}
}

func (s serie) GetLabel() string {
	return s.label
}

func (s serie) GetValue() float64 {
	return s.value
}

func (s serie) GetFill() string {
	return s.fill
}

func (s serie) GetColor() string {
	return s.color.String()
}
