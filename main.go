package main

import (
	"fmt"
	dk "geo-algors/dijkstra-go/dijkstra"
)

func main() {
	//var newGraph = dk.HighLevelGraph[int]{}
	var newGraph = dk.NewGraph[int, string]()

	type edges[T dk.Number, TN dk.Ordered] struct {
		fromNode TN
		toNode   TN
		weight   T
	}

	var edgesContainer = []edges[int, string]{
		{"X", "A", 7},
		{"A", "B", 3},
		{"A", "D", 4},
		{"X", "B", 2},
		{"X", "C", 3},
		{"B", "H", 5},
		{"B", "D", 4},
		{"C", "L", 2},
		{"D", "F", 1},
		{"X", "E", 4},
		{"F", "H", 3},
		{"G", "Y", 2},
		{"G", "H", 2},
		{"I", "J", 6},
		{"I", "L", 4},
		{"I", "K", 4},
		{"K", "Y", 5},
		{"J", "L", 1},
	}

	for _, edge := range edgesContainer {
		newGraph.AddEdge(edge.fromNode, edge.toNode, edge.weight)
	}

	fmt.Println(newGraph.GetWeights())
	fmt.Println(newGraph.GetEdges())

}
