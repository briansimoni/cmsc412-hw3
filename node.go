package main

import (
	"reflect"
)

type node struct {
	id       string
	edges    []edge

	// using distance and parent for breadth first search algorithm
	distance int
	parent   string // id of the parent node

	closenessCentrality float64
}

func (n node) WeightedDegreeCentrality() float64 {

	sum := 0.0

	for i := range n.edges {
		sum += n.edges[i].Weight
	}

	return sum
}

func (n *node) AddEdge(e edge) {
	for i := range n.edges {
		if reflect.DeepEqual(e, n.edges[i]) {
			return
		}
	}

	n.edges = append(n.edges, e)
}

func (n *node) RemoveEdge(e edge) {
	//flip it first
	placeHolder := e.toNode
	e.toNode = e.fromNode
	e.fromNode = placeHolder
	es := make([]edge, 0)
	for i := range n.edges {
		if reflect.DeepEqual(e, n.edges[i]) {
			continue
		}
		es = append(es, n.edges[i])
	}
	n.edges = es
}
