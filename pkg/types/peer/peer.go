package peer

type Type uint32

const (
	Unknown Type = 0
	Local   Type = 1
	Remote  Type = 2
)
