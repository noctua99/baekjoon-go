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

func main() {
	defer w.Flush()

	dp := make([]int, 11)
	dp[1], dp[2], dp[3] = 1, 2, 4
	for i := 4; i <= 10; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	t := nextInt()
	for i := 0; i < t; i++ {
		fmt.Fprintln(w, dp[nextInt()])
	}
}
