package context

import (
	"github.com/picatz/hook/pkg/types/action"
)

type HTTP interface {
	OnRequestHeaders(numHeaders int, endOfStream bool) action.Type
	OnRequestBody(bodySize int, endOfStream bool) action.Type
	OnRequestTrailers(numTrailers int) action.Type
	OnResponseHeaders(numHeaders int, endOfStream bool) action.Type
	OnResponseBody(bodySize int, endOfStream bool) action.Type
	OnResponseTrailers(numTrailers int) action.Type
	OnStreamDone()
}

type HTTPDefault struct{}

func (*HTTPDefault) OnRequestHeaders(int, bool) action.Type  { return action.Continue }
func (*HTTPDefault) OnRequestBody(int, bool) action.Type     { return action.Continue }
func (*HTTPDefault) OnRequestTrailers(int) action.Type       { return action.Continue }
func (*HTTPDefault) OnResponseHeaders(int, bool) action.Type { return action.Continue }
func (*HTTPDefault) OnResponseBody(int, bool) action.Type    { return action.Continue }
func (*HTTPDefault) OnResponseTrailers(int) action.Type      { return action.Continue }
func (*HTTPDefault) OnStreamDone()                           {}
