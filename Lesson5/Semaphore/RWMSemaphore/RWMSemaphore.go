package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Semaphore interface {
	Acquire(ctx context.Context, n int) error
	Release(n int) error
	TryAcquire(n int) bool
}

type CondSemaphore struct {
	mu        sync.Mutex
	available int
	max       int
}

func NewCondSemaphore(max int) *CondSemaphore {
	return &CondSemaphore{
		available: max,
		max:       max,
	}
}

func (s *CondSemaphore) Acquire(ctx context.Context, n int) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		s.mu.Lock()
		if s.available >= n {
			s.available -= n
			s.mu.Unlock()
			return nil
		}
		s.mu.Unlock()
		time.Sleep(time.Millisecond)
	}
}

func (s *CondSemaphore) TryAcquire(n int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.available >= n {
		s.available -= n
		return true
	}
	return false
}

func (s *CondSemaphore) Release(n int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.available+n > s.max {
		return fmt.Errorf("release overflow: available=%d, release=%d, max=%d",
			s.available, n, s.max)
	}
	s.available += n
	return nil
}

func main() {
	sem := NewCondSemaphore(3)
	ctx := context.Background()

	if err := sem.Acquire(ctx, 2); err != nil {
		panic(err)
	}
	fmt.Println("Acquired 2, left:", sem.available)

	ok := sem.TryAcquire(2)
	fmt.Println("TryAcquire(2):", ok, "left:", sem.available)

	if err := sem.Release(2); err != nil {
		panic(err)
	}
	fmt.Println("Released 2, left:", sem.available)

	if err := sem.Release(5); err != nil {
		fmt.Println("Expected error:", err)
	}
}
