package workerPool

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		m = len(tasks) + 1
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tasksCh := make(chan Task)
	var wg sync.WaitGroup
	var errCount int32

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case task, ok := <-tasksCh:
					if !ok {
						return
					}
					if err := task(); err != nil {
						if atomic.AddInt32(&errCount, 1) >= int32(m) {
							cancel()
						}
					}
				}
			}
		}()
	}

	go func() {
		defer close(tasksCh)
		for _, task := range tasks {
			select {
			case <-ctx.Done():
				return
			case tasksCh <- task:
			}
		}
	}()

	wg.Wait()

	if int(atomic.LoadInt32(&errCount)) >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
