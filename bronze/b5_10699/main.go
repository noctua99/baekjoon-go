package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// fastIO
var (
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	fmt.Fprintln(w, time.Now().Format("2006-01-02"))
	w.Flush()
}
