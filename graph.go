// Package GoGraph implements graph algorithms and a Graph class to build graphs to
// run those algorithms on.
package GoGraph

import (
	"errors"
)

type error interface {
	Error() string
}

// Graph represents a graph data structure and has parameters of size of
// the graph (number of vertices) and a slice adjList which corresponds to the
// adjacency list of the graph. This is implemented such that the index of the
// adjList corresponds to the vertex of the Graph.
type Graph struct {
	size  int
	graph []adjList
}

// AdjList represents the adjacency list for each vertex in the Graph above. It maps a vertex
// v to its edge weight. If making an undirected graph, make the edge weight 0.
type adjList struct {
	list map[int]int
}

// InitializeGraph initializes the input Graph by filling the adjancey list with empty maps
// to avoid errors when adding edges.
func (g *Graph) InitializeGraph() {
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
}

// Size returns the size parameter of the Graph
func (g *Graph) Size() (size int) {
	return g.size
}

// AdjList returns the adjList parameter of the Graph
func (g *Graph) AdjList() (adjList []adjList) {
	return g.graph
}

// HasEdge is called on a Graph and two vertices. It outputs true if there is a
// directed edge between the two vertices and false otherwise. HasEdge throws an
// error if vertices outside the bounds are inputted.
func (g *Graph) HasEdge(u int, v int) (bool, error) {
	if u < 0 || u >= g.size || v < 0 || v >= g.size {
		return false, errors.New("invalid vertices")
	}
	a := g.graph[u]
	_, isPresent := a.list[v]
	return isPresent, nil

}

// GetWeight is called on a Graph and two vertices u and v. It returns the weight
// of a directed edge from u to v. If the edge u to v does not exist, or the
// vertices are out of range of the size of the graph, GetWeight throws an error.
func (g *Graph) GetWeight(u int, v int) (int, error) {
	if u < 0 || u >= g.size || v < 0 || v >= g.size {
		return -1, errors.New("invalid vertices")
	}
	has, err := g.HasEdge(u, v)
	if err == nil && !has {
		return -1, errors.New("there is no edge from u to v")
	}
	a := g.graph[u]
	value := a.list[v]
	return value, nil
}

// AddEdge creates a vertex from u to v with a weight newWeight in a Graph.
// If the vertices are not in the bounds of the Graph, AddEdge throws an error.
func (g *Graph) AddEdge(u int, v int, newWeight int) error {
	if u < 0 || u >= g.size || v < 0 || v >= g.size {
		return errors.New("invalid vertices")
	}
	g.graph[u].list[v] = newWeight
	return nil
}

// Neighbors returns all the outgoing vertices/edges from a given vertex u
// in a given Graph.
func (g *Graph) Neighbors(u int) []int {
	a := g.graph[u]
	keys := make([]int, 0, len(a.list))
	for k := range a.list {
		keys = append(keys, k)
	}
	return keys
}
