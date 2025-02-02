package pool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {
	t.Run("WithMaxActiveConn", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.Equal(t, defaultMaxActiveConn, got.maxActiveConn)
		})

		t.Run("override", func(t *testing.T) {
			want := 101

			// act
			got := newPoolOptions(WithMaxActiveConn(want))

			// assert
			require.Equal(t, want, got.maxActiveConn)
		})
	})

	t.Run("WithMaxIdleConn", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.Equal(t, defaultMaxIdleConn, got.maxIdleConn)
		})

		t.Run("override, defaults", func(t *testing.T) {
			want := 101

			// act
			got := newPoolOptions(WithMaxIdleConn(want))

			// assert
			require.Equal(t, defaultMaxActiveConn, got.maxIdleConn)
		})

		t.Run("override with active", func(t *testing.T) {
			wantActive := 101
			wantIdle := 51

			// act
			got := newPoolOptions(WithMaxActiveConn(wantActive), WithMaxIdleConn(wantIdle))

			// assert
			require.Equal(t, wantIdle, got.maxIdleConn)
			require.Equal(t, wantActive, got.maxActiveConn)
		})
	})

	t.Run("WithMaxConnAge", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.Equal(t, time.Duration(defaultMaxConnAge), got.maxConnAge)
		})

		t.Run("override", func(t *testing.T) {
			want := 101 * time.Second

			// act
			got := newPoolOptions(WithMaxConnAge(want))

			// assert
			require.Equal(t, want, got.maxConnAge)
		})
	})

	t.Run("WithIdleTimeout", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.Equal(t, time.Duration(defaultIdleTimeout), got.idleTimeout)
		})

		t.Run("override", func(t *testing.T) {
			want := 101 * time.Second

			// act
			got := newPoolOptions(WithIdleTimeout(want))

			// assert
			require.Equal(t, want, got.idleTimeout)
		})
	})

	t.Run("WithDialTimeout", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.Equal(t, time.Duration(defaultDialTimeout), got.dialTimeout)
		})

		t.Run("override", func(t *testing.T) {
			want := 101 * time.Second

			// act
			got := newPoolOptions(WithDialTimeout(want))

			// assert
			require.Equal(t, want, got.dialTimeout)
		})
	})

	t.Run("WithDialRetries", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.Equal(t, defaultDialRetries, got.dialRetries)
		})

		t.Run("override", func(t *testing.T) {
			want := 101

			// act
			got := newPoolOptions(WithDialRetries(want))

			// assert
			require.Equal(t, want, got.dialRetries)
		})
	})

	t.Run("WithDialer", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.Equal(t, newDefaultDialer(), got.dialer)
		})

		t.Run("override", func(t *testing.T) {
			want := NewDialerMock(t)

			// act
			got := newPoolOptions(WithDialer(want))

			// assert
			require.Equal(t, want, got.dialer)
		})
	})

	t.Run("WithOnCloseHook", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.NotPanics(t, func() {
				got.hooks.OnClose(time.Second, CloseReasonError, ErrConnFailed)
			})
		})

		t.Run("override", func(t *testing.T) {
			want := NewOnCloseHookMock(t)
			want.OnCloseMock.
				Times(1).
				Expect(time.Second, CloseReasonError, ErrConnFailed).
				Return()

			// act
			got := newPoolOptions(WithOnCloseHook(want))
			got.hooks.OnClose(time.Second, CloseReasonError, ErrConnFailed)
		})
	})

	t.Run("WithOnDialHook", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.NotPanics(t, func() {
				got.hooks.OnDial(time.Second, ErrConnFailed)
			})
		})

		t.Run("override", func(t *testing.T) {
			want := NewOnDialHookMock(t)
			want.OnDialMock.
				Times(1).
				Expect(time.Second, ErrConnFailed).
				Return()

			// act
			got := newPoolOptions(WithOnDialHook(want))
			got.hooks.OnDial(time.Second, ErrConnFailed)
		})
	})

	t.Run("WithOnGetTimeoutHook", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			// act
			got := newPoolOptions()

			// assert
			require.NotPanics(t, func() {
				got.hooks.OnDial(time.Second, ErrConnFailed)
			})
		})

		t.Run("override", func(t *testing.T) {
			want := NewOnGetTimeoutHookMock(t)
			want.OnGetTimeoutMock.
				Times(1).
				Expect(time.Second).
				Return()

			// act
			got := newPoolOptions(WithOnGetTimeoutHook(want))
			got.hooks.OnGetTimeout(time.Second)
		})
	})
}
