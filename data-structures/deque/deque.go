package deque

type Deque[T any] []T

func New[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (d *Deque[T]) PushBack(value T) {
	*d = append(*d, value)
}

func (d *Deque[T]) PushFront(value T) {
	*d = append([]T{value}, *d...)
}

func (d *Deque[T]) PopFront() T {
	if len(*d) == 0 {
		panic("queue: PopFront() called on empty queue")
	}
	v := (*d)[0]
	*d = (*d)[1:]
	return v
}

func (d *Deque[T]) PopBack() T {
	if len(*d) == 0 {
		panic("queue: PopBack() called on empty queue")
	}
	v := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return v
}

func (d *Deque[T]) Len() int {
	return len(*d)
}

func (d *Deque[T]) Empty() bool {
	return len(*d) == 0
}

func (d *Deque[T]) Front() T {
	if len(*d) == 0 {
		panic("queue: Front() called on empty queue")
	}
	return (*d)[0]
}

func (d *Deque[T]) Back() T {
	if len(*d) == 0 {
		panic("queue: Back() called on empty queue")
	}
	return (*d)[len(*d)-1]
}
