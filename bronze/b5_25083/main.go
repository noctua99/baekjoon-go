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
	fmt.Fprintln(w, "         ,r'\"7\nr`-_   ,'  ,/\n \\. \". L_r'\n   `~\\/\n      |\n      |")
	w.Flush()
}
