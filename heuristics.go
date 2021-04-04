package main

// Returns most constrained variable
func mrvHeuristic() int {
	min := INF
	minVertex := INF

	for vertex := range graph {
		// getting the domain length of a vertex
		domain := len(domains[vertex])
		// if domain is minimum and vertex is not explored
		if domain < min && used[vertex] == 0 {
			min = domain
			minVertex = vertex
		}
	}
	// setting used variable
	used[minVertex] = 1

	return minVertex
}

// Returns the sorted domain of least constraining values
func lcvHeuristic(vertex int) []int {
	var domain = sliceToMap(domains[vertex])

	for _, vertexColor := range domains[vertex] {
		for _, neighbor := range graph[vertex] {
			for _, neighborColor := range domains[neighbor] {
				// assigning degrees to the colors in the domain
				if vertexColor != neighborColor {
					domain[vertexColor]++
				}
			}
		}
	}

	// sorting and returning the domain according to the highest degree
	sorted := sortMap(domain)

	// returning new domain as result
	var result []int
	for _, lcv := range sorted {
		result = append(result, lcv.Color)
	}

	return result
}

// Returns a vertex with the maximum degree
// Can be implemented within backtracking as a tie-breaker or a separate heuristic
func degreeHeuristic() int {
	max := -INF
	maxVertex := -INF

	for vertex := range graph {
		// getting degree of a vertex
		degree := len(graph[vertex])
		// if the degree is maximum and vertex is not explored
		if degree > max && used[vertex] == 0 {
			max = degree
			maxVertex = vertex
		}
	}
	// setting used variable
	used[maxVertex] = 1

	// returning vertex with maximum degree
	return maxVertex
}