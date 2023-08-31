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

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	t := nextInt()
	for i := 0; i < t; i++ {
		categoryToNum := make(map[string]int)
		n := nextInt()
		for j := 0; j < n; j++ {
			nextStr()
			categoryToNum[nextStr()]++
		}

		ans := 1
		for _, num := range categoryToNum {
			ans *= num + 1
		}
		ans -= 1

		fmt.Fprintln(w, ans)
	}
}
