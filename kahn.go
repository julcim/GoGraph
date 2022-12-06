package goGraph

func Kahn(g *Graph) (output []int) {
	queue := make(chan int, g.size)
	output = make([]int, g.size)
	indegree := make([]int, g.size)
	for i := 0; i < g.size; i++ {
		for j := 0; j < len(g.graph[i].list); j++ {
			indegree[j]++
		}
	}
	for i := 0; i < g.size; i++ {
		if indegree[i] == 0 {
			queue <- i
		}
	}
	for len(queue) > 0 {
		v := <-queue
		output = append(output, v)
		for j := 0; j < len(g.graph[i].list); j++ {
			indegree[j]--
			if indegree[j] == 0 {
				queue <- j
			}
		}
	}
	return output
}
