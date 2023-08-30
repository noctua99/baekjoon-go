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

var minimum int

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func dfs(n int) {
	if n > minimum {
		if isPrime(n) {
			fmt.Fprintln(w, n)
			return
		}
	}
	for _, val := range []int{1, 3, 7, 9} {
		if isPrime(n*10 + val) {
			dfs(n*10 + val)
		}
	}
}

func main() {
	defer w.Flush()

	n := nextInt()
	minimum = int(math.Pow10(n - 1))

	dfs(2)
	dfs(3)
	dfs(5)
	dfs(7)
}
