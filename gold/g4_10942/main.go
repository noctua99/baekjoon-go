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

func isPalindrome(arr []int, start, end int) bool {
	for start < end {
		if arr[start] != arr[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		arr[i] = nextInt()
	}

	m := nextInt()
	for i := 0; i < m; i++ {
		start := nextInt()
		end := nextInt()
		if isPalindrome(arr, start, end) {
			fmt.Fprintln(w, 1)
		} else {
			fmt.Fprintln(w, 0)
		}
	}
}
