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
	fmt.Fprintln(w, "|\\_/|\n|q p|   /}\n( 0 )\"\"\"\\\n|\"^\"`    |\n||_/=\\\\__|")
	w.Flush()
}
