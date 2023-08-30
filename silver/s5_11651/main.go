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

type coordinate struct {
	x int
	y int
}

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func solve(coordinates []coordinate) {
	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i].y == coordinates[j].y {
			return coordinates[i].x < coordinates[j].x
		}
		return coordinates[i].y < coordinates[j].y
	})

	for _, c := range coordinates {
		fmt.Fprintln(w, c.x, c.y)
	}
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	coordinates := make([]coordinate, n)

	for i := 0; i < n; i++ {
		coordinates[i] = coordinate{
			x: nextInt(),
			y: nextInt(),
		}
	}

	solve(coordinates)
}
