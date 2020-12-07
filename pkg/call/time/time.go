package time

import (
	"github.com/picatz/hook/pkg/call/host"
	"github.com/picatz/hook/pkg/types/status"
)

func GetCurrent() int64 {
	var t int64
	host.ProxyGetCurrentTimeNanoseconds(&t)
	return t
}

func SetTickPeriodMilliSeconds(millSec uint32) error {
	return status.AsError(host.ProxySetTickPeriodMilliseconds(millSec))
}
