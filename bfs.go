package goGraph

func Bfs(g *Graph, source int) []int {
	discovered := make([]bool, g.size)
	parent := make([]int, g.size)
	for i := 0; i < g.size; i++ {
		discovered[i] = false
		parent[i] = -1
	}
	queue := make([]int, 0)
	queue = append(queue, source)
	discovered[source] = true
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for u := range g.graph[v].list {
			if !discovered[u] {
				discovered[u] = true
				queue = append(queue, u)
				parent[u] = v
			}
		}
	}
	return parent

}
