package main

import "fmt"

func test() {
	// not testing if there was no solution
	if len(colors) == 0 { return }

	solution := "Correct"
	for vertex := range graph {
		for _, neighbor := range graph[vertex] {
			if colors[vertex] == colors[neighbor] {
				solution = "Incorrect"
			}
		}
	}
	fmt.Println(solution)
}
