package main

import (
	"bufio"
	"fmt"
	"math"
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

func nextFloat() float64 {
	scanner.Scan()
	n, _ := strconv.ParseFloat(scanner.Text(), 64)
	return n
}

func solve(scores []float64) float64 {
	var maxScore, sum float64
	for _, score := range scores {
		maxScore = math.Max(maxScore, score)
		sum += score
	}
	return sum / float64(len(scores)) / maxScore * 100.0
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	scores := make([]float64, n)

	for i := 0; i < n; i++ {
		scores[i] = nextFloat()
	}

	fmt.Fprintln(w, solve(scores))
}
