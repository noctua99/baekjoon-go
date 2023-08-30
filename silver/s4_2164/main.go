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

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func solve(arr *[]int) int {
	for len(*arr) > 1 {
		*arr = (*arr)[1:]
		*arr = append((*arr)[1:], (*arr)[0])
	}
	return (*arr)[0]
}

func main() {
	defer w.Flush()

	n := nextInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	fmt.Fprintln(w, solve(&arr))
}
