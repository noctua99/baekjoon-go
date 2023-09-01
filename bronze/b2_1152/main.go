package main

import (
	"bufio"
	"bytes"
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
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	input := scanner.Bytes()
	input = bytes.TrimSpace(input)

	if len(input) == 0 {
		fmt.Fprintln(w, 0)
	} else {
		fmt.Fprintln(w, bytes.Count(input, []byte{byte(' ')})+1)
	}
}
