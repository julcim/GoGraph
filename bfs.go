package goGraph

func Bfs(g *Graph, source int) []int {
	var discovered []bool
	var parent []int
	for i := 0; i < g.size; i++ {
		discovered[i] = false
		parent[i] = -1
	}
	queue := make(chan int, g.size)
	queue <- source
	discovered[source] = true
	for len(queue) > 0 {
		v := <-queue
		for u := range g.graph[v].list {
			if !discovered[u] {
				discovered[u] = true
				queue <- u
				parent[u] = v
			}
		}
	}
	return parent

}
