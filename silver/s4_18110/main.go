package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// fastIO
var (
	scanner = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func solve(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)
	start := int(float64(len(arr))*0.15 + 0.5)
	end := len(arr) - start

	sum := 0
	for i := start; i < end; i++ {
		sum += arr[i]
	}

	return int(float64(sum)/float64(end-start) + 0.5)
}

func main() {
	defer w.Flush()

	n := nextInt()
	scores := make([]int, n)

	for i := 0; i < n; i++ {
		scores[i] = nextInt()
	}

	fmt.Fprintln(w, solve(scores))
}
