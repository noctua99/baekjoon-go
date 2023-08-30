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

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("queue: PopFront() called on empty queue")
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
	return (*s)[len(*s)-1]
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

	stack := New[int]()

	n := nextInt()
	for i := 0; i < n; i++ {
		command := nextStr()
		switch command {
		case "push":
			x := nextInt()
			stack.Push(x)
		case "pop":
			if stack.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, stack.Pop())
			}
		case "size":
			fmt.Fprintln(w, stack.Len())
		case "empty":
			if stack.Empty() {
				fmt.Fprintln(w, 1)
			} else {
				fmt.Fprintln(w, 0)
			}
		case "top":
			if stack.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, stack.Peek())
			}
		}
	}
}
