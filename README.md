# Graph Coloring (map coloring) Constraint Satisfaction Problem (CSP) with Backtracking implemented in Go

Program solves the graph coloring problem with backtracking, MRV & LCV heuristics, including forward checking, AC-3 algorithm as preprocessing (demo), as well as degree heuristic. 

## Installation

Go should be installed beforehand https://golang.org/doc/install and GOPATH should be set.

Then either download .zip file of this package or go get from command line to GOPATH:

```
$ go get github.com/shahaliyev/go-graph-coloring
```

### Usage

If the installation part is complete, it remains to run the command line with the path to the input file:

```
$ go run go-graph-coloring input/4.txt
```

Before running, make sure that the package is located in GOPATH and the input path is specified correctly.

### Package

The main package consists of the following files:

* **main.go** brings all the children together
* **backtrack.go** implements the core function with recursive backtracking 
* **heuristics.go** implements MRV (Minimal Remaining Value) and LCV (Least Constrained Value) heuristics, as well as the degree heuristic (which can be added as a tie-breaker for mrv, lcv)
* **forward.go** implements forward checking and reduces the domains of vertices
* **ac3.go** implements the AC-3 algorithm as preprocessing (for demonstration purposes).
* **helpers.go** includes helper functions to ease computation
* **io.go** takes care of input and output operations.
* **test.go** checks if the solution is correct
* **global.go** stores global declarations