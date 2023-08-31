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

	nameToNum := make(map[string]int)
	numToName := make(map[int]string)

	n, m := nextInt(), nextInt()
	for i := 1; i <= n; i++ {
		name := nextStr()
		nameToNum[name] = i
		numToName[i] = name
	}
	for i := 0; i < m; i++ {
		input := nextStr()
		if num, err := strconv.Atoi(input); err == nil {
			fmt.Fprintln(w, numToName[num])
		} else {
			fmt.Fprintln(w, nameToNum[input])
		}
	}
}
