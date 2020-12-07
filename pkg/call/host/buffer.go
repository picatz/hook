package host

import (
	"github.com/picatz/hook/pkg/call/utils"
	"github.com/picatz/hook/pkg/types/buffer"
	"github.com/picatz/hook/pkg/types/status"
)

func GetBufferOrError(bufType buffer.Type, size int) ([]byte, error) {
	ret, st := GetBuffer(bufType, 0, size)
	return ret, status.AsError(st)
}

func GetBuffer(bufType buffer.Type, start, maxSize int) ([]byte, status.Type) {
	var (
		retData *byte
		retSize int
	)
	switch st := ProxyGetBufferBytes(bufType, start, maxSize, &retData, &retSize); st {
	case status.OK:
		return utils.BytePtrToByteSlice(retData, retSize), st
	default:
		return nil, st
	}
}
