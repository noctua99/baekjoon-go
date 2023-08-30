package stack

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("queue: Pop() called on empty queue")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		panic("queue: Peek() called on empty queue")
	}
	return (*s)[len(*s)-1]
}
