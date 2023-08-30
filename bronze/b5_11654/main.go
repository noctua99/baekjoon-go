package main

import (
	"bufio"
	"fmt"
	"os"
)

// fastIO
var (
	scanner = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()

	scanner.Scan()
	fmt.Fprintln(w, scanner.Bytes()[0])
}
