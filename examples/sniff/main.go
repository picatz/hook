package main

import (
	"github.com/picatz/hook/pkg/call/http"
	"github.com/picatz/hook/pkg/call/log"
	"github.com/picatz/hook/pkg/types/action"
)

func main() {
	http.OnRequestHeaders(func(int, bool) action.Type {
		headers, err := http.GetRequestHeaders()
		if err != nil {
			log.Errorf("failed to get request headers: %v", err)
			return action.Continue
		}
		log.Infof("request headers sniffed: %v", headers)
		return action.Continue
	})

	http.OnRequestBody(func(size int, end bool) action.Type {
		if size > 0 {
			body, err := http.GetRequestBody(0, size)
			if err != nil {
				log.Errorf("failed to get request body: %v", err)
			} else {
				log.Infof("request body sniffed: %v", string(body))
			}
		}
		return action.Continue
	})

	http.OnResponseHeaders(func(int, bool) action.Type {
		headers, err := http.GetResponseHeaders()
		if err != nil {
			log.Errorf("failed to get response headers: %v", err)
			return action.Continue
		}
		log.Infof("request headers sniffed: %v", headers)
		return action.Continue
	})

	http.OnResponseBody(func(size int, end bool) action.Type {
		if size > 0 {
			body, err := http.GetResponseBody(0, size)
			if err != nil {
				log.Errorf("failed to get response body: %v", err)
			} else {
				log.Infof("response body sniffed: %v", string(body))
			}
		}
		return action.Continue
	})
}
