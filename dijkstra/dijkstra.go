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

type Ordered interface {
	types.Ordered
}

type NodePair[TN Ordered] [2]TN

type NodeAndWeight[T Ordered] struct {
	node *Node[T]
}

type Node[T Ordered] struct {
	key T
}

type HighLevelGraph[T Number, TNode Ordered] struct {
	edges   map[TNode][]TNode     // Ordered types Number, String..
	weights map[NodePair[TNode]]T // weights = {['A', 'B']: 2}
}

func NewGraph[T Number, TNode Ordered]() *HighLevelGraph[T, TNode] {
	return &HighLevelGraph[T, TNode]{}
}

func (hlg *HighLevelGraph[T, TN]) AddEdge(fromNode, toNode TN, weight T) {
	if hlg.edges == nil {
		hlg.edges = make(map[TN][]TN)
	}
	if hlg.weights == nil {
		hlg.weights = make(map[NodePair[TN]]T)
	}
	hlg.edges[fromNode] = append(hlg.edges[fromNode], toNode)
	hlg.edges[toNode] = append(hlg.edges[toNode], fromNode)
	hlg.weights[NodePair[TN]{fromNode, toNode}] = weight
}

func (hlg *HighLevelGraph[T, TN]) GetEdges() map[TN][]TN {
	return hlg.edges
}

func (hlg *HighLevelGraph[T, TN]) GetWeights() map[NodePair[TN]]T {
	return hlg.weights
}

func Dijkstra[T Number, TN Ordered](graph *HighLevelGraph[T, TN], initial, end TN) {
	//  shortest paths is a map of nodes
	//  whose value is a tuple of (previous node, weight)
}
