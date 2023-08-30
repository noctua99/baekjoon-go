package main

import (
	"bufio"
	"fmt"
	"os"
)

// fastIO
var (
	scanner = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()

	scanner.Scan()
	grade := scanner.Text()

	switch grade {
	case "A+":
		fmt.Fprintln(w, "4.3")
	case "A0":
		fmt.Fprintln(w, "4.0")
	case "A-":
		fmt.Fprintln(w, "3.7")
	case "B+":
		fmt.Fprintln(w, "3.3")
	case "B0":
		fmt.Fprintln(w, "3.0")
	case "B-":
		fmt.Fprintln(w, "2.7")
	case "C+":
		fmt.Fprintln(w, "2.3")
	case "C0":
		fmt.Fprintln(w, "2.0")
	case "C-":
		fmt.Fprintln(w, "1.7")
	case "D+":
		fmt.Fprintln(w, "1.3")
	case "D0":
		fmt.Fprintln(w, "1.0")
	case "D-":
		fmt.Fprintln(w, "0.7")
	case "F":
		fmt.Fprintln(w, "0.0")
	}
}
