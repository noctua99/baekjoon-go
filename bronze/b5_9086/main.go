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

	t := nextInt()
	for i := 0; i < t; i++ {
		scanner.Scan()
		str := scanner.Text()
		fmt.Fprintf(w, "%s%s\n", string(str[0]), string(str[len(str)-1]))
	}
}
