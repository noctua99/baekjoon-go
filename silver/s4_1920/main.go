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

	n := nextInt()
	myMap := make(map[int]bool)

	for i := 0; i < n; i++ {
		myMap[nextInt()] = true
	}

	m := nextInt()
	for i := 0; i < m; i++ {
		if myMap[nextInt()] {
			fmt.Fprintln(w, 1)
		} else {
			fmt.Fprintln(w, 0)
		}
	}
}
