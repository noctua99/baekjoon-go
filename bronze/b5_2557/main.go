package main

import (
	"bufio"
	"fmt"
	"os"
)

// fastIO
var (
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	fmt.Fprintln(w, "Hello World!")
	w.Flush()
}
