package workerPool

import (
	"errors"
	"sync/atomic"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	const tasksCount = 100
	var ranCount int32

	tasks := make([]Task, tasksCount)
	for i := 0; i < tasksCount; i++ {
		tasks[i] = func() error {
			atomic.AddInt32(&ranCount, 1)
			return nil
		}
	}

	err := Run(tasks, 5, tasksCount)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if atomic.LoadInt32(&ranCount) != tasksCount {
		t.Errorf("expected ranCount %d, got %d", tasksCount, ranCount)
	}
}

func TestRunNoErrorLimit(t *testing.T) {
	const tasksCount = 10
	var ranCount int32

	tasks := make([]Task, tasksCount)
	for i := 0; i < tasksCount; i++ {
		tasks[i] = func() error {
			atomic.AddInt32(&ranCount, 1)
			return errors.New("task error")
		}
	}

	err := Run(tasks, 3, 0)
	if err != nil {
		t.Fatalf("expected nil error when m<=0, got %v", err)
	}
	if atomic.LoadInt32(&ranCount) != tasksCount {
		t.Errorf("expected ranCount %d, got %d", tasksCount, ranCount)
	}
}
