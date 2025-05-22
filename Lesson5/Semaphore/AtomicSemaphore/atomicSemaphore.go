package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

type atomicSemaphore struct {
	available int64
	max       int64
}

func (s *atomicSemaphore) Acquire(ctx context.Context, n int64) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			current := atomic.LoadInt64(&s.available)
			if current < n {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			newVal := current - n
			if atomic.CompareAndSwapInt64(&s.available, current, newVal) {
				return nil
			}
		}
	}
}

func (s *atomicSemaphore) TryAcquire(n int64) bool {
	for {
		current := atomic.LoadInt64(&s.available)
		if current < n {
			return false
		}
		if atomic.CompareAndSwapInt64(&s.available, current, current-n) {
			return true
		}
	}
}

func (s *atomicSemaphore) Release(n int64) error {
	for {
		current := atomic.LoadInt64(&s.available)
		newVal := current + n
		if newVal > s.max {
			return fmt.Errorf("the semaphore is full, the operation is not completed")
		}
		if atomic.CompareAndSwapInt64(&s.available, current, newVal) {
			return nil
		}
	}
}

func main() {

}
