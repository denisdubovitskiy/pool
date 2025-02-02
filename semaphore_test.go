package pool

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSemaphore(t *testing.T) {
	t.Run("Take", func(t *testing.T) {
		const (
			wantCount   = 10
			parallelism = 10_000
		)

		var (
			sem   = New(wantCount)
			count = 0
			wg    sync.WaitGroup
		)

		wg.Add(parallelism)

		for i := 0; i < parallelism; i++ {
			go func() {
				defer wg.Done()

				if ok := sem.Take(); ok {
					count++
				}
			}()
		}

		wg.Wait()

		require.Equal(t, wantCount, count)
		require.Equal(t, wantCount, sem.Taken())
	})

	t.Run("TakeAll", func(t *testing.T) {
		const (
			wantCount = 10
		)

		var (
			sem = New(wantCount)
		)

		ok := sem.Take()

		require.True(t, ok)
		require.Equal(t, 1, sem.Taken())

		sem.TakeAll()

		ok = sem.Take()
		require.False(t, ok)
		require.Equal(t, wantCount, sem.Taken())
	})

	t.Run("TakeAll", func(t *testing.T) {
		const (
			wantCount = 1
		)

		var (
			sem = New(wantCount)
		)

		ok := sem.Take()

		require.True(t, ok)
		require.Equal(t, 1, sem.Taken())

		sem.Release()

		require.Equal(t, 0, sem.Taken())
	})
}
