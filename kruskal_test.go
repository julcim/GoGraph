package GoGraph

import (
	"reflect"
	"testing"
)

func TestKruskal(t *testing.T) {
	var g Graph
	g.graph = make([]adjList, 3)
	g.size = 3
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	g.AddEdge(0, 1, 2)
	g.AddEdge(1, 0, 2)
	g.AddEdge(1, 2, 2)
	g.AddEdge(2, 1, 2)
	actual := g.Kruskal()
	for i := 0; i < g.size; i++ {
		if !reflect.DeepEqual(actual.graph[i].list, g.graph[i].list) {
			t.Errorf("wrong")
		}
	}

}
