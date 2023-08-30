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

	a, b := nextInt(), nextInt()

	if a > b {
		fmt.Fprintln(w, ">")
	} else if a == b {
		fmt.Fprintln(w, "==")
	} else {
		fmt.Fprintln(w, "<")
	}
}
