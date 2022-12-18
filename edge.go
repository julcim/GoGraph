package goGraph

// An Edge is used by the Kruskal class to correctly use the Union-Find data structure.
// An Edge has a start vertex, end vertex, and weight, which can be set to 0 for an
// unweighted graph.
type Edge struct {
	start  int
	end    int
	weight int
}
