package metric

type Type uint32

const (
	Counter   Type = 0
	Gauge     Type = 1
	Histogram Type = 2
)

