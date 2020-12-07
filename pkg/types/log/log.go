package log

type Level uint32

const (
	Trace    Level = 0
	Debug    Level = 1
	Info     Level = 2
	Warn     Level = 3
	Error    Level = 4
	Critical Level = 5
)

func (l Level) String() string {
	switch l {
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Critical:
		return "critical"
	default:
		return "unkown"
	}
}
