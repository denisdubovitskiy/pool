package pool

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
)

func TestPool(t *testing.T) {
	t.Run("put with no error", func(t *testing.T) {
		server := newTCPServer()
		addr, closeServer := server.Serve(t)
		t.Cleanup(closeServer)

		pool := NewPool(
			addr,
			WithMaxActiveConn(10),
			WithMaxIdleConn(10),
			WithMaxConnAge(5*time.Millisecond),
		)
		t.Cleanup(pool.Close)

		var wg sync.WaitGroup

		for i := 0; i < 1000; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				time.Sleep(5 * time.Millisecond)

				conn, err := pool.Get(context.Background())
				assert.NoError(t, err)
				assert.NotNil(t, conn)

				data := []byte(fmt.Sprintf(
					"%s%s\n",
					randstr.String(1),
					randstr.String(int(rand.Int32N(200))),
				))

				n, err := conn.Write(data)
				assert.NoError(t, err)
				assert.Equal(t, len(data), n)

				recv := make([]byte, len(data))

				n, err = conn.Read(recv)
				assert.NoError(t, err)
				assert.Equal(t, len(recv), n)

				pool.Put(conn, nil)
			}()
		}

		wg.Wait()
	})

	t.Run("put with error", func(t *testing.T) {
		server := newTCPServer()
		addr, closeServer := server.Serve(t)
		t.Cleanup(closeServer)

		pool := NewPool(
			addr,
			WithMaxActiveConn(10),
			WithMaxIdleConn(10),
			WithMaxConnAge(100*time.Millisecond),
		)
		t.Cleanup(pool.Close)

		var wg sync.WaitGroup

		for i := 0; i < 1000; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				conn, err := pool.Get(context.Background())
				assert.NoError(t, err)
				assert.NotNil(t, conn)

				data := []byte(fmt.Sprintf(
					"%s%s\n",
					randstr.String(1),
					randstr.String(int(rand.Int32N(200))),
				))

				n, err := conn.Write(data)
				assert.NoError(t, err)
				assert.Equal(t, len(data), n)

				recv := make([]byte, len(data))

				n, err = conn.Read(recv)
				assert.NoError(t, err)
				assert.Equal(t, len(recv), n)

				pool.Put(conn, assert.AnError)
			}()
		}

		wg.Wait()
	})

	t.Run("test idle timeout", func(t *testing.T) {
		server := newTCPServer()
		addr, closeServer := server.Serve(t)
		t.Cleanup(closeServer)

		pool := NewPool(
			addr,
			WithMaxActiveConn(10),
			WithMaxIdleConn(10),
			WithMaxConnAge(10*time.Second),
			WithIdleTimeout(10*time.Millisecond),
		)
		t.Cleanup(pool.Close)

		var wg sync.WaitGroup
		wg.Add(10)

		for i := 0; i < 10; i++ {

			go func() {
				defer wg.Done()

				time.Sleep(5 * time.Millisecond)

				conn, err := pool.Get(context.Background())
				assert.NoError(t, err)
				assert.NotNil(t, conn)

				data := []byte(fmt.Sprintf(
					"%s%s\n",
					randstr.String(1),
					randstr.String(int(rand.Int32N(200))),
				))

				n, err := conn.Write(data)
				assert.NoError(t, err)
				assert.Equal(t, len(data), n)

				recv := make([]byte, len(data))

				n, err = conn.Read(recv)
				assert.NoError(t, err)
				assert.Equal(t, len(recv), n)

				pool.Put(conn, nil)
			}()
		}

		wg.Wait()

		time.Sleep(100 * time.Millisecond)

		wg.Add(10)
		for i := 0; i < 10; i++ {

			go func() {
				defer wg.Done()

				time.Sleep(5 * time.Millisecond)

				conn, err := pool.Get(context.Background())
				assert.NoError(t, err)
				assert.NotNil(t, conn)

				data := []byte(fmt.Sprintf(
					"%s%s\n",
					randstr.String(1),
					randstr.String(int(rand.Int32N(200))),
				))

				n, err := conn.Write(data)
				assert.NoError(t, err)
				assert.Equal(t, len(data), n)

				recv := make([]byte, len(data))

				n, err = conn.Read(recv)
				assert.NoError(t, err)
				assert.Equal(t, len(recv), n)

				pool.Put(conn, nil)
			}()
		}

		wg.Wait()
	})

	t.Run("put after close", func(t *testing.T) {
		server := newTCPServer()
		addr, closeServer := server.Serve(t)
		t.Cleanup(closeServer)

		pool := NewPool(
			addr,
			WithMaxActiveConn(10),
			WithMaxIdleConn(10),
		)

		conn, err := pool.Get(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, conn)

		data := []byte(fmt.Sprintf(
			"%s%s\n",
			randstr.String(1),
			randstr.String(int(rand.Int32N(200))),
		))

		n, err := conn.Write(data)
		assert.NoError(t, err)
		assert.Equal(t, len(data), n)

		recv := make([]byte, len(data))

		n, err = conn.Read(recv)
		assert.NoError(t, err)
		assert.Equal(t, len(recv), n)

		pool.Close()

		time.Sleep(100 * time.Millisecond)

		pool.Put(conn, nil)
	})

	t.Run("get after close", func(t *testing.T) {
		server := newTCPServer()
		addr, closeServer := server.Serve(t)
		t.Cleanup(closeServer)

		pool := NewPool(
			addr,
			WithMaxActiveConn(10),
			WithMaxIdleConn(10),
		)

		pool.Close()

		_, err := pool.Get(context.Background())

		assert.Error(t, err)
		assert.Equal(t, ErrPoolClosed, err)
	})

	t.Run("stats", func(t *testing.T) {
		const capacity = 10

		server := newTCPServer()
		addr, closeServer := server.Serve(t)
		t.Cleanup(closeServer)

		pool := NewPool(
			addr,
			WithMaxActiveConn(capacity),
			WithMaxIdleConn(capacity),
		)

		funcs := make([]func(), 0, capacity)

		for {
			stats := pool.Stats()
			if stats.IdleConns == capacity {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}

		stats := pool.Stats()

		assert.Equal(t, capacity, stats.ActiveConns)
		assert.Equal(t, capacity, stats.IdleConns)
		assert.Equal(t, 0, stats.InUseConns)

		for i := 0; i < capacity; i++ {
			conn, err := pool.Get(context.Background())
			assert.NoError(t, err)
			assert.NotNil(t, conn)

			funcs = append(funcs, func() {
				pool.Put(conn, nil)
			})

			stats := pool.Stats()

			wantInUse := i + 1
			wantIdle := capacity - wantInUse

			assert.Equal(t, wantInUse, stats.InUseConns)
			assert.Equal(t, capacity, stats.ActiveConns)
			assert.Equal(t, wantIdle, stats.IdleConns)
		}

		for i, put := range funcs {
			put()

			wantIdle := i + 1
			wantInUse := capacity - wantIdle

			stats := pool.Stats()

			assert.Equal(t, stats.ActiveConns, capacity)
			assert.Equal(t, wantInUse, stats.InUseConns)
			assert.Equal(t, capacity, stats.ActiveConns)
		}

		pool.Close()

		time.Sleep(100 * time.Millisecond)
	})

	t.Run("dial", func(t *testing.T) {
		mc := minimock.NewController(t)
		timeout := time.Second
		mockConn := NewConnMock(mc)

		mockDialer := NewDialerMock(mc)
		mockDialer.DialTimeoutMock.Times(1).
			Expect("tcp", "127.0.0.1:8000", timeout).
			Return(mockConn, nil)

		pool := NewPool(
			"127.0.0.1:8000",
			WithDialer(mockDialer),
			WithMaxActiveConn(1),
			WithDialTimeout(timeout),
		)

		conn, err := pool.Get(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, conn)
	})

	t.Run("dial failed", func(t *testing.T) {
		mc := minimock.NewController(t)
		timeout := time.Second

		mockDialer := NewDialerMock(mc)
		mockDialer.DialTimeoutMock.
			Expect("tcp", "127.0.0.1:8000", timeout).
			Return(nil, assert.AnError)

		pool := NewPool(
			"127.0.0.1:8000",
			WithDialer(mockDialer),
			WithMaxActiveConn(1),
			WithDialTimeout(timeout),
		)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		conn, err := pool.Get(ctx)

		assert.Error(t, err)
		assert.Nil(t, conn)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func newTCPServer() *tcpServer {
	return &tcpServer{}
}

type tcpServer struct{}

func (s *tcpServer) Serve(t *testing.T) (string, func()) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			conn, err := lis.Accept()
			if err != nil {
				break
			}

			if conn == nil {
				log.Printf("accepted new connection")
				continue
			}

			wg.Add(1)

			go func() {
				defer wg.Done()

				s.serveConn(conn)
			}()
		}
	}()

	return lis.Addr().String(), func() {
		_ = lis.Close()
		wg.Wait()
	}
}

func (s *tcpServer) serveConn(conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(
		bufio.NewReader(conn),
		bufio.NewWriter(conn),
	)

	for {
		req, err := rw.ReadString('\n')
		if err != nil {
			_, _ = rw.WriteString("failed to read input")
			rw.Flush()
			return
		}

		_, _ = rw.WriteString(req)
		rw.Flush()
	}
}
