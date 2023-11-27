package dijkstra

import (
	types "golang.org/x/exp/constraints"
)

/*
	HighLevelGraph

edges is a map of all possible next nodes

	e.g. {'X': ['A', 'B', 'C', 'E'], ...}
	weights has all the weights between two nodes,
	with the two nodes as a tuple as the key
	e.g. {[ 'X', 'A' ]: 7, [ 'X', 'B' ]: 2, ...}
*/
type Number interface {
	types.Float | types.Integer
}

type NodePair [2]string

type HighLevelGraph[T Number] struct {
	edges   map[string][]string
	weights map[NodePair]T // weights = {['A', 'B']: 2}
}

func NewGraph[T Number]() *HighLevelGraph[T] {
	return &HighLevelGraph[T]{}
}

func (hlg *HighLevelGraph[T]) AddEdge(fromNode, toNode string, weight T) {
	if hlg.edges == nil {
		hlg.edges = make(map[string][]string)
	}
	if hlg.weights == nil {
		hlg.weights = make(map[NodePair]T)
	}
	hlg.edges[fromNode] = append(hlg.edges[fromNode], toNode)
	hlg.edges[toNode] = append(hlg.edges[toNode], fromNode)
	hlg.weights[NodePair{fromNode, toNode}] = weight
}

func (hlg *HighLevelGraph[T]) GetEdges() map[string][]string {
	return hlg.edges
}

func (hlg *HighLevelGraph[T]) GetWeights() map[NodePair]T {
	return hlg.weights
}
