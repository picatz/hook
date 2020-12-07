package main

import (
	"github.com/picatz/hook/pkg/call/http"
	"github.com/picatz/hook/pkg/call/vm/state"
	"github.com/picatz/hook/pkg/types/action"
	"github.com/picatz/hook/pkg/types/context"
)

func main() {
	state.SetNewHttpContext(newHeader)
}

type header struct {
	context.HTTPDefault
	contextID uint32
}

func newHeader(rootContextID, contextID uint32) context.HTTP {
	return &header{contextID: contextID}
}

func (ctx *header) OnHTTPRequestHeaders(int, bool) action.Type {
	http.SetRequestHeader("wasm", "enabled")
	return action.Continue
}

func (ctx *header) OnHTTPResponseHeaders(int, bool) action.Type {
	http.SetResponseHeader("wasm", "enabled")
	return action.Continue
}
