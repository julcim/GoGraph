# GoGraph
Final project for CIS 1930. Implements a graph data structure and graph algorithms for Go.

## Installing
Run `$ go get github.com/julcim/GoGraph` in your terminal to install the package, and then use it as you desire. 

## Simple Usages
```
var g Graph
g.size = 3
g.graph = make([]adjList, 3)
g.InitializeGraph()
g.AddEdge(0, 1, 3)
g.AddEdge(0, 2, 1)
g.AddEdge(2, 3, 1)
g.AddEdge(1, 3, 4)
distances := g.Dijkstra(0)
// distance[0] = 0
// distance[1] = 3
// distance[2] = 1
// distance[3] = 2
```

## Documentation
Documentation for the package can be found [here.](https://pkg.go.dev/github.com/julcim/gograph)

Note: When using this library avoid reading and writing to a graph concurrently or writing concurrently to maintain thread safety. Reading concurrently will not pose a threat to thread safety.

## License
GoGraph uses MIT License
