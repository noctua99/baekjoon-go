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

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer w.Flush()

	alp := make([]int, 26)
	for i := range alp {
		alp[i] = -1
	}

	s := nextStr()
	for i, c := range s {
		if alp[c-'a'] == -1 {
			alp[c-'a'] = i
		}
	}

	for i := range alp {
		w.WriteString(strconv.Itoa(alp[i]))
		w.WriteByte(' ')
	}
}
