package queue

import (
	"github.com/picatz/hook/pkg/call/host"
	"github.com/picatz/hook/pkg/call/utils"
	"github.com/picatz/hook/pkg/types/status"
)

func Register(name string) (uint32, error) {
	var queueID uint32
	ptr := utils.StringToBytePtr(name)
	st := host.ProxyRegisterSharedQueue(ptr, len(name), &queueID)
	return queueID, status.AsError(st)
}

// TODO: verify and fix this
func Resolve(vmID, queueName string) (uint32, error) {
	var ret uint32
	st := host.ProxyResolveSharedQueue(
		utils.StringToBytePtr(vmID),
		len(vmID),
		utils.StringToBytePtr(queueName),
		len(queueName),
		&ret,
	)
	return ret, status.AsError(st)
}

func Dequeue(queueID uint32) ([]byte, error) {
	var raw *byte
	var size int
	st := host.ProxyDequeueSharedQueue(queueID, &raw, &size)
	if st != status.OK {
		return nil, status.AsError(st)
	}
	return utils.BytePtrToByteSlice(raw, size), nil
}

func Enqueue(queueID uint32, data []byte) error {
	return status.AsError(host.ProxyEnqueueSharedQueue(queueID, &data[0], len(data)))
}
