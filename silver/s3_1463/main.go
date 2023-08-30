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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func solve(n int) int {
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i%3 == 0 {
			dp[i] = min(dp[i], dp[i/3]+1)
		}
		if i%2 == 0 {
			dp[i] = min(dp[i], dp[i/2]+1)
		}
	}

	return dp[n]
}

func main() {
	defer w.Flush()

	n := nextInt()

	fmt.Fprintln(w, solve(n))
}
