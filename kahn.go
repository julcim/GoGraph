package GoGraph

// Kahn runs Kahn's algorithm for topologically sorting a Directed Acyclic Graph (DAG).
// For a DAG, Kahn will return an output slice in topologically sorted order from index 0.
// The behavior of this function is undefined if not called on a DAG.
func (g *Graph) Kahn() (output []int) {
	queue := make([]int, 0)
	output = make([]int, 0)
	indegree := make([]int, g.size)
	for i := 0; i < g.size; i++ {
		for u := range g.graph[i].list {
			indegree[u]++
		}
	}
	for i := 0; i < g.size; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]
		output = append(output, v)
		for u := range g.graph[v].list {
			indegree[u]--
			if indegree[u] == 0 {
				queue = append(queue, u)
			}
		}
	}
	return output
}
