package goGraph

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
	AddEdge(&g, 0, 1, 2)
	AddEdge(&g, 1, 0, 2)
	AddEdge(&g, 1, 2, 2)
	AddEdge(&g, 2, 1, 2)
	actual := Kruskal(&g)
	for i := 0; i < g.size; i++ {
		if !reflect.DeepEqual(actual.graph[i].list, g.graph[i].list) {
			t.Errorf("wrong")
		}
	}

}
