package stack

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type StackOnSlice struct {
	items []int
	size  int
}

func NewStackOnSlice() *StackOnSlice {
	return &StackOnSlice{
		items: make([]int, 0),
		size:  0,
	}
}

func (s *StackOnSlice) Print() string {
	result := ""
	for i := s.size - 1; i >= 0; i-- {
		result += fmt.Sprintf("%d ", s.items[i])
	}
	return result
}

func (s *StackOnSlice) Empty() bool {
	return s.size == 0
}

func (s *StackOnSlice) increaseSlice() {
	s.items = append(s.items, 0)
}

func (s *StackOnSlice) Push(item int) {
	if s.size == len(s.items) {
		s.increaseSlice()
	}
	s.items[s.size] = item
	s.size++
}

func (s *StackOnSlice) Pop() (int, error) {
	if s.size == 0 {
		return 0, ErrEmptyStack
	}
	s.size--
	return s.items[s.size], nil
}
