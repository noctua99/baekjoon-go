package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

// fastIO
var (
	scanner = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

var (
	n, m    int
	visited []bool
	arr     []int
	ans     []int
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

	for i := 0; i < n; i++ {
		if !visited[i] {
			visited[i] = true
			ans[depth] = arr[i]
			dfs(depth + 1)
			visited[i] = false
		}
	}
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n, m = nextInt(), nextInt()
	visited = make([]bool, n)
	arr = make([]int, n)
	ans = make([]int, m)

	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	sort.Ints(arr)

	dfs(0)
}
