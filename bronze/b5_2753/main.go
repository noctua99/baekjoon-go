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

	year := nextInt()

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Fprintln(w, 1)
	} else {
		fmt.Fprintln(w, 0)
	}
}
