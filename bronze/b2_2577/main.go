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

	a, b, c := nextInt(), nextInt(), nextInt()
	res := a * b * c

	count := make([]int, 10)
	for res > 0 {
		count[res%10]++
		res /= 10
	}

	for i := 0; i < 10; i++ {
		fmt.Fprintln(w, count[i])
	}
}
