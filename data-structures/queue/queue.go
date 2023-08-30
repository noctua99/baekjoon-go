package queue

type Queue[T any] []T

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Pop() T {
	if len(*q) == 0 {
		panic("queue: PopFront() called on empty queue")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Peek() T {
	if len(*q) == 0 {
		panic("queue: PopFront() called on empty queue")
	}
	return (*q)[0]
}
