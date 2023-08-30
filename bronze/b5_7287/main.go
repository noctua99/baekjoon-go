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
	fmt.Fprintln(w, 27)
	fmt.Fprintln(w, "go_daecoolnoc")
	w.Flush()
}
