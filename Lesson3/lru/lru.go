package lru

type Node[KeyT comparable, ValueT any] struct {
	key   KeyT
	value ValueT
	prev  *Node[KeyT, ValueT]
	next  *Node[KeyT, ValueT]
}

type List[KeyT comparable, ValueT any] struct {
	head *Node[KeyT, ValueT]
	tail *Node[KeyT, ValueT]
}

func NewList[KeyT comparable, ValueT any]() *List[KeyT, ValueT] {
	var key KeyT
	var value ValueT
	list := &List[KeyT, ValueT]{
		head: &Node[KeyT, ValueT]{key, value, nil, nil},
		tail: &Node[KeyT, ValueT]{key, value, nil, nil},
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func (l *List[KeyT, ValueT]) PushToFront(node *Node[KeyT, ValueT]) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *List[KeyT, ValueT]) Remove(node *Node[KeyT, ValueT]) {
	if node == nil {
		return
	}
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (l *List[KeyT, ValueT]) MoveToFront(node *Node[KeyT, ValueT]) {
	if node == nil {
		return
	}
	l.Remove(node)
	l.PushToFront(node)
}

func (l *List[KeyT, ValueT]) Back() *Node[KeyT, ValueT] {
	if l.head == l.tail {
		return nil
	}
	return l.tail.prev
}

type LRUCache[KeyT comparable, ValueT any] struct {
	capacity int
	cache    map[KeyT]*Node[KeyT, ValueT]
	list     *List[KeyT, ValueT]
}

func NewLRUCache[KeyT comparable, ValueT any](capacity int) *LRUCache[KeyT, ValueT] {
	return &LRUCache[KeyT, ValueT]{
		capacity: capacity,
		cache:    make(map[KeyT]*Node[KeyT, ValueT]),
		list:     NewList[KeyT, ValueT](),
	}
}

func (lru *LRUCache[KeyT, ValueT]) Get(key KeyT) (ValueT, bool) {
	if node, found := lru.cache[key]; found {
		lru.list.MoveToFront(node)
		return node.value, true
	}
	var value ValueT
	return value, false
}

func (lru *LRUCache[KeyT, ValueT]) Put(key KeyT, value ValueT) {
	if node, found := lru.cache[key]; found {
		lru.list.MoveToFront(node)
		node.value = value
		return
	}
	if len(lru.cache) == lru.capacity {
		back := lru.list.Back()
		if back != nil {
			lru.list.Remove(back)
			delete(lru.cache, back.key)
		}
	}
	newNode := &Node[KeyT, ValueT]{key, value, nil, nil}
	lru.list.PushToFront(newNode)
	lru.cache[key] = newNode
}
