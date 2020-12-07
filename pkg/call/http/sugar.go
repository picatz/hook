package http

import (
	"github.com/picatz/hook/pkg/call/vm/state"
	"github.com/picatz/hook/pkg/types/action"
	"github.com/picatz/hook/pkg/types/context"
)

func init() {
	state.SetNewHttpContext(newInternalHTTPContext)
}

var internal = &httpContext{}

type Header = [2]string
type Headers = []Header

type httpContext struct {
	context.HTTPDefault

	// context IDs
	rootContextID uint32
	contextID     uint32

	// custom funcs via sugar methods
	onRequestHeaders   func(int, bool) action.Type
	onRequestBody      func(int, bool) action.Type
	onRequestTrailers  func(int) action.Type
	onResponseHeaders  func(int, bool) action.Type
	onResponseBody     func(int, bool) action.Type
	onResponseTrailers func(int) action.Type
	onStreamDone       func()
}

func newInternalHTTPContext(rootContextID, contextID uint32) context.HTTP {
	internal.rootContextID = rootContextID
	internal.contextID = contextID
	return internal
}

func (i *httpContext) OnRequestHeaders(size int, end bool) action.Type {
	if i.onRequestHeaders != nil {
		return i.onRequestHeaders(size, end)
	}
	return action.Continue
}

func (i *httpContext) OnRequestBody(size int, end bool) action.Type {
	if i.onRequestBody != nil {
		return i.onRequestBody(size, end)
	}
	return action.Continue
}

func (i *httpContext) OnRequestTrailers(size int) action.Type {
	if i.onRequestTrailers != nil {
		return i.onRequestTrailers(size)
	}
	return action.Continue
}

func (i *httpContext) OnResponseHeaders(size int, end bool) action.Type {
	if i.onResponseHeaders != nil {
		return i.onResponseHeaders(size, end)
	}
	return action.Continue
}

func (i *httpContext) OnResponseBody(size int, end bool) action.Type {
	if i.onResponseBody != nil {
		return i.onResponseBody(size, end)
	}
	return action.Continue
}

func (i *httpContext) OnResponseTrailers(size int) action.Type {
	if i.onResponseHeaders != nil {
		return i.onResponseTrailers(size)
	}
	return action.Continue
}

func (i *httpContext) OnStreamDone() {
	if i.onStreamDone != nil {
		i.onStreamDone()
	}
}

func OnRequestHeaders(f func(int, bool) action.Type) {
	internal.onRequestHeaders = f
}

func OnRequestBody(f func(int, bool) action.Type) {
	internal.onRequestBody = f
}

func OnRequestTrailers(f func(int) action.Type) {
	internal.onRequestTrailers = f
}

func OnResponseHeaders(f func(int, bool) action.Type) {
	internal.onResponseHeaders = f
}

func OnResponseBody(f func(int, bool) action.Type) {
	internal.onResponseBody = f
}

func OnResponseTrailers(f func(int) action.Type) {
	internal.onResponseTrailers = f
}

func OnStreamDone(f func()) {
	internal.onStreamDone = f
}
