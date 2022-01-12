package cache

type node[V any] struct {
	prev  *node[V]
	next  *node[V]
	value V
}

type list[V any] struct {
	front *node[V]
	back  *node[V]
}

func (l *list[V]) pushBack(value V) {
	node := &node[V]{
		prev:  nil,
		next:  nil,
		value: value,
	}

	if l.back == nil {
		l.front = node
		l.back = node
	} else {
		l.back.next = node
		node.prev = l.back
		l.back = node
	}
}

func (l *list[V]) popFront() V {
	node := l.front
	l.front = node.next

	if l.front != nil {
		l.front.prev = nil
	}

	node.next = nil
	node.prev = nil
	return node.value
}
