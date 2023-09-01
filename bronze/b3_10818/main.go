package main

import (
	"bufio"
	"fmt"
	"math"
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

	minimum, maximum := math.MaxInt, math.MinInt

	n := nextInt()
	for i := 0; i < n; i++ {
		num := nextInt()
		if num < minimum {
			minimum = num
		}
		if num > maximum {
			maximum = num
		}
	}

	fmt.Fprintln(w, minimum, maximum)
}
