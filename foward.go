package main

// Forward checks and reduces the domains to search
func forwardCheck(vertex int) {
	for _, neighbor := range graph[vertex] {
		// if unassigned
		if colors[neighbor] == 0 {
			// removes the color from the domain of a neighbor
			domains[neighbor] = remove(domains[neighbor], colors[vertex])
		}
	}
}

// Restores the domains in case of backtracking
func forwardCheckBacktrack(vertex int) {
	for _, neighbor := range graph[vertex] {
		if !contains(domains[neighbor], colors[vertex]){
			domains[neighbor] = append(domains[neighbor], colors[vertex])
		}
	}
}
