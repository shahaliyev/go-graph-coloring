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
	// returning when complete
	if len(colors) == len(graph) { return true }

	// getting the next vertex according to heuristic
	vertex := mrvHeuristic()

	// getting a sorted domain with least constrained values
	domains[vertex] = lcvHeuristic(vertex)

	// iterating over the domain of a vertex
	for _, color := range domains[vertex] {
		// tracking the number of domains visited
		domainCount++

		// recursive search if constraints are met
		if isSafe(vertex, color) {
			// getting arcs queue for ac3
			arcs := getArcs(vertex)

			// checking consistency with AC3
			if ac3(arcs) {
				// coloring vertex
				colors[vertex] = color
				// updating domains
				updateDomains(vertex, color)
				// the next recursion
				if colorGraph() { return true }
			}

			backtrack(vertex, color)
		}
	}
	delete(used, vertex)

	return false
}

// Updates domains for ac3
func updateDomains(vertex, color int) {
	domains[vertex] = remove(domains[vertex], color)

	for _, neighbor := range graph[vertex] {
		if colors[neighbor] == 0 {
			domains[neighbor] = remove(domains[neighbor], colors[vertex])
		}
	}
}

// Backtracks and restores values
func backtrack(vertex, color int) {
	for _, neighbor := range graph[vertex] {
		if !contains(domains[neighbor], colors[vertex]) && colors[vertex] != 0 {
			domains[neighbor] = append(domains[neighbor], colors[vertex])
		}
	}

	domains[vertex] = append(domains[vertex], color)
	delete(colors, vertex)
}
