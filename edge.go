package goGraph

type Edge struct {
	start  int
	end    int
	weight int
}

func GetEdgeWeight(e *Edge) int {
	return e.weight
}
