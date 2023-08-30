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

func (q *Queue[T]) Front() T {
	return (*q)[0]
}

func (q *Queue[T]) Back() T {
	return (*q)[len(*q)-1]
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

	queue := New[int]()

	n := nextInt()
	for i := 0; i < n; i++ {
		command := nextStr()
		switch command {
		case "push":
			x := nextInt()
			queue.Push(x)
		case "pop":
			if queue.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, queue.Pop())
			}
		case "size":
			fmt.Fprintln(w, queue.Len())
		case "empty":
			if queue.Empty() {
				fmt.Fprintln(w, 1)
			} else {
				fmt.Fprintln(w, 0)
			}
		case "front":
			if queue.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, queue.Front())
			}
		case "back":
			if queue.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, queue.Back())
			}
		}
	}
}
