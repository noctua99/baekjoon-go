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

	modToCount := make(map[int]bool)
	for i := 0; i < 10; i++ {
		mod := nextInt() % 42
		modToCount[mod] = true
	}

	fmt.Fprintln(w, len(modToCount))
}
