package main

import (
	"github.com/picatz/hook/pkg/call/http"
	"github.com/picatz/hook/pkg/types/action"
)

func main() {
	http.OnResponseHeaders(func(int, bool) action.Type {
		http.SetResponseHeader("X-XSS-Protection", "1; mode=block")
		http.SetResponseHeader("X-Frame-Options", "SAMEORIGIN")
		http.SetResponseHeader("X-Content-Type-Options", "nosniff")
		http.SetResponseHeader("X-Download-Options", "noopen")
		http.SetResponseHeader("Strict-Transport-Security", "max-age=5184000")
		return action.Continue
	})
}
