package goGraph

import (
	"sort"
)

func Kruskal(g *Graph) *Graph {
	var t *Graph
	edges := make([]Edge, 1)
	for i := 0; i < GetSize(g); i++ {
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

	var rank []int
	var parent []int
	for i := 0; i < GetSize(g); i++ {
		rank[i] = 0
		parent[i] = i
	}
	for i := 0; i < len(edges)-1; i++ {
		if find(edges[i].start, parent) != find(edges[i].end, parent) {
			AddEdge(t, edges[i].start, edges[i].end, GetWeight(g, edges[i].start, edges[i].end))
			AddEdge(t, edges[i].end, edges[i].start, GetWeight(g, edges[i].start, edges[i].end))
		}
		union(edges[i].end, edges[i].start, parent, rank)
	}
	return t

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
