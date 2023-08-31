package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	notHeard := make(map[string]bool)
	var notHeardAndSeen []string

	n, m := nextInt(), nextInt()
	for i := 0; i < n; i++ {
		notHeard[nextStr()] = true
	}
	for i := 0; i < m; i++ {
		name := nextStr()
		if notHeard[name] {
			notHeardAndSeen = append(notHeardAndSeen, name)
		}
	}

	sort.Strings(notHeardAndSeen)

	fmt.Fprintln(w, len(notHeardAndSeen))
	for _, name := range notHeardAndSeen {
		fmt.Fprintln(w, name)
	}
}
