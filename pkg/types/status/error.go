package status

import (
	"errors"
	"fmt"
)

var (
	ErrorNotFound        = errors.New("not found")
	ErrorBadArgument     = errors.New("bad argument")
	ErrorEmpty           = errors.New("empty")
	ErrorCasMismatch     = errors.New("cas mismatch")
	ErrorInternalFailure = errors.New("internal failure")
)

//go:inline
func AsError(status Type) error {
	switch status {
	case OK:
		return nil
	case NotFound:
		return ErrorNotFound
	case BadArgument:
		return ErrorBadArgument
	case Empty:
		return ErrorEmpty
	case CasMismatch:
		return ErrorCasMismatch
	case InternalFailure:
		return ErrorInternalFailure
	default:
		return errors.New(fmt.Sprintf("unkown error, %d is invalid status type", status))
	}
}
