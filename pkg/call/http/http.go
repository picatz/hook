package http

import (
	"github.com/picatz/hook/pkg/call/host"
	"github.com/picatz/hook/pkg/call/utils"
	"github.com/picatz/hook/pkg/call/vm/state"
	"github.com/picatz/hook/pkg/types/buffer"
	"github.com/picatz/hook/pkg/types/http/callout"
	"github.com/picatz/hook/pkg/types/status"
	"github.com/picatz/hook/pkg/types/stream"
	"github.com/picatz/hook/pkg/types/wmap"
)

func SendResponse(statusCode uint32, headers [][2]string, body string) status.Type {
	shs := utils.HeadersToBytes(headers)
	hp := &shs[0]
	hl := len(shs)
	return host.ProxySendLocalResponse(statusCode, nil, 0, utils.StringToBytePtr(body), len(body), hp, hl, -1)
}

func GetCallResponseHeaders() ([][2]string, error) {
	ret, st := host.GetMap(wmap.HTTPCallResponseHeaders)
	return ret, status.AsError(st)
}

func GetCallResponseBody(start, maxSize int) ([]byte, error) {
	ret, st := host.GetBuffer(buffer.HTTPCallResponseBody, start, maxSize)
	return ret, status.AsError(st)
}

func GetCallResponseTrailers() ([][2]string, error) {
	ret, st := host.GetMap(wmap.HTTPCallResponseTrailers)
	return ret, status.AsError(st)
}

func GetDownstreamData(start, maxSize int) ([]byte, error) {
	ret, st := host.GetBuffer(buffer.DownstreamData, start, maxSize)
	return ret, status.AsError(st)
}

func GetUpstreamData(start, maxSize int) ([]byte, error) {
	ret, st := host.GetBuffer(buffer.DownstreamData, start, maxSize)
	return ret, status.AsError(st)
}

func GetRequestHeaders() ([][2]string, error) {
	ret, st := host.GetMap(wmap.HTTPRequestHeaders)
	return ret, status.AsError(st)
}

func SetRequestHeaders(headers [][2]string) error {
	return status.AsError(host.SetMap(wmap.HTTPRequestHeaders, headers))
}

func GetRequestHeader(key string) (string, error) {
	ret, st := host.GetMapValue(wmap.HTTPRequestHeaders, key)
	return ret, status.AsError(st)
}

func RemoveRequestHeader(key string) error {
	return status.AsError(host.RemoveMapValue(wmap.HTTPRequestHeaders, key))
}

func SetRequestHeader(key, value string) error {
	return status.AsError(host.SetMapValue(wmap.HTTPRequestHeaders, key, value))
}

func AddHTTPRequestHeader(key, value string) error {
	return status.AsError(host.AddMapValue(wmap.HTTPRequestHeaders, key, value))
}

func GetRequestBody(start, maxSize int) ([]byte, error) {
	ret, st := host.GetBuffer(buffer.HTTPRequestBody, start, maxSize)
	return ret, status.AsError(st)
}

func SetRequestBody(body []byte) error {
	var bufferData *byte
	if len(body) != 0 {
		bufferData = &body[0]
	}
	st := host.ProxySetBufferBytes(buffer.HTTPRequestBody, 0, len(body), bufferData, len(body))
	return status.AsError(st)
}

func GetRequestTrailers() ([][2]string, error) {
	ret, st := host.GetMap(wmap.HTTPRequestTrailers)
	return ret, status.AsError(st)
}

func SetRequestTrailers(headers [][2]string) error {
	return status.AsError(host.SetMap(wmap.HTTPRequestTrailers, headers))
}

func GetRequestTrailer(key string) (string, error) {
	ret, st := host.GetMapValue(wmap.HTTPRequestTrailers, key)
	return ret, status.AsError(st)
}

func RemoveRequestTrailer(key string) error {
	return status.AsError(host.RemoveMapValue(wmap.HTTPRequestTrailers, key))
}

func SetRequestTrailer(key, value string) error {
	return status.AsError(host.SetMapValue(wmap.HTTPRequestTrailers, key, value))
}

func AddRequestTrailer(key, value string) error {
	return status.AsError(host.AddMapValue(wmap.HTTPRequestTrailers, key, value))
}

func ResumeRequest() error {
	return status.AsError(host.ProxyContinueStream(stream.Request))
}

func GetResponseHeaders() ([][2]string, error) {
	ret, st := host.GetMap(wmap.HTTPResponseHeaders)
	return ret, status.AsError(st)
}

func SetResponseHeaders(headers [][2]string) error {
	return status.AsError(host.SetMap(wmap.HTTPResponseHeaders, headers))
}

func GetResponseHeader(key string) (string, error) {
	ret, st := host.GetMapValue(wmap.HTTPResponseHeaders, key)
	return ret, status.AsError(st)
}

func RemoveResponseHeader(key string) error {
	return status.AsError(host.RemoveMapValue(wmap.HTTPResponseHeaders, key))
}

func SetResponseHeader(key, value string) error {
	return status.AsError(host.SetMapValue(wmap.HTTPResponseHeaders, key, value))
}

func AddResponseHeader(key, value string) error {
	return status.AsError(host.AddMapValue(wmap.HTTPResponseHeaders, key, value))
}

func GetResponseBody(start, maxSize int) ([]byte, error) {
	ret, st := host.GetBuffer(buffer.HTTPResponseBody, start, maxSize)
	return ret, status.AsError(st)
}

func SetResponseBody(body []byte, maxSize int) error {
	var bufferData *byte
	if len(body) != 0 {
		bufferData = &body[0]
	}
	st := host.ProxySetBufferBytes(buffer.HTTPResponseBody, 0, maxSize, bufferData, len(body))
	return status.AsError(st)
}

func GetResponseTrailers() ([][2]string, error) {
	ret, st := host.GetMap(wmap.HTTPResponseTrailers)
	return ret, status.AsError(st)
}

func SetResponseTrailers(headers [][2]string) error {
	return status.AsError(host.SetMap(wmap.HTTPResponseTrailers, headers))
}

func GetResponseTrailer(key string) (string, error) {
	ret, st := host.GetMapValue(wmap.HTTPResponseTrailers, key)
	return ret, status.AsError(st)
}

func RemoveResponseTrailer(key string) error {
	return status.AsError(host.RemoveMapValue(wmap.HTTPResponseTrailers, key))
}

func SetResponseTrailer(key, value string) error {
	return status.AsError(host.SetMapValue(wmap.HTTPResponseTrailers, key, value))
}

func AddResponseTrailer(key, value string) error {
	return status.AsError(host.AddMapValue(wmap.HTTPResponseTrailers, key, value))
}

func ResumeResponse() error {
	return status.AsError(host.ProxyContinueStream(stream.Response))
}

func DispatchCall(
	upstream string,
	headers [][2]string,
	body string,
	trailers [][2]string,
	timeoutMillisecond uint32,
	callBack callout.Callback,
) (calloutID uint32, err error) {
	shs := utils.HeadersToBytes(headers)
	hp := &shs[0]
	hl := len(shs)

	sts := utils.HeadersToBytes(trailers)
	tp := &sts[0]
	tl := len(sts)

	u := utils.StringToBytePtr(upstream)
	st := host.ProxyHTTPCall(
		u,
		len(upstream),
		hp,
		hl,
		utils.StringToBytePtr(body),
		len(body),
		tp,
		tl,
		timeoutMillisecond,
		&calloutID,
	)

	switch st {
	case status.OK:
		state.RegisterHTTPCallout(calloutID, callBack)
		return calloutID, nil
	default:
		return 0, status.AsError(st)
	}
}
