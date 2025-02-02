package pool

import (
	"net"
	"time"
)

type Dialer interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

func newDefaultDialer() Dialer {
	return &defaultDialer{}
}

type defaultDialer struct{}

func (d *defaultDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout("tcp", address, timeout)
}
