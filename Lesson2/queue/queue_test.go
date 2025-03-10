package queue

import "testing"

func TestNewQueueOnSlice(t *testing.T) {
	queue := NewQueueOnSlice()

	if queue == nil {
		t.Fatal("NewQueueOnSlice() returned it nil")
	}
	if len(queue.items) != 0 {
		t.Errorf("expected length items=0, got=%d", len(queue.items))
	}
}

func TestPushQueue(t *testing.T) {
	queue := NewQueueOnSlice()

	queue.PushQueue(10)
	if len(queue.items) != 1 {
		t.Errorf("after PushQueue(10), the size should be 1, got %d", len(queue.items))
	}

	queue.PushQueue(20)
	queue.PushQueue(30)
	queue.PushQueue(40)

	if len(queue.items) != 4 {
		t.Errorf("after Pushing several items, the size should be 4, got %d", len(queue.items))
	}

	expected := []int{10, 20, 30, 40}
	for i, v := range expected {
		if queue.items[i] != v {
			t.Errorf("expected queue.items[%d] = %d, got %d", i, v, queue.items[i])
		}
	}
	queue.PushQueue(50)
	if len(queue.items) != 5 {
		t.Errorf("after PushQueue(50), the size should be 5, got %d", len(queue.items))
	}
	if queue.items[4] != 50 {
		t.Errorf("after PushQueue(50), the fifth element should be 50, got %d", queue.items[4])
	}
}

func TestPopQueue(t *testing.T) {
	queue := NewQueueOnSlice()

	_, err := queue.PopQueue()
	if err == nil {
		t.Error("an error was expected when Poping from an empty queue, but the error was not received")
	}

	queue.PushQueue(100)
	val, err := queue.PopQueue()
	if err != nil {
		t.Errorf("not expected error on PopQueue, got %v", err)
	}
	if val != 100 {
		t.Errorf("expected PopQueue() = 100, got %d", val)
	}
	if !queue.EmptyQueue() {
		t.Error("the queue should be empty after PopQueue")
	}

	queue.PushQueue(1)
	queue.PushQueue(2)
	queue.PushQueue(3)

	val, err = queue.PopQueue()
	if err != nil {
		t.Errorf("not expected error on PopQueue, got %v", err)
	}
	if val != 1 {
		t.Errorf("expected PopQueue() = 1, got %d", val)
	}

	val, err = queue.PopQueue()
	if err != nil {
		t.Errorf("not expected error on PopQueue, got %v", err)
	}
	if val != 2 {
		t.Errorf("expected PopQueue() = 2, got %d", val)
	}
	val, err = queue.PopQueue()
	if err != nil {
		t.Errorf("not expected error on PopQueue, got %v", err)
	}
	if val != 3 {
		t.Errorf("expected PopQueue() = 3, got %d", val)
	}

	if !queue.EmptyQueue() {
		t.Error("queue should be empty after Pop")
	}
}

func TestEmptyQueue(t *testing.T) {
	queue := NewQueueOnSlice()

	if !queue.EmptyQueue() {
		t.Error("new queue should be empty")
	}

	queue.PushQueue(5)
	if queue.EmptyQueue() {
		t.Error("queue should not be empty after a Push")
	}

	queue.PopQueue()
	if !queue.EmptyQueue() {
		t.Error("queue should be empty after Pop")
	}
}

func TestSPrintQueue(t *testing.T) {
	stack := NewQueueOnSlice()

	expected := ""
	result := stack.PrintQueue()
	if result != expected {
		t.Errorf("expected PrintQueue() = \"%s\", got \"%s\"", expected, result)
	}

	stack.PushQueue(1)
	stack.PushQueue(2)
	stack.PushQueue(3)

	expected = "123"
	result = stack.PrintQueue()
	if result != expected {
		t.Errorf("expected PrintQueue() = \"%s\", got \"%s\"", expected, result)
	}

	stack.PopQueue()
	expected = "23"
	result = stack.PrintQueue()
	if result != expected {
		t.Errorf("expected PrintQueue() = \"%s\", got \"%s\"", expected, result)
	}
}
