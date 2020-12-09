package action

type Type uint32

type (
	General      = Type
	HeaderAction = Type
	DataAction   = Type
)

const (
	Continue General = 0
	Pause    General = 1
)

const (
	ContinueAndEndStream    HeaderAction = 2
	PauseAndBufferHeader    HeaderAction = 3
	PauseAndWatermarkHeader HeaderAction = 4
)

const (
	StopIterationAndBufferData       DataAction = 1
	StopIterationAndWatermarkData    DataAction = 2
	StopAllIterationAndWatermarkData DataAction = 4
)
