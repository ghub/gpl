// Exercise 1.4: Modify dup2 to print the names of all files in which each
// duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, count map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := filename + "\t" + input.Text()
		count[key]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
