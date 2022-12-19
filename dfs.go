package GoGraph

// DFS runs the depth-first-search algorithm on a Graph.
// After running, it outputs the parent array from the depth-first-search, which is
// structured as parent[b] = a if a is deemed to be b's parent in the depth-first-search.
func (g *Graph) DFS() []int {
	// color white = 0
	// color gray = 1
	// color black = 2
	color := make([]int, g.Size())
	parent := make([]int, g.Size())
	d := make([]int, g.Size())
	f := make([]int, g.Size())
	var time int
	for x := range g.graph {
		color[x] = 0
	}
	time = 0
	for x := range g.graph {
		if color[x] == 0 {
			dfsVisit(g, x, color, parent, d, f, time)
		}
	}
	return parent

}

func dfsVisit(g *Graph, u int, color []int, parent []int, d []int, f []int, time int) {
	time++
	d[u] = time
	color[u] = 1
	for v := range g.graph[u].list {
		if color[v] == 0 {
			parent[v] = u
			dfsVisit(g, v, color, parent, d, f, time)
		}
	}
	color[u] = 2
	time++
	f[u] = time
}
