package context

import (
	"github.com/picatz/hook/pkg/types/action"
	"github.com/picatz/hook/pkg/types/peer"
)

type Stream interface {
	OnDownstreamData(dataSize int, endOfStream bool) action.Type
	OnDownstreamClose(peerType peer.Type)
	OnNewConnection() action.Type
	OnUpstreamData(dataSize int, endOfStream bool) action.Type
	OnUpstreamClose(peerType peer.Type)
	OnStreamDone()
}

type StreamDefault struct{}

func (*StreamDefault) OnDownstreamData(int, bool) action.Type { return action.Continue }
func (*StreamDefault) OnDownstreamClose(peer.Type)            {}
func (*StreamDefault) OnNewConnection() action.Type           { return action.Continue }
func (*StreamDefault) OnUpstreamData(int, bool) action.Type   { return action.Continue }
func (*StreamDefault) OnUpstreamClose(peer.Type)              {}
func (*StreamDefault) OnStreamDone()                          {}
