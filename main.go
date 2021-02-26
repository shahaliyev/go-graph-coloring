package main

import (
	"os"
)

func main() {
	// reading input
	input := readInput(os.Args[1]); defer input.Close()

	// building graph and setting input values
	buildFromFile(input)

	// building domains
	buildDomains()

	// preprocessing with AC-3 algorithm (does nothing)
	ac3()

	// backtracking and coloring graph
	display := colorGraph()

	// displaying output
	output(display)

	// testing
	test()
}
