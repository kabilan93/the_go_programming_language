package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// read entire input from files into memory i one big gulp and split it into lies all at once
	counts := make(map[string]int)
	files := []string{"./input.txt", "./input2.txt"} //os.Args[1:]

	for _, filename := range files {
		data, err := ioutil.ReadFile(filename) // reads a the file content till EOF and returs a byte slice
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err) // takes in var and affs it to the writer to print
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
