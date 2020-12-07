package buffer

type Type uint32

const (
	HTTPRequestBody      Type = 0
	HTTPResponseBody     Type = 1
	DownstreamData       Type = 2
	UpstreamData         Type = 3
	HTTPCallResponseBody Type = 4
	GRPCReceiveBuffer    Type = 5
	VMConfiguration      Type = 6
	PluginConfiguration  Type = 7
	CallData             Type = 8
)
