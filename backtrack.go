package main

// Checks if constraints are met
func isSafe(from int, color int) bool {
	for _, neighbor := range graph[from] {
		if color == colors[neighbor] { return false }
	}

	return true
}

// Backtracks and colors the graph
func colorGraph() bool {
	var color int

	// returning when complete
	if len(colors) == len(graph) { return true }

	// getting the next vertex according to heuristic
	vertex := mrvHeuristic()

	// getting a sorted domain with least constrained values
	domain := lcvHeuristic(vertex)

	// iterating over the domain of a vertex
	for _, lcv := range domain {
		// getting color value
		color = lcv.Color
		// tracking the number of domains visited
		domainCount++
		// recursive search if constraints are met
		if isSafe(vertex, color) {
			colors[vertex] = color
			// interweaving ac-3 with forward check
			forwardCheck(vertex)
			if colorGraph() {
				return true
			}
			// backtracking
			forwardCheckBacktrack(vertex)
			delete(colors, vertex)
		}
	}
	// backtracking
	delete(used, vertex)

	return false
}
