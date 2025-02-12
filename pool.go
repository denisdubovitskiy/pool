package pool

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Stats struct {
	InUseConns  int
	ActiveConns int
	IdleConns   int
}

type Pool struct {
	addr          string
	maxActiveConn int
	maxIdleConn   int
	maxConnAge    time.Duration
	idleTimeout   time.Duration
	dialTimeout   time.Duration
	dialRetries   int
	retryInterval time.Duration
	conns         chan *Conn
	activeConns   *atomic.Int32
	closing       atomic.Bool
	wg            sync.WaitGroup
	semaphore     *Semaphore
	dialer        Dialer
	hooks         *hooks
}

var (
	ErrPoolClosed = errors.New("pool: pool is closed")
	ErrConnFailed = errors.New("pool: conn failed")
)

func NewPool(addr string, opts ...Option) *Pool {
	opt := newPoolOptions(opts...)

	p := &Pool{
		addr:          addr,
		maxActiveConn: opt.maxActiveConn,
		maxIdleConn:   opt.maxIdleConn,
		dialRetries:   opt.dialRetries,
		conns:         make(chan *Conn, opt.maxActiveConn),
		semaphore:     New(opt.maxActiveConn),
		activeConns:   &atomic.Int32{},
		idleTimeout:   opt.idleTimeout,
		dialTimeout:   opt.dialTimeout,
		dialer:        opt.dialer,
		hooks:         opt.hooks,
	}

	p.startMaintainingConnections()

	return p
}

func (p *Pool) startMaintainingConnections() {
	p.wg.Add(1)

	go func() {
		defer p.wg.Done()

		p.maintainConnections()
	}()
}

func (p *Pool) newConnection(conn net.Conn) *Conn {
	p.activeConns.Add(1)
	return &Conn{
		Conn:      conn,
		createdAt: time.Now(),
		lastUsed:  time.Now(),
		onClose: func() {
			p.activeConns.Add(-1)
			p.semaphore.Release()
		},
	}
}

func (p *Pool) dial() (*Conn, error) {
	var (
		err  error
		conn net.Conn
	)

	// Initial dial and retries
	totalRetries := 1 + p.dialRetries

	for i := 0; i < totalRetries; i++ {
		dialStarted := time.Now()
		conn, err = p.dialer.DialTimeout("tcp", p.addr, p.dialTimeout)
		p.hooks.OnDial(time.Since(dialStarted), err)

		if err != nil {
			time.Sleep(p.retryInterval)
			continue
		}
		log.Println("New connection established:", p.addr)

		return p.newConnection(conn), nil
	}

	if err == nil {
		err = ErrConnFailed
	}

	return nil, err
}

func (p *Pool) Get(ctx context.Context) (*Conn, error) {
	if p.closing.Load() {
		return nil, ErrPoolClosed
	}

	start := time.Now()

	for !p.closing.Load() {
		select {
		case <-ctx.Done():
			p.hooks.OnGetTimeout(time.Since(start))
			return nil, fmt.Errorf(
				"pool: failed to get a connection from the pool: %w",
				ctx.Err(),
			)

		case conn := <-p.conns:
			if conn == nil {
				return nil, ErrPoolClosed
			}

			connAge := conn.age()

			if p.maxConnAge > 0 && connAge > p.maxConnAge {
				p.hooks.OnClose(connAge, CloseReasonMaxAgeExceeded, conn.Close())
				continue
			}

			if p.idleTimeout > 0 && time.Since(conn.lastUsed) > p.idleTimeout {
				p.hooks.OnClose(connAge, CloseReasonIdleTimeoutExceeded, conn.Close())
				continue
			}

			return conn, nil
		}
	}

	return nil, ErrPoolClosed
}

func (p *Pool) Put(conn *Conn, err error) {
	if err != nil {
		p.hooks.OnClose(conn.age(), CloseReasonError, conn.Close())
		return
	}

	if p.closing.Load() {
		p.hooks.OnClose(conn.age(), CloseReasonPoolClosing, conn.Close())
		return
	}

	conn.markUsed()

	select {
	case p.conns <- conn:

	default:
		// This code will unlikely be executed due to semaphore.
		p.hooks.OnClose(
			conn.age(),
			CloseReasonPoolCapacityExceeded,
			conn.Close(),
		)
	}
}

func (p *Pool) Close() {
	// Prohibiting new outgoing connections.
	p.semaphore.TakeAll()

	// Signal that the pool is going to be closed and finishing the loop of
	// filling a pool with new connections.
	if !p.closing.CompareAndSwap(false, true) {
		return
	}

	close(p.conns)

	for conn := range p.conns {
		if conn == nil {
			continue
		}

		p.hooks.OnClose(
			conn.age(),
			CloseReasonPoolClosing,
			conn.Close(),
		)
	}

	// Waiting background tasks to be finished.
	p.wg.Wait()
}

func (p *Pool) maintainConnections() {
	for !p.closing.Load() {
		// If the Take() call returns false, creating new outgoing connections
		// is prohibited because the pool is full or new connections are already
		// being created.
		if ok := p.semaphore.Take(); !ok {
			runtime.Gosched()
			continue
		}

		err := p.createConnection()

		// Immediately release the token if the connection was not created due
		// to a host connection error.
		if err != nil && !errors.Is(err, ErrPoolClosed) {
			p.semaphore.Release()
		}
	}
}

func (p *Pool) createConnection() error {
	// Just sanity check.
	if p.closing.Load() {
		return ErrPoolClosed
	}

	conn, err := p.dial()
	if err != nil {
		return err
	}

	// Creating a connection takes time, so while the connection was being
	// created, the pool may have already been closed.
	if p.closing.Load() {
		p.hooks.OnClose(
			conn.age(),
			CloseReasonPoolClosing,
			conn.Close(),
		)
		return ErrPoolClosed
	}

	select {
	case p.conns <- conn:
	default:
		// This code will unlikely be executed due to semaphore.
		p.hooks.OnClose(
			conn.age(),
			CloseReasonPoolCapacityExceeded,
			conn.Close(),
		)
	}

	return nil
}

func (p *Pool) Stats() Stats {
	// Non-atomic yet.
	poolConns := len(p.conns)
	activeConns := int(p.activeConns.Load())
	return Stats{
		InUseConns:  activeConns - poolConns,
		ActiveConns: int(p.activeConns.Load()),
		IdleConns:   poolConns,
	}
}
