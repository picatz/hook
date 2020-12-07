package main

import (
	"fmt"

	"github.com/picatz/hook/pkg/call/http"
	"github.com/picatz/hook/pkg/call/log"
	"github.com/picatz/hook/pkg/call/vm/state"
	"github.com/picatz/hook/pkg/types/action"
	"github.com/picatz/hook/pkg/types/context"
)

func main() {
	state.SetNewHttpContext(newReplacer)
}

type replacer struct {
	context.HTTPDefault
	replaceWith string
}

func newReplacer(rootContextID, contextID uint32) context.HTTP {
	return &replacer{replaceWith: "replaced with wasm\n"}
}

func (ctx *replacer) OnHTTPResponseHeaders(size int, end bool) action.Type {
	http.SetResponseHeader("content-length", fmt.Sprintf("%d", len(ctx.replaceWith)))
	return action.Continue
}

func (ctx *replacer) OnHTTPResponseBody(size int, end bool) action.Type {
	if size > 0 {
		err := http.SetResponseBody([]byte(ctx.replaceWith), size) // replace using old size
		if err != nil {
			log.Errorf("failed to set http response body: %v", err)
			return action.Continue
		}
	}
	return action.Continue
}
