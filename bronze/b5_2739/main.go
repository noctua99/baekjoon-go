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

	n := nextInt()
	for i := 1; i <= 9; i++ {
		fmt.Fprintf(w, "%d * %d = %d\n", n, i, n*i)
	}
}
