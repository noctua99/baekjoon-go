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

	score := nextInt()

	if score >= 90 {
		fmt.Fprintln(w, "A")
	} else if score >= 80 {
		fmt.Fprintln(w, "B")
	} else if score >= 70 {
		fmt.Fprintln(w, "C")
	} else if score >= 60 {
		fmt.Fprintln(w, "D")
	} else {
		fmt.Fprintln(w, "F")
	}
}
