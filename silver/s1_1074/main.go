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

func main() {
	defer w.Flush()

	scanner.Scan()
	input := scanner.Text()

	arr := strings.Split(input, "-")
	ans := 0

	for _, s := range strings.Split(arr[0], "+") {
		n, _ := strconv.Atoi(s)
		ans += n
	}
	for i := 1; i < len(arr); i++ {
		for _, s := range strings.Split(arr[i], "+") {
			n, _ := strconv.Atoi(s)
			ans -= n
		}
	}

	fmt.Fprintln(w, ans)
}
