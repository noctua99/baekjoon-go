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

func factorial(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}

	return ans
}

func main() {
	defer w.Flush()

	n := nextInt()

	fmt.Fprintln(w, factorial(n))
}
