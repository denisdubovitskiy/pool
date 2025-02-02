package pool

import (
	"net"
	"time"
)

type Conn struct {
	net.Conn
	lastUsed  time.Time
	createdAt time.Time
	onClose   func()
}

func (c *Conn) Close() error {
	c.onClose()
	return c.Conn.Close()
}

func (c *Conn) markUsed() {
	c.lastUsed = time.Now()
}

func (c *Conn) age() time.Duration {
	return c.lastUsed.Sub(c.createdAt)
}
