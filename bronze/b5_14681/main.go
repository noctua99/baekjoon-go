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

	x, y := nextInt(), nextInt()

	if x > 0 {
		if y > 0 {
			fmt.Fprintln(w, 1)
		} else {
			fmt.Fprintln(w, 4)
		}
	} else {
		if y > 0 {
			fmt.Fprintln(w, 2)
		} else {
			fmt.Fprintln(w, 3)
		}
	}
}
