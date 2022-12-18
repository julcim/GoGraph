package goGraph

import (
	"fmt"
	"sort"
)

// Kruskal runs the Kruskal algorithm on a given Graph and outputs a Graph pointer to a
// Minimum Spanning Tree of that graph.
func Kruskal(g *Graph) *Graph {
	var t Graph
	t.size = g.size
	t.graph = make([]adjList, t.size)
	InitializeGraph(&t)
	edges := make([]Edge, 0)
	fmt.Println(g)
	for i := 0; i < g.size; i++ {
		for u := range g.graph[i].list {
			var e Edge
			e.start = i
			e.end = u
			e.weight = GetWeight(g, i, u)
			edges = append(edges, e)
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight > edges[j].weight
	})

	// reverse := []Edge{}
	// for i := range edges {
	// 	// reverse the order
	// 	reverse = append(reverse, edges[len(edges)-1-i])
	// }
	// edges = reverse
	fmt.Println("edges", edges)
	rank := make([]int, g.size)
	parent := make([]int, g.size)
	for i := 0; i < GetSize(g); i++ {
		rank[i] = 0
		parent[i] = i
	}
	for i := 0; i < len(edges)-1; i++ {
		if find(edges[i].start, parent) != find(edges[i].end, parent) {
			weight := GetWeight(g, edges[i].start, edges[i].end)
			AddEdge(&t, edges[i].start, edges[i].end, weight)
			AddEdge(&t, edges[i].end, edges[i].start, weight)
		}
		union(edges[i].end, edges[i].start, parent, rank)
	}
	return &t

}

func reverse(arr []Edge) []Edge {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
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
