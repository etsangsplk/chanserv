// Package chanserv implements an effective 2-dimensional message queue based upon AstraNet multiplexer and discovery capabilities.
package chanserv

import (
	"net"
	"time"
)

type Frame interface {
	Bytes() []byte
}

type MetaData interface {
	RemoteAddr() string
}

type Source interface {
	Header() []byte
	Meta() MetaData
	Out() <-chan Frame
}

type SourceFunc func(reqBody []byte) <-chan Source

type Multiplexer interface {
	Bind(net, laddr string) (net.Listener, error)
	DialTimeout(network string, address string, timeout time.Duration) (net.Conn, error)
}

type Server interface {
	ListenAndServe(vAddr string, src SourceFunc) error
}

type RequestTag int

const (
	TagMeta RequestTag = iota
	TagBucket
)

type Client interface {
	LookupAndPost(vAddr string, body []byte, tags map[RequestTag]string) (<-chan Source, error)
}
