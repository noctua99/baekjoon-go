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

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	set := make([]bool, 21)

	m := nextInt()
	for i := 0; i < m; i++ {
		cmd := nextStr()
		switch cmd {
		case "add":
			set[nextInt()] = true
		case "remove":
			set[nextInt()] = false
		case "check":
			if set[nextInt()] {
				fmt.Fprintln(w, 1)
			} else {
				fmt.Fprintln(w, 0)
			}
		case "toggle":
			x := nextInt()
			set[x] = !set[x]
		case "all":
			for i := 1; i <= 20; i++ {
				set[i] = true
			}
		case "empty":
			set = make([]bool, 21)
		}
	}
}
