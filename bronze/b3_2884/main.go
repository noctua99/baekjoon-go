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

	h, m := nextInt(), nextInt()

	if m >= 45 {
		m -= 45
	} else {
		h -= 1
		if h < 0 {
			h += 24
		}
		m += 15
	}

	w.WriteString(strconv.Itoa(h))
	w.WriteByte(' ')
	w.WriteString(strconv.Itoa(m))
}
