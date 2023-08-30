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
)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func dfs(v int) {
	visited[v] = true
	fmt.Fprintf(w, "%d ", v)

	for _, node := range graph[v] {
		if !visited[node] {
			dfs(node)
		}
	}
}

func bfs(v int) {
	var queue []int
	queue = append(queue, v)
	visited[v] = true
	fmt.Fprintf(w, "%d ", v)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, node := range graph[cur] {
			if !visited[node] {
				queue = append(queue, node)
				visited[node] = true
				fmt.Fprintf(w, "%d ", node)
			}
		}
	}
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n, m, v := nextInt(), nextInt(), nextInt()
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

	dfs(v)

	for i := 1; i <= n; i++ {
		visited[i] = false
	}
	fmt.Fprintln(w)

	bfs(v)
}
