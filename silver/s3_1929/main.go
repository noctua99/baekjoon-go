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
	isPrime []bool
)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func setIsPrime(n int) {
	isPrime = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
}

func solve(start, end int) {
	for i := start; i <= end; i++ {
		if isPrime[i] {
			fmt.Fprintln(w, i)
		}
	}
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	m, n := nextInt(), nextInt()

	setIsPrime(n)
	solve(m, n)
}
