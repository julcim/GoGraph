package goGraph

import (
	"testing"
)

func TestDfs(t *testing.T) {
	var g Graph
	g.graph = make([]adjList, 3)
	g.size = 3
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	AddEdge(&g, 0, 1, 2)
	AddEdge(&g, 0, 2, 2)
	actual := DFS(&g)
	var expected []int
	expected = append(expected, 0, 0, 0)
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("wrong")
		}
	}

}

func TestDfs1(t *testing.T) {
	var g Graph
	g.graph = make([]adjList, 4)
	g.size = 4
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	AddEdge(&g, 0, 1, 2)
	AddEdge(&g, 0, 2, 2)
	AddEdge(&g, 2, 3, 2)
	actual := DFS(&g)
	var expected []int
	expected = append(expected, 0, 0, 0, 2)
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("wrong")
		}
	}

}
