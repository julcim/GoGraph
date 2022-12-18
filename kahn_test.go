package goGraph

import (
	"fmt"
	"testing"
)

func TestKahn(t *testing.T) {
	var g Graph
	g.graph = make([]adjList, 3)
	g.size = 3
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	AddEdge(&g, 0, 1, 2)
	AddEdge(&g, 1, 2, 2)
	actual := Kahn(&g)
	var expected []int
	expected = append(expected, 0, 1, 2)
	for i := 0; i < len(expected); i++ {
		fmt.Println(actual[i])
		if actual[i] != expected[i] {
			t.Errorf("wrong")
		}
	}

}
