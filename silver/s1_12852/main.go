package main

import (
	"bufio"
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

func main() {
	defer w.Flush()

	n := nextInt()

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

	w.WriteString(strconv.Itoa(dp[n]))
	w.WriteByte('\n')

	for n > 0 {
		w.WriteString(strconv.Itoa(n))
		w.WriteByte(' ')
		if n%3 == 0 && dp[n/3] == dp[n]-1 {
			n /= 3
		} else if n%2 == 0 && dp[n/2] == dp[n]-1 {
			n /= 2
		} else {
			n -= 1
		}
	}
}
