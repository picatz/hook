package main

import (
	"github.com/picatz/hook/pkg/call/http"
	"github.com/picatz/hook/pkg/call/log"
	"github.com/picatz/hook/pkg/types/action"
)

func main() {
	http.OnRequestHeaders(func(int, bool) action.Type {
		// remember to configure the httpbin upstream service so Envoy knows about it
		// https://github.com/proxy-wasm/proxy-wasm-rust-sdk/issues/8#issuecomment-630236035
		log.Info("dispatching external http call to httpbin")
		callID, err := http.Request(
			"httpbin",
			http.WithMethod("GET"),
			http.WithPath("/bytes/1"),
			http.WithAuthority("/bytes/1"),
		)

		//
		// The same dispatched call, but the longer form.
		//
		// callID, err := http.DispatchCall(
		// 	"httpbin",
		// 	[][2]string{
		// 		{":method", "GET"},
		// 		{":path", "/bytes/1"},
		// 		{":authority", "httpbin.org"},
		// 	},
		// 	"",            // emtpy body
		// 	[][2]string{}, // no trailers
		// 	30000,         // 30s
		// 	func(numHeaders int, bodySize int, numTrailers int) {
		// 		log.Infof("num headers %d", numHeaders)
		// 		log.Infof("body size %d", bodySize)
		// 		log.Infof("num trailers %d", numTrailers)
		// 	},
		// )

		if err != nil {
			log.Errorf("failed to register request: %v", err)
			return action.Pause
		}

		log.Infof("registered request: %v", callID)
		return action.Continue
	})
}
