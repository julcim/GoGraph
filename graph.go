package goGraph

type Graph struct {
	size  int
	graph []adjList
}

type adjList struct {
	list map[int]int
}

// getters
func GetSize(g *Graph) (size int) {
	return g.size
}

func GetAdjList(g *Graph) (adjList []adjList) {
	return g.graph
}

// returns whether the current graph has a directed edge between
// 2 vertices
func HasEdge(g *Graph, u int, v int) bool {
	if u < 0 || u >= GetSize(g) || v < 0 || v >= GetSize(g) {
		panic("invalid vertices")
	}
	a := g.graph[u]
	_, isPresent := a.list[v]
	return isPresent

}

// gets the weight of a directed edge
func GetWeight(g *Graph, u int, v int) int {
	if u < 0 || u >= GetSize(g) || v < 0 || v >= GetSize(g) {
		panic("invalid vertices")
	}
	if !(HasEdge(g, u, v)) {
		panic("there is no edge from u to v")
	}
	a := g.graph[u]
	value := a.list[v]
	return value
}

// creates an edge from vertex u to v
func AddEdge(g *Graph, u int, v int, newWeight int) {
	if u < 0 || u >= GetSize(g) || v < 0 || v >= GetSize(g) {
		panic("invalid vertices")
	}
	g.graph[u].list[v] = newWeight
}

// returns all the vertices outgoing from a given vertex
func Neighbors(g *Graph, u int) []int {
	a := g.graph[u]
	keys := make([]int, 0, len(a.list))
	for k := range a.list {
		keys = append(keys, k)
	}
	return keys
}
