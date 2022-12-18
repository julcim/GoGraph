package goGraph

import (
	"fmt"
	"math"
)

// func remove(slice []int, value int) []int {
// 	result := []int{}
// 	for _, v := range slice {
// 		if v != value {
// 			result = append(result, v)
// 		}
// 	}
// 	copy(slice, result)
// 	return slice[:len(result)]
// }

func Dijkstra(g *Graph, source int) []int {
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
