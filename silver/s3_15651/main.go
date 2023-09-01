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

var (
	n, m int
	ans  []int
)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func dfs(depth int) {
	if depth == m {
		for i := range ans {
			w.WriteString(strconv.Itoa(ans[i]))
			w.WriteByte(' ')
		}
		w.WriteByte('\n')
		return
	}

	for i := 1; i <= n; i++ {
		ans[depth] = i
		dfs(depth + 1)
	}
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n, m = nextInt(), nextInt()
	ans = make([]int, m)

	dfs(0)
}
