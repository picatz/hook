package data

import (
	"github.com/picatz/hook/pkg/call/host"
	"github.com/picatz/hook/pkg/call/utils"
	"github.com/picatz/hook/pkg/types/status"
)

func Get(key string) (value []byte, cas uint32, err error) {
	var raw *byte
	var size int

	st := host.ProxyGetSharedData(utils.StringToBytePtr(key), len(key), &raw, &size, &cas)
	if st != status.OK {
		return nil, 0, status.AsError(st)
	}
	return utils.BytePtrToByteSlice(raw, size), cas, nil
}


func Set(key string, data []byte, cas uint32) error {
	st := host.ProxySetSharedData(utils.StringToBytePtr(key), len(key), &data[0], len(data), cas)
	return status.AsError(st)
}
