package goGraph

var time int

// color white = 0
// color gray = 1
// color black = 2

func DFS(g *Graph) []int {
	color := make([]int, GetSize(g))
	parent := make([]int, GetSize(g))
	d := make([]int, GetSize(g))
	f := make([]int, GetSize(g))
	for x := range g.graph {
		color[x] = 0
	}
	time = 0
	for x := range g.graph {
		if color[x] == 0 {
			DFSVisit(g, x, color, parent, d, f)
		}
	}
	return parent

}

func DFSVisit(g *Graph, u int, color []int, parent []int, d []int, f []int) {
	time++
	d[u] = time
	color[u] = 1
	for v := range g.graph[u].list {
		if color[v] == 0 {
			parent[v] = u
			DFSVisit(g, v, color, parent, d, f)
		}
	}
	color[u] = 2
	time++
	f[u] = time
}
