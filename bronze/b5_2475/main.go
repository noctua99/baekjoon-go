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
	scanner.Split(bufio.ScanWords)

	sum := 0
	for i := 0; i < 5; i++ {
		num := nextInt()
		sum += num * num
	}

	fmt.Fprintln(w, sum%10)
}
