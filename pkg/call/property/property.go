package property

import (
	"github.com/picatz/hook/pkg/call/host"
	"github.com/picatz/hook/pkg/call/utils"
	"github.com/picatz/hook/pkg/types/status"
)

func GetProperty(path []string) ([]byte, error) {
	var ret *byte
	var retSize int
	raw := utils.PropertyPathToBytes(path)

	err := status.AsError(host.ProxyGetProperty(&raw[0], len(raw), &ret, &retSize))
	if err != nil {
		return nil, err
	}

	return utils.BytePtrToByteSlice(ret, retSize), nil
}

func SetProperty(path string, data []byte) error {
	return status.AsError(host.ProxySetProperty(utils.StringToBytePtr(path), len(path), &data[0], len(data)))
}
