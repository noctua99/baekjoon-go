package main

import (
	"bufio"
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

	t := nextInt()
	for i := 0; i < t; i++ {
		h, _, n := nextInt(), nextInt(), nextInt()
		roomNumber := n%h*100 + (n/h + 1)
		if n%h == 0 {
			roomNumber = h*100 + n/h
		}
		w.WriteString(strconv.Itoa(roomNumber))
		w.WriteByte('\n')
	}
}
