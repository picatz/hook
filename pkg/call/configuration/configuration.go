package configuration

import (
	"github.com/picatz/hook/pkg/call/host"
	"github.com/picatz/hook/pkg/types/buffer"
)

func GetPlugin(size int) ([]byte, error) {
	return host.GetBufferOrError(buffer.PluginConfiguration, size)
}

func GetVM(size int) ([]byte, error) {
	return host.GetBufferOrError(buffer.PluginConfiguration, size)
}
