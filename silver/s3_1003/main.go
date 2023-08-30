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

	fibo := make([]int, 41)
	fibo[1] = 1
	for i := 2; i <= 40; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
	}

	t := nextInt()
	for i := 0; i < t; i++ {
		n := nextInt()
		if n == 0 {
			fmt.Fprintln(w, 1, 0)
		} else {
			fmt.Fprintln(w, fibo[n-1], fibo[n])
		}
	}
}
