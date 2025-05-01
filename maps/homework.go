package maps

import "cmp"

type node[K cmp.Ordered, V any] struct {
	key   K
	value V

	left  *node[K, V]
	right *node[K, V]
}

type OrderedMap[K cmp.Ordered, V any] struct {
	head *node[K, V]

	len int
}

func NewOrderedMap[K cmp.Ordered, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{}
}

// Insert - добавить элемент в словарь
func (m *OrderedMap[K, V]) Insert(key K, value V) bool {
	isNew := false
	m.head = m.insertNode(m.head, key, value, &isNew)

	if isNew {
		m.len++
	}

	return isNew
}

func (m *OrderedMap[K, V]) insertNode(n *node[K, V], key K, value V, isNew *bool) *node[K, V] {
	if n == nil {
		*isNew = true

		return &node[K, V]{
			key:   key,
			value: value,
		}
	}

	if key == n.key {
		n.value = value
		return n
	}

	if key < n.key {
		n.left = m.insertNode(n.left, key, value, isNew)

		return n
	}

	n.right = m.insertNode(n.right, key, value, isNew)

	return n
}

// Erase - удалить элемент из словаря
func (m *OrderedMap[K, V]) Erase(key K) bool {
	if m.head == nil {
		return false
	}

	found := false
	m.head = m.eraseNode(m.head, key, &found)

	if found {
		m.len--
	}

	return found
}

func (m *OrderedMap[K, V]) eraseNode(n *node[K, V], key K, found *bool) *node[K, V] {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = m.eraseNode(n.left, key, found)

		return n
	}

	if key > n.key {
		n.right = m.eraseNode(n.right, key, found)
		return n

	}

	*found = true

	if n.left == nil && n.right == nil {
		return nil
	}

	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	successor := n.right
	for successor.left != nil {
		successor = successor.left
	}

	n.key = successor.key
	n.value = successor.value

	n.right = m.eraseNode(n.right, successor.key, found)

	return n
}

// Contains - проверить существование элемента в словаре
func (m *OrderedMap[K, V]) Contains(key K) bool {
	current := m.head

	for current != nil {
		if key == current.key {
			return true
		}

		if key < current.key {
			current = current.left

			continue
		}

		current = current.right
	}

	return false
}

// Size - получить количество элементов в словаре
func (m *OrderedMap[K, V]) Size() int {
	return m.len
}

// ForEach - применить функцию к каждому элементу словаря от меньшего к большему
func (m *OrderedMap[K, V]) ForEach(action func(K, V)) {
	stack := make([]*node[K, V], 0, m.len)
	current := m.head

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		action(current.key, current.value)

		current = current.right
	}
}
