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

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func solve(arr []int) int {
	count := 0
	for _, data := range arr {
		if isPrime(data) {
			count++
		}
	}

	return count
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	fmt.Fprintln(w, solve(arr))
}
