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

func solve(coins []int, k int) int {
	count := 0
	for i := len(coins) - 1; i >= 0; i-- {
		if k >= coins[i] {
			count += k / coins[i]
			k %= coins[i]
		}
	}
	return count
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	coins := make([]int, n)

	for i := 0; i < n; i++ {
		coins[i] = nextInt()
	}

	fmt.Fprintln(w, solve(coins, k))
}
