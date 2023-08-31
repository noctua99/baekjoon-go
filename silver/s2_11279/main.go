package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// fastIO
var (
	scanner = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

type intHeap []int

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *intHeap) Pop() (v any) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	myHeap := &intHeap{}
	heap.Init(myHeap)

	n := nextInt()
	for i := 0; i < n; i++ {
		x := nextInt()
		if x > 0 {
			heap.Push(myHeap, x)
		} else {
			if myHeap.Len() > 0 {
				fmt.Fprintln(w, heap.Pop(myHeap))
			} else {
				fmt.Fprintln(w, 0)
			}
		}
	}
}
