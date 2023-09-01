package main

import (
	"bufio"
	"math"
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

func getMean(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return int(math.Round(float64(sum) / float64(len(arr))))
}

func getMedian(arr []int) int {
	return arr[len(arr)/2]
}

func getMode(arr []int) int {
	numToCount := make(map[int]int, len(arr))
	for i := range arr {
		numToCount[arr[i]]++
	}

	maxCount := -1
	for _, count := range numToCount {
		if count > maxCount {
			maxCount = count
		}
	}

	var modes []int
	for num, count := range numToCount {
		if count == maxCount {
			modes = append(modes, num)
		}
	}

	if len(modes) > 1 {
		sort.Ints(modes)
		return modes[1]
	}

	return modes[0]
}

func getRange(arr []int) int {
	return arr[len(arr)-1] - arr[0]
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	sort.Ints(arr)

	// mean: average of all the numbers
	w.WriteString(strconv.Itoa(getMean(arr)))
	w.WriteByte('\n')
	// median: middle number, when in order
	w.WriteString(strconv.Itoa(getMedian(arr)))
	w.WriteByte('\n')
	// mode: most common number.
	w.WriteString(strconv.Itoa(getMode(arr)))
	w.WriteByte('\n')
	// range: largest number minus the smallest number
	w.WriteString(strconv.Itoa(getRange(arr)))
	w.WriteByte('\n')
}
