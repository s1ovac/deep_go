package slices

// go test -v homework_test.go

type Numbers interface {
	int | int8 | int16 | int32 | int64
}

type CircularQueue[T Numbers] struct {
	values []T
	front  int
	rear   int
}

func NewCircularQueue[T Numbers](size int) CircularQueue[T] {
	return CircularQueue[T]{
		values: make([]T, size),
		front:  -1,
		rear:   -1,
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	if q.Empty() {
		q.front = 0
	}

	q.rear = (q.rear + 1) % len(q.values)
	q.values[q.rear] = value

	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}

	if q.front == q.rear {
		q.front = -1
		q.rear = -1

		return true
	}

	q.front = (q.front + 1) % len(q.values)

	return true
}

func (q *CircularQueue[T]) Front() T {
	if q.Empty() {
		return -1
	}

	return q.values[q.front]
}

func (q *CircularQueue[T]) Back() T {
	if q.Empty() {
		return -1
	}

	return q.values[q.rear]
}

func (q *CircularQueue[T]) Empty() bool {
	return q.front == -1
}

func (q *CircularQueue[T]) Full() bool {
	return q.front == (q.rear+1%len(q.values)) || q.front == 0 && q.rear == len(q.values)-1
}
