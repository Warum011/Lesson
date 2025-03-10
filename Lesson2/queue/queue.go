package queue

import (
	"errors"
	"fmt"
)

type queueOnSlice struct {
	items []int
}

func NewQueueOnSlice() *queueOnSlice {
	return &queueOnSlice{
		items: make([]int, 0),
	}
}

func (q *queueOnSlice) PrintQueue() string {
	result := ""
	for _, item := range q.items {
		result += fmt.Sprintf("%d", item)
	}
	return result
}

func (q *queueOnSlice) EmptyQueue() bool {
	return len(q.items) == 0
}

func (q *queueOnSlice) PushQueue(item int) {
	q.items = append(q.items, item)
}

func (q *queueOnSlice) PopQueue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	lastItem := q.items[0]
	q.items = q.items[1:]
	return lastItem, nil
}
