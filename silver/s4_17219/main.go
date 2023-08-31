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

	urlToPassword := make(map[string]string)

	n, m := nextInt(), nextInt()
	for i := 0; i < n; i++ {
		url, password := nextStr(), nextStr()
		urlToPassword[url] = password
	}
	for i := 0; i < m; i++ {
		url := nextStr()
		fmt.Fprintln(w, urlToPassword[url])
	}
}
