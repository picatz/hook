package main

import (
	"github.com/picatz/hook/pkg/call/http"
	"github.com/picatz/hook/pkg/types/action"
)

func main() {
	http.OnRequestHeaders(func(int, bool) action.Type {
		http.SetRequestHeader("wasm", "enabled")
		return action.Continue
	})

	http.OnResponseHeaders(func(int, bool) action.Type {
		http.SetResponseHeader("wasm", "enabled")
		return action.Continue
	})
}
