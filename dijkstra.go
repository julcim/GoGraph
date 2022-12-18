package goGraph

import (
	"fmt"
	"math"
)

// Dijkstra runs the known Dijkstra algorithm on a given Graph and source.
// It outputs the shortest path from the source to all vertices in the graph in the
// form of an array. This array is structured such that distance[a] is the
// distance from the source to a.
func Dijkstra(g *Graph, source int) (distance []int) {
	fmt.Println("entered func")
	dist := make([]int, g.size)
	parent := make([]int, g.size)
	queue := make(map[int]int, 0)
	for x := range g.graph {
		dist[x] = 1000000000
		parent[x] = -1
	}
	dist[source] = 0
	for i := 0; i < g.size; i++ {
		queue[i] = dist[i]
	}
	for len(queue) > 0 {
		minDistance := math.MaxInt32
		u := -1
		for i, distance := range queue {
			if distance < minDistance {
				minDistance = distance
				u = i
			}
		}
		delete(queue, u)
		for v := range g.graph[u].list {
			fmt.Println("v=", v)
			if dist[v] > dist[u]+GetWeight(g, u, v) {
				dist[v] = dist[u] + GetWeight(g, u, v)
				parent[v] = u

			}
		}
	}
	fmt.Println(dist)
	return dist
}
