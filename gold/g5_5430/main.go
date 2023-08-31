package main

import (
	"bufio"
	"fmt"
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

func solve() {
	p := nextStr()
	_ = nextInt()
	arrStr := nextStr()
	var arr []string

	arrStr = strings.Trim(arrStr, "[]")
	if arrStr != "" {
		arr = strings.Split(arrStr, ",")
	}

	var reversed bool
	for _, c := range p {
		if c == 'R' {
			reversed = !reversed
		} else {
			if len(arr) == 0 {
				fmt.Fprintln(w, "error")
				return
			}
			if reversed {
				arr = arr[:len(arr)-1]
			} else {
				arr = arr[1:]
			}
		}
	}
	if reversed {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	fmt.Fprintf(w, "[%s]\n", strings.Join(arr, ","))
}

func main() {
	defer w.Flush()
	scanner.Buffer(make([]byte, 0), 400000)

	t := nextInt()
	for i := 0; i < t; i++ {
		solve()
	}
}
