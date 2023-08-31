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

	var dp [101]int
	dp[1], dp[2], dp[3], dp[4], dp[5] = 1, 1, 1, 2, 2
	for i := 6; i <= 100; i++ {
		dp[i] = dp[i-1] + dp[i-5]
	}

	t := nextInt()
	for i := 0; i < t; i++ {
		fmt.Fprintln(w, dp[nextInt()])
	}
}
