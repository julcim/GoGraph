package goGraph

import (
	"fmt"
	"testing"
)

func TestBfs(t *testing.T) {
	var g Graph
	g.graph = make([]adjList, 3)
	g.size = 3
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	fmt.Println(GetSize(&g))
	AddEdge(&g, 0, 1, 2)
	AddEdge(&g, 1, 2, 2)
	actual := Bfs(&g, 0)
	var expected []int
	expected = append(expected, -1, 0, 1)
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("wrong")
		}
	}

}
