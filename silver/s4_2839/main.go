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

func solve(n int) int {
	count := 0
	for n%5 != 0 {
		n -= 3
		count++
	}
	if n < 0 {
		return -1
	}
	return count + n/5
}

func main() {
	defer w.Flush()

	n := nextInt()
	fmt.Fprintln(w, solve(n))
}
