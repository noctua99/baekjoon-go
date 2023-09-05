package main

import (
	"bufio"
	"os"
	"strings"
)

// fastIO
var (
	scanner = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer w.Flush()
	scanner.Buffer([]byte{}, 1000001)

	alp := make([]int, 26)

	s := strings.ToLower(nextStr())
	for _, c := range s {
		alp[c-'a']++
	}

	maxAlp, maxCnt, isTie := ' ', -1, true
	for i := range alp {
		if alp[i] > maxCnt {
			maxAlp = int32('A' + i)
			maxCnt = alp[i]
			isTie = false
		} else if alp[i] == maxCnt {
			isTie = true
		}
	}

	if isTie {
		w.WriteByte('?')
	} else {
		w.WriteString(string(maxAlp))
	}
}
