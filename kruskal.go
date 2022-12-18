package goGraph

import (
	"sort"
)

// Kruskal runs the Kruskal algorithm on a given Graph and outputs a Graph pointer to a
// Minimum Spanning Tree of that graph.
func Kruskal(g *Graph) *Graph {
	var t Graph
	t.size = g.size
	t.graph = make([]adjList, t.size)
	t.InitializeGraph()
	edges := make([]Edge, 0)
	for i := 0; i < g.size; i++ {
		for u := range g.graph[i].list {
			var e Edge
			e.start = i
			e.end = u
			weight, err := g.GetWeight(i, u)
			if err == nil {
				e.weight = weight
			}
			edges = append(edges, e)
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight > edges[j].weight
	})

	rank := make([]int, g.size)
	parent := make([]int, g.size)
	for i := 0; i < g.Size(); i++ {
		rank[i] = 0
		parent[i] = i
	}
	for i := 0; i < len(edges)-1; i++ {
		if find(edges[i].start, parent) != find(edges[i].end, parent) {
			weight, err := g.GetWeight(edges[i].start, edges[i].end)
			if err == nil {
				t.AddEdge(edges[i].start, edges[i].end, weight)
				t.AddEdge(edges[i].end, edges[i].start, weight)
			}
		}
		union(edges[i].end, edges[i].start, parent, rank)
	}
	return &t

}

func find(v int, arr []int) int {
	if v != arr[v] {
		arr[v] = find(arr[v], arr)
	}
	return arr[v]
}

func union(x int, y int, parent []int, rank []int) {
	rx := find(x, parent)
	ry := find(y, parent)
	if rx == ry {
		return
	}
	if rank[rx] > rank[ry] {
		parent[ry] = rx
	} else {
		parent[rx] = ry
		if rank[rx] == rank[ry] {
			rank[ry]++
		}
	}
}
