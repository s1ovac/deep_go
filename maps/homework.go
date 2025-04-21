package maps

import "cmp"

type Node[T cmp.Ordered] struct {
	Key   T
	Value T

	left   *Node[T]
	right  *Node[T]
	parent *Node[T]
}

type OrderedMap[T cmp.Ordered] struct {
	*Node[T]
}

func NewOrderedMap[T cmp.Ordered]() OrderedMap[T] {
	return OrderedMap[T]{
		Node: new(Node[T]),
	}
}

func (m *OrderedMap[T]) Insert(key, value T) {
	var y *Node[T]
	for m.parent != nil {
		if m.parent.Key == key {
			m.parent.Value = value

			return
		}

		y = m.parent

		if m.parent.Key > key {
			m.parent = m.left
		} else {
			m.parent = m.right
		}
	}

	newNode := new(Node[T])
	if y == nil {
		m.parent = newNode

		return
	}

	if y.Key > key {
		y.left = newNode
	} else {
		y.right = newNode
	}
}

func (m *OrderedMap[T]) Erase(key int) {
	// need to implement
}

func (m *OrderedMap[T]) Contains(key int) bool {
	return false // need to implement
}

func (m *OrderedMap[T]) Size() int {
	return 0
}

func (m *OrderedMap[T]) ForEach(action func(int, int)) {
	// need to implement
}
