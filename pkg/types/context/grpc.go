package context

import (
	"github.com/picatz/hook/pkg/types/action"
	"github.com/picatz/hook/pkg/types/peer"
)

type GRPC interface {
	OnDownstreamData(dataSize int, endOfStream bool) action.Type
	OnDownstreamClose(peerType peer.Type)
	OnNewConnection() action.Type
	OnUpstreamData(dataSize int, endOfStream bool) action.Type
	OnUpstreamClose(peerType peer.Type)
	OnRequestHeaders(numHeaders int, endOfStream bool) action.Type
	OnRequestMetadata(size int) action.Type
	OnRequestBody(bodySize int, endOfStream bool) action.Type
	OnRequestTrailers(numTrailers int) action.Type
	OnResponseHeaders(numHeaders int, endOfStream bool) action.Type
	OnResponseBody(bodySize int, endOfStream bool) action.Type
	OnResponseMetadata(size int) action.Type
	OnResponseTrailers(numTrailers int) action.Type
	OnStreamDone()
}

type GRPCDefault struct{}

func (*GRPCDefault) OnRequestHeaders(int, bool) action.Type  { return action.Continue }
func (*GRPCDefault) OnRequestMetadata(int) action.Type       { return action.Continue }
func (*GRPCDefault) OnRequestBody(int, bool) action.Type     { return action.Continue }
func (*GRPCDefault) OnRequestTrailers(int) action.Type       { return action.Continue }
func (*GRPCDefault) OnResponseHeaders(int, bool) action.Type { return action.Continue }
func (*GRPCDefault) OnResponseMetadata(int) action.Type      { return action.Continue }
func (*GRPCDefault) OnResponseBody(int, bool) action.Type    { return action.Continue }
func (*GRPCDefault) OnResponseTrailers(int) action.Type      { return action.Continue }
func (*GRPCDefault) OnDownstreamData(int, bool) action.Type  { return action.Continue }
func (*GRPCDefault) OnDownstreamClose(peer.Type)             {}
func (*GRPCDefault) OnNewConnection() action.Type            { return action.Continue }
func (*GRPCDefault) OnUpstreamData(int, bool) action.Type    { return action.Continue }
func (*GRPCDefault) OnUpstreamClose(peer.Type)               {}
func (*GRPCDefault) OnStreamDone()                           {}
