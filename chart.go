package chart

const (
	Vertical   = 1
	Horizontal = 2
)

// Chart allow you to create a chart
type Chart interface {
	GetTick() float64
	GetSeries() []Series
	Generate(int) string
}

// Series is in chart
type Series interface {
	GetLabel() string
	GetValue() float64
	GetColor() string
	GetFill() string
}

type barChart struct {
	title  string
	tick   float64
	series []Series
}

func (bc barChart) GetSeries() []Series {
	return bc.series
}

func (bc barChart) GetTick() float64 {
	return bc.tick
}

type stackChart struct {
	title  string
	tick   float64
	series []Series
}

func (sc stackChart) GetSeries() []Series {
	return sc.series
}

func (sc stackChart) GetTick() float64 {
	return sc.tick
}

type series struct {
	label, fill string
	value       float64
	color       AnsiColor
}

// NewSeries create a new series
func NewSeries(label, fill string, value float64) Series {
	return &series{label, fill, value, White}
}

// NewColoredSeries create a new colored series
func NewColoredSeries(label, fill string, value float64, color AnsiColor) Series {
	return &series{label, fill, value, color}
}

func (s series) GetLabel() string {
	return s.label
}

func (s series) GetValue() float64 {
	return s.value
}

func (s series) GetFill() string {
	return s.fill
}

func (s series) GetColor() string {
	return s.color.String()
}
