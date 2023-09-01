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

	n := nextInt()
	input := nextStr()

	sum := 0
	for i := 0; i < n; i++ {
		sum += int(input[i] - '0')
	}

	fmt.Fprintln(w, sum)
}
