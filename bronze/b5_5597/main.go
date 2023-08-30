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

	submitted := make([]bool, 31)
	for i := 0; i < 28; i++ {
		n := nextInt()
		submitted[n] = true
	}

	for i := 1; i <= 30; i++ {
		if !submitted[i] {
			fmt.Fprintln(w, i)
		}
	}
}
