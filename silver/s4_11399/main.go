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
	sort.Ints(arr)

	sum := 0
	for i, val := range arr {
		sum += val * (len(arr) - i)
	}

	return sum
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	fmt.Fprintln(w, solve(arr))
}
