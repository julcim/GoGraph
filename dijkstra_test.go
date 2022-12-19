package GoGraph

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	var g Graph
	g.graph = make([]adjList, 4)
	g.size = 4
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	g.AddEdge(0, 1, 7)
	g.AddEdge(0, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(1, 3, 1)
	actual := g.Dijkstra(0)
	var expected []int
	expected = append(expected, 0, 7, 1, 3)
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("wrong")
		}
	}

}
