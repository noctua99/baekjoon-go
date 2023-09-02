package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	p := nextInt()
	for i := 0; i < p; i++ {
		r, s := nextInt(), nextStr()
		for i := range s {
			for j := 0; j < r; j++ {
				w.WriteString(strings.Repeat(s[i:i+1], r))
			}
		}
		w.WriteByte('\n')
	}
}
