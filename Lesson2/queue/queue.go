package queue

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

type QueueOnSlice struct {
	items []int
}

func NewQueueOnSlice() *QueueOnSlice {
	return &QueueOnSlice{
		items: make([]int, 0),
	}
}

func (q *QueueOnSlice) PrintQueue() string {
	result := ""
	for _, item := range q.items {
		result += fmt.Sprintf("%d", item)
	}
	return result
}

func (q *QueueOnSlice) EmptyQueue() bool {
	return len(q.items) == 0
}

func (q *QueueOnSlice) PushQueue(item int) {
	q.items = append(q.items, item)
}

func (q *QueueOnSlice) PopQueue() (int, error) {
	if len(q.items) == 0 {
		return 0, ErrEmptyQueue
	}
	lastItem := q.items[0]
	q.items = append([]int(nil), q.items[1:]...)
	return lastItem, nil
}
