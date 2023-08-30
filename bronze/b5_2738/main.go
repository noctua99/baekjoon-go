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
	scanner.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	matA := make([][]int, n)
	matB := make([][]int, n)

	for i := 0; i < n; i++ {
		matA[i] = make([]int, m)
		matB[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matA[i][j] = nextInt()
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matB[i][j] = nextInt()
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprintf(w, "%d ", matA[i][j]+matB[i][j])
		}
		fmt.Fprintln(w)
	}
}
