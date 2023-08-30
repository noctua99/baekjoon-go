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

type member struct {
	order int
	age   int
	name  string
}

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func solve(members []member) {
	sort.Slice(members, func(i, j int) bool {
		if members[i].age == members[j].age {
			return members[i].order < members[j].order
		}
		return members[i].age < members[j].age
	})

	for _, member := range members {
		fmt.Fprintln(w, member.age, member.name)
	}
}

func main() {
	defer w.Flush()
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	members := make([]member, n)

	for i := 0; i < n; i++ {
		members[i] = member{
			order: i,
			age:   nextInt(),
			name:  nextStr(),
		}
	}

	solve(members)
}
