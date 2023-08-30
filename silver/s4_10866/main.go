package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// fastIO
var (
	scanner = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

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
		panic("queue: PopFront() called on empty queue")
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
		panic("queue: PopFront() called on empty queue")
	}
	return (*d)[0]
}

func (d *Deque[T]) Back() T {
	if len(*d) == 0 {
		panic("queue: PopFront() called on empty queue")
	}
	return (*d)[len(*d)-1]
}

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	deque := New[int]()

	n := nextInt()
	for i := 0; i < n; i++ {
		command := nextStr()
		switch command {
		case "push_front":
			x := nextInt()
			deque.PushFront(x)
		case "push_back":
			x := nextInt()
			deque.PushBack(x)
		case "pop_front":
			if deque.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, deque.PopFront())
			}
		case "pop_back":
			if deque.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, deque.PopBack())
			}
		case "size":
			fmt.Fprintln(w, deque.Len())
		case "empty":
			if deque.Empty() {
				fmt.Fprintln(w, 1)
			} else {
				fmt.Fprintln(w, 0)
			}
		case "front":
			if deque.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, deque.Front())
			}
		case "back":
			if deque.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, deque.Back())
			}
		}
	}
}
