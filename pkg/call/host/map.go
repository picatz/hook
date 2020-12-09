package host

import (
	"github.com/picatz/hook/pkg/call/utils"
	"github.com/picatz/hook/pkg/types/status"
	"github.com/picatz/hook/pkg/types/wmap"
)

func RemoveMapValue(mapType wmap.Type, key string) status.Type {
	return ProxyRemoveHeaderMapValue(mapType, utils.StringToBytePtr(key), len(key))
}

func SetMapValue(mapType wmap.Type, key, value string) status.Type {
	return ProxyReplaceHeaderMapValue(mapType, utils.StringToBytePtr(key), len(key), utils.StringToBytePtr(value), len(value))
}

func AddMapValue(mapType wmap.Type, key, value string) status.Type {
	return ProxyAddHeaderMapValue(mapType, utils.StringToBytePtr(key), len(key), utils.StringToBytePtr(value), len(value))
}

func GetMapValue(mapType wmap.Type, key string) (string, status.Type) {
	var rvs int
	var raw *byte
	if st := ProxyGetHeaderMapValue(mapType, utils.StringToBytePtr(key), len(key), &raw, &rvs); st != status.OK {
		return "", st
	}

	ret := utils.BytePtrToString(raw, rvs)
	return ret, status.OK
}

func SetMap(mapType wmap.Type, headers [][2]string) status.Type {
	shs := utils.HeadersToBytes(headers)
	hp := &shs[0]
	hl := len(shs)
	return ProxySetHeaderMapPairs(mapType, hp, hl)
}

func GetMap(mapType wmap.Type) ([][2]string, status.Type) {
	var rvs int
	var raw *byte

	st := ProxyGetHeaderMapPairs(mapType, &raw, &rvs)
	if st != status.OK {
		return nil, st
	}

	bs := utils.BytePtrToByteSlice(raw, rvs)

	return utils.BytesToHeaders(bs), status.OK
}
