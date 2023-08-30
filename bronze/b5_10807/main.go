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
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	myMap := make(map[int]int)
	for i := 0; i < n; i++ {
		num := nextInt()
		myMap[num]++
	}

	fmt.Fprintln(w, myMap[nextInt()])
}
