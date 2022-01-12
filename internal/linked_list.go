package internal

type Node[V any] struct {
	prev  *Node[V]
	next  *Node[V]
	value V
}

type List[V any] struct {
	front *Node[V]
	back  *Node[V]
}

func (l *List[V]) PushBack(value V) *Node[V] {
	return l.pushBack(&Node[V]{
		prev:  nil,
		next:  nil,
		value: value,
	})
}

func (l *List[V]) PopFront() V {
	node := l.front
	l.front = node.next

	if l.front != nil {
		l.front.prev = nil
	}

	node.next = nil
	node.prev = nil
	return node.value
}

func (l *List[V]) MoveBack(node *Node[V]) {
	l.Remove(node)
	l.pushBack(node)
}

func (l *List[V]) Remove(node *Node[V]) {
	if l.front == node {
		l.front = node.next
		l.front.prev = nil
	}
	if l.back == node {
		l.back = node.prev
		l.back.next = nil
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.next = nil
	node.prev = nil
}

func (l *List[V]) pushBack(node *Node[V]) *Node[V] {
	if l.back == nil {
		l.front = node
		l.back = node
	} else {
		l.back.next = node
		node.prev = l.back
		l.back = node
	}

	return node
}
