package main

import (
	"fmt"
	dk "geo-algors/dijkstra-go/dijkstra"
)

func main() {
	var newGraph = dk.HighLevelGraph[int]{}

	type edges[T dk.Number] struct {
		fromNode string
		toNode   string
		weight   T
	}

	var edgesContainer = []edges[int]{
		{"X", "A", 7},
		{"X", "B", 2},
		{"X", "C", 3},
		{"X", "E", 4},
		{"A", "B", 3},
		{"A", "D", 4},
		{"B", "D", 4},
		{"B", "H", 5},
		{"C", "L", 2},
		{"D", "F", 1},
		{"F", "H", 3},
		{"G", "H", 2},
		{"G", "Y", 2},
		{"I", "J", 6},
		{"I", "K", 4},
		{"I", "L", 4},
		{"J", "L", 1},
		{"K", "Y", 5},
	}

	for _, edge := range edgesContainer {
		newGraph.AddEdge(edge.fromNode, edge.toNode, edge.weight)
	}

	//fmt.Println(newGraph.GetWeights())
	fmt.Println(newGraph.GetEdges())

	//test()
}

func test() {
	dict := make(map[string]int)
	dict["Dennis"] = 22
	fmt.Println(dict)
}
