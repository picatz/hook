package main

import (
	"strings"

	"github.com/picatz/hook/pkg/call/http"
	"github.com/picatz/hook/pkg/call/log"
	"github.com/picatz/hook/pkg/types/action"
)

func internalError() action.Type {
	http.SendResponse(
		500,
		[][2]string{{"reason", "internal-error"}},
		"Internal Error\n",
	)
	return action.Pause
}

func rejectRequest() action.Type {
	http.SendResponse(
		401,
		[][2]string{{"reason", "no-token"}},
		"Unauthorized\n",
	)
	return action.Pause
}

func authorizedToken() bool {
	// fake token auth logic
	header, err := http.GetRequestHeader("token")
	return err == nil && strings.TrimSpace(header) != ""
}

func stripToken() error {
	return http.RemoveRequestHeader("token")
}

func main() {
	http.OnRequestHeaders(func(int, bool) action.Type {
		if !authorizedToken() {
			log.Errorf("unauthorized access detected")
			return rejectRequest()
		}

		err := stripToken()
		if err != nil {
			log.Errorf("failed to strip token header from request after auth", err)
			return internalError()
		}

		return action.Continue
	})
}
