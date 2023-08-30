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

var isPrime []bool

func setIsPrime(n int) {
	isPrime = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
}

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

// According to the Strong Goldbach Conjecture, every even natural number greater than 2 is the sum of two primes.
// 1. Any even number not less than 8 can be decomposed into the sum of four primes.
// 2. Any odd numbers not less than 9 can be decomposed into the sum of four primes.
// Therefore, any number not less than 8 can be decomposed into the sum of four primes.
func solve(n int) {
	if n < 8 {
		fmt.Fprintln(w, -1)
		return
	}

	var ans []int

	// When n is even number
	if n%2 == 0 {
		ans = append(ans, 2, 2)
		n -= 4

		for i := 2; i < n; i++ {
			if isPrime[i] && isPrime[n-i] {
				ans = append(ans, i, n-i)
				break
			}
		}
	} else { // When n is odd number
		ans = append(ans, 2)
		n -= 2

		for i := 4; i <= n; i += 2 {
			if isPrime[n-i] {
				ans = append(ans, n-i)
				n = i
				break
			}
		}

		for i := 2; i < n; i++ {
			if isPrime[i] && isPrime[n-i] {
				ans = append(ans, i, n-i)
				break
			}
		}
	}

	for i := range ans {
		fmt.Fprintf(w, "%d ", ans[i])
	}
}

func main() {
	defer w.Flush()

	n := nextInt()

	setIsPrime(n)
	solve(n)
}
