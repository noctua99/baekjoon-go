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

func solve(arr []int) {
	sort.Ints(arr)
	for i := range arr {
		fmt.Fprintln(w, arr[i])
	}
}

func main() {
	defer w.Flush()

	n := nextInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	solve(arr)
}
