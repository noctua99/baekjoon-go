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

	t := nextInt()
	for i := 0; i < t; i++ {
		countO := 0
		sum := 0
		quizRes := nextStr()
		for i := range quizRes {
			if quizRes[i] == 'O' {
				countO++
				sum += countO
			} else {
				countO = 0
			}
		}
		fmt.Fprintln(w, sum)
	}
}
