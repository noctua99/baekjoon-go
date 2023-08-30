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
	input := scanner.Text()
	for _, c := range input {
		if c >= 'a' {
			fmt.Fprint(w, string(c-32))
		} else {
			fmt.Fprint(w, string(c+32))
		}
	}
}
