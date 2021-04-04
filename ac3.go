package main

import (
	"container/list"
)

// Defines Arc type
type Arc struct {
	from int
	to int
}

// Revises according to constraints
func revise(X, Y int) bool {
	revised := false

	for _, colorX := range domains[X] {
		flag := false
		for _, colorY :=  range domains[Y] {
			if colorX != colorY {
				flag = true
			}
		}
		// if no value satisfied the constraint
		if flag == false {
			// removing color from the domain of x
			domains[X] = remove(domains[X], colorX)
			revised = true
		}
	}

	return revised
}

// AC-3 algorithm
func ac3(queue *list.List) bool {
	for queue.Len() > 0 {
		// getting arc from the queue
		front := queue.Front()
		queue.Remove(front)
		arc := front.Value.(Arc)
		// if revision took place
		if revise(arc.from, arc.to) {
			if len(domains[arc.from]) == 0 { return false }
			// adding arcs from neighbors to the queue
			for _, neighbor := range graph[arc.from] {
				if colors[neighbor] == 0 && neighbor != arc.to {
					queue.PushBack(Arc{neighbor, arc.from})
				}
			}
		}
	}

	return true
}

// Gets neighboring arcs as queue for ac3
func getArcs(vertex int) *list.List {
	var queue = list.New()

	for _, neighbor := range graph[vertex] {
		// if unassigned
		if colors[neighbor] == 0 {
			queue.PushBack(Arc{neighbor, vertex})
		}
	}

	return queue
}