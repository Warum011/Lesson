package main

import (
	"context"
	"fmt"
	"sync"
)

type CondSemaphore struct {
	mu        sync.RWMutex
	available int
	max       int
}

func (s *CondSemaphore) Acquire(ctx context.Context, n int) error {

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			s.mu.RLock()
			if s.available >= n {
				s.mu.RUnlock()
				s.mu.Lock()
				if s.available >= n {
					s.available -= n
					s.mu.Unlock()
					return nil
				} else {
					s.mu.Unlock()
				}
			} else {
				s.mu.RUnlock()
			}
		}
	}
}

func (s *CondSemaphore) TryAcquire(n int) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	defer s.mu.Unlock()

	if s.available >= n {
		s.mu.RUnlock()
		s.mu.Lock()
		if s.available >= n {
			s.available -= n
			s.mu.Unlock()
			return true
		}
	}
	return false
}

func (s *CondSemaphore) Release(n int) error {
	s.mu.RLock()
	if s.max >= (n + s.available) {
		s.mu.RUnlock()
		s.mu.Lock()
		if s.max >= (n + s.available) {
			s.available += n
			return nil
		}
	}
	return fmt.Errorf("the channel is full, the operation is not completed")
}

func main() {

}
