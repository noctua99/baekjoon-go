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
	sum := 0
	xorSum := 0
	for i := 0; i < n; i++ {
		command := nextInt()
		switch command {
		case 1: // Add x to the tail of A.
			x := nextInt()
			sum += x
			xorSum ^= x
		case 2: // Remove x from A. If there is more than one x in A, only the earliest one is removed. Only queries with x in A are always given.
			x := nextInt()
			sum -= x
			xorSum ^= x
		case 3: // Outputs the sum of all elements in A.
			fmt.Fprintln(w, sum)
		case 4: // Outputs the value of the XOR of all elements in A.
			fmt.Fprintln(w, xorSum)
		}
	}
}
