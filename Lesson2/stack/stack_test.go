package stack

import "testing"

func TestNewStackOnSlice(t *testing.T) {
	stack := NewStackOnSlice()

	if stack == nil {
		t.Fatal("NewStackOnSlice() returned it nil")
	}
}

func TestStackPush(t *testing.T) {
	stack := NewStackOnSlice()

	stack.Push(10)
	if stack.size != 1 {
		t.Errorf("after Push(10), the size should be 1, got %d", stack.size)
	}

	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	if stack.size != 4 {
		t.Errorf("after Pushing several items, the size should be 4, got %d", stack.size)
	}

	expected := []int{10, 20, 30, 40}
	for i, v := range expected {
		if stack.items[i] != v {
			t.Errorf("expected stack.items[%d] = %d, got %d", i, v, stack.items[i])
		}
	}
	stack.Push(50)
	if stack.size != 5 {
		t.Errorf("after Push(50), the size should be 5, got %d", stack.size)
	}
	if stack.items[4] != 50 {
		t.Errorf("after Push(50), the fifth element should be 50, got %d", stack.items[4])
	}
}

func TestStackPop(t *testing.T) {
	stack := NewStackOnSlice()

	_, err := stack.Pop()
	if err == nil {
		t.Error("an error was expected when Poping from an empty stack, but the error was not received")
	}

	stack.Push(100)
	val, err := stack.Pop()
	if err != nil {
		t.Errorf("not expected error on Pop, got %v", err)
	}
	if val != 100 {
		t.Errorf("expected Pop() = 100, got %d", val)
	}
	if !stack.Empty() {
		t.Error("the stack should be empty after Pop")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val, err = stack.Pop()
	if err != nil {
		t.Errorf("not expected error on Pop, got %v", err)
	}
	if val != 3 {
		t.Errorf("expected Pop() = 3, got %d", val)
	}

	val, err = stack.Pop()
	if err != nil {
		t.Errorf("not expected error on Pop, got %v", err)
	}
	if val != 2 {
		t.Errorf("expected Pop() = 2, got %d", val)
	}
	val, err = stack.Pop()
	if err != nil {
		t.Errorf("not expected error on Pop, got %v", err)
	}
	if val != 1 {
		t.Errorf("expected Pop() = 1, got %d", val)
	}

	if !stack.Empty() {
		t.Error("stack should be empty after Pop")
	}
}

func TestStackEmpty(t *testing.T) {
	stack := NewStackOnSlice()

	if !stack.Empty() {
		t.Error("new stack should be empty")
	}

	stack.Push(5)
	if stack.Empty() {
		t.Error("stack should not be empty after a Push")
	}

	stack.Pop()
	if !stack.Empty() {
		t.Error("stack should be empty after Pop")
	}
}

func TestStackPrint(t *testing.T) {
	stack := NewStackOnSlice()

	expected := ""
	result := stack.Print()
	if result != expected {
		t.Errorf("expected Print() = \"%s\", got \"%s\"", expected, result)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expected = "3 2 1 "
	result = stack.Print()
	if result != expected {
		t.Errorf("expected Print() = \"%s\", got \"%s\"", expected, result)
	}

	stack.Pop()
	expected = "2 1 "
	result = stack.Print()
	if result != expected {
		t.Errorf("expected Print() = \"%s\", got \"%s\"", expected, result)
	}
}

func TestStackincreaseSlice(t *testing.T) {
	stack := NewStackOnSlice()

	if stack.size != 0 {
		t.Errorf("expected size = 0, got %d", stack.size)
	}
	if len(stack.items) != 0 {
		t.Errorf("expected len(items) = 0, got %d", len(stack.items))
	}

	for i := 1; i <= 10; i++ {
		stack.Push(i)
	}

	if stack.size != 10 {
		t.Errorf("expected size = 10, got %d", stack.size)
	}
	if len(stack.items) < 10 {
		t.Errorf("expected length items >= 10, got %d", len(stack.items))
	}

	for i := 0; i < 10; i++ {
		if stack.items[i] != i+1 {
			t.Errorf("expected items[%d] = %d, got %d", i, i+1, stack.items[i])
		}
	}
}
