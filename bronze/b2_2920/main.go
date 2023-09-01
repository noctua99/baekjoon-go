package main

import (
	"bufio"
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

func isAscending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func isDescending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	arr := make([]int, 8)
	for i := 0; i < 8; i++ {
		arr[i] = nextInt()
	}

	if isAscending(arr) {
		w.WriteString("ascending\n")
	} else if isDescending(arr) {
		w.WriteString("descending\n")
	} else {
		w.WriteString("mixed\n")
	}
}
