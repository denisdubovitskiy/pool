package pool

import (
	"sync/atomic"
)

func New(i int) *Semaphore {
	tokens := &atomic.Int32{}
	tokens.Store(0)
	return &Semaphore{
		max:    int32(i),
		tokens: tokens,
		closed: &atomic.Bool{},
	}
}

type Semaphore struct {
	tokens *atomic.Int32
	closed *atomic.Bool
	max    int32
}

func (s *Semaphore) Take() bool {
	for {
		oldVal := s.tokens.Load()
		newVal := oldVal + 1

		if newVal > s.max {
			return false
		}

		if s.tokens.CompareAndSwap(oldVal, newVal) {
			return true
		}
	}
}

func (s *Semaphore) TakeAll() {
	s.tokens.Store(s.max)
}

func (s *Semaphore) Release() {
	for {
		oldVal := s.tokens.Load()
		newVal := oldVal - 1

		if s.tokens.CompareAndSwap(oldVal, newVal) {
			return
		}
	}
}

func (s *Semaphore) Taken() int {
	return int(s.tokens.Load())
}
