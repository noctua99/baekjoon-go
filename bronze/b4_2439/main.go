package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

	for i := 1; i <= n; i++ {
		w.WriteString(strings.Repeat(" ", n-i))
		w.WriteString(strings.Repeat("*", i))
		w.WriteByte('\n')
	}
}
