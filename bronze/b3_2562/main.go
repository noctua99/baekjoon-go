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

	maximum, maxIndex := math.MinInt, -1

	for i := 1; i <= 9; i++ {
		num := nextInt()
		if num > maximum {
			maximum = num
			maxIndex = i
		}
	}

	fmt.Fprintln(w, maximum)
	fmt.Fprintln(w, maxIndex)
}
