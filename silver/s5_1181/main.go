package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func nextString() string {
	scanner.Scan()
	return scanner.Text()
}

func solve(words []string) {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) < len(words[j])
	})

	fmt.Fprintln(w, words[0])
	for i := 1; i < len(words); i++ {
		if words[i] != words[i-1] {
			fmt.Fprintln(w, words[i])
		}
	}
}

func main() {
	defer w.Flush()

	n := nextInt()
	words := make([]string, n)

	for i := 0; i < n; i++ {
		words[i] = nextString()
	}

	solve(words)
}
