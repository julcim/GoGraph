package goGraph

import (
	"testing"
)

func TestHasEdge(t *testing.T) {
	var g Graph
	g.graph = make([]adjList, 3)
	g.size = 3
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	g.AddEdge(0, 1, 2)
	g.AddEdge(1, 2, 2)
	actual := g.HasEdge(0, 1)
	expected := true
	if actual != expected {
		t.Errorf("wrong")
	}

}

func TestHasEdge1(t *testing.T) {
	var g Graph
	g.graph = make([]adjList, 3)
	g.size = 3
	for i := 0; i < g.size; i++ {
		g.graph[i].list = map[int]int{}
	}
	g.AddEdge(0, 1, 2)
	g.AddEdge(1, 2, 2)
	actual := g.HasEdge(0, 2)
	expected := false
	if actual != expected {
		t.Errorf("wrong")
	}

}
