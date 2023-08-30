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

	for scanner.Scan() {
		input := scanner.Bytes()
		sum := int(input[0]-'0') + int(input[2]-'0')
		fmt.Fprintln(w, sum)
	}
}
