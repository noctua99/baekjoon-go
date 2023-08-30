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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()

	fmt.Fprintln(w, abs(n-m))
}
