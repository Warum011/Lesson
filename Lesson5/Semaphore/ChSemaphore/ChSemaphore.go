package main

import (
	"context"
	"fmt"
	"time"
)

type ChanSemaphore struct {
	ch  chan struct{}
	max int
}

func NewChanSemaphore(max int) *ChanSemaphore {
	s := &ChanSemaphore{
		ch:  make(chan struct{}, max),
		max: int(max),
	}
	for i := 0; i < max; i++ {
		s.ch <- struct{}{}
	}
	return s
}

func (s *ChanSemaphore) Acquire(ctx context.Context, n int) error {
	for range n {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-s.ch:
		}
	}
	return nil
}

func (s *ChanSemaphore) TryAcquire(n int) bool {
	for i := 0; i < n; i++ {
		select {
		case <-s.ch:
		default:
			s.Release(i)
			return false
		}
	}
	return true
}

func (s *ChanSemaphore) Release(n int) error {

	if (len(s.ch) + n) <= s.max {
		for range n {
			s.ch <- struct{}{}
		}
	} else {
		return fmt.Errorf("the channel is full, the operation is not completed")
	}
	return nil
}

func main() {
	sem := NewChanSemaphore(10)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := sem.Acquire(ctx, 3); err != nil {
		fmt.Println("acquire error:", err)
		return
	}

	err := sem.Release(3)
	if err != nil {
		fmt.Println(err)
	}

	success := sem.TryAcquire(5)
	if success {
		fmt.Println("successfully acquired 5 tokens using TryAcquire")
		sem.Release(5)
	} else {
		fmt.Println("failed to acquire 5 tokens immediately")
	}
}
