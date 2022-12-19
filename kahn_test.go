package GoGraph

import (
	"testing"
)

func TestKahn(t *testing.T) {
	var g Graph
	g.graph = make([]adjList, 3)
	g.size = 3
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	g.AddEdge(0, 1, 2)
	g.AddEdge(1, 2, 2)
	actual := g.Kahn()
	var expected []int
	expected = append(expected, 0, 1, 2)
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("wrong")
		}
	}

}
