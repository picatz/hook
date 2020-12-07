package wmap

type Type uint32

const (
	HTTPRequestHeaders       Type = 0
	HTTPRequestTrailers      Type = 1
	HTTPResponseHeaders      Type = 2
	HTTPResponseTrailers     Type = 3
	HTTPCallResponseHeaders  Type = 6
	HTTPCallResponseTrailers Type = 7
)
