package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Builds graphs by scanning each line of the input file
func buildFromFile(input *os.File) {
	var line string
	var section = "color count"

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		// scanning next line
		line = scanner.Text()
		// skipping comments
		if strings.HasPrefix(line, "#") { continue }
		// getting values
		if section == "graph" {
			// getting vertices from tokens
			tokens := strings.Split(line, ",")
			from := stringToInt(tokens[0])
			to := stringToInt(tokens[1])
			// building graph & queue if not duplicate
			if !contains(graph[from], to) {
				graph[from] = append(graph[from], to)
				graph[to] = append(graph[to], from)
				queue.PushBack(Arc{from, to})
				queue.PushBack(Arc{to, from})
			}
		} else
		if section == "color count" {
			// getting colorCount value
			tokens := strings.Split(line, " ")
			colorCount = stringToInt(tokens[2])
			// changing section
			section = "graph"
		}
	}
	// error checking
	err := scanner.Err(); check(err)
}

func output(display bool) {
	if display {
		fmt.Println("Domains visited:", domainCount)
		fmt.Println("Colored graph:", colors)
	} else {
		fmt.Println("No solution")
	}
}
