package status

type Type uint32

const (
	OK              Type = 0
	NotFound        Type = 1
	BadArgument     Type = 2
	Empty           Type = 7
	CasMismatch     Type = 8
	InternalFailure Type = 10
)
