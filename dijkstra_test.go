package goGraph

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	fmt.Println("entered test")
	var g Graph
	g.graph = make([]adjList, 4)
	g.size = 4
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	AddEdge(&g, 0, 1, 7)
	AddEdge(&g, 0, 2, 1)
	AddEdge(&g, 2, 3, 2)
	AddEdge(&g, 1, 3, 1)
	actual := Dijkstra(&g, 0)
	var expected []int
	expected = append(expected, 0, 7, 1, 3)
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("wrong")
		}
	}

}
