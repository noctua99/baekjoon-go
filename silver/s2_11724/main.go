package main

import (
	"bufio"
	"fmt"
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
	graph   [][]int
	visited []bool
	count   int
)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func dfs(v int) {
	visited[v] = true

	for _, node := range graph[v] {
		if !visited[node] {
			dfs(node)
		}
	}
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	graph = make([][]int, n+1)
	visited = make([]bool, n+1)

	for i := 1; i <= n; i++ {
		graph[i] = []int{}
	}

	for i := 0; i < m; i++ {
		start := nextInt()
		end := nextInt()
		graph[start] = append(graph[start], end)
		graph[end] = append(graph[end], start)
	}

	for i := 1; i <= n; i++ {
		sort.Ints(graph[i])
	}

	for i := 1; i <= n; i++ {
		if !visited[i] {
			count++
			dfs(i)
		}
	}

	fmt.Fprintln(w, count)
}
