package main

import (
	"math"
	"fmt"
)

const infinity = math.MaxInt32

// a graph is simply a collection of nodes and edges (the node struct contains edge data)
type graph struct {
	nodes map[string]node
	kCore int
}

func NewGraph() graph {
	nodes := make(map[string]node)
	return graph{nodes: nodes}
}

// function to see if the node is already in the graph
func (g *graph) IsInGraph(id string) bool {
	if _, ok := g.nodes[id]; ok {
		return true
	}
	return false
}

func (g *graph) InsertNode(n node) {
	if g.IsInGraph(n.id) {
		ref := g.nodes[n.id]
		ref.AddEdge(n.edges[0])
		g.nodes[n.id] = ref
	} else {
		g.nodes[n.id] = n
	}
}


// remove all of the relevant edges, then remove the node
func (g *graph) RemoveNode(n node) {
	id := n.id
	for i := range n.edges {
		toNode := n.edges[i].toNode
		e := n.edges[i]
		ref := g.nodes[toNode]
		ref.RemoveEdge(e)
		g.nodes[toNode] = ref
	}
	delete(g.nodes, id)
}


func (g *graph) maxKcore() ([]string, int){
	k := 1
	memory := make([]string, 0)

	for len(g.nodes) != 0{
		increment := true
		for _, node := range g.nodes {
			if len(node.edges) <= k {
				memory = append(memory, node.id)
				g.RemoveNode(node)
				increment = false
			}
		}
		if increment && len(g.nodes) != 0 {
			k++
			memory = memory [:0]
		}
	}
	return memory, k
}


func (g *graph) PrintGraphInformation() {
	for id, _ := range g.nodes {
		fmt.Println("Node", id)
		for i := 0; i < len(g.nodes[id].edges); i++ {
			fmt.Println(g.nodes[id].edges[i].toNode)
		}

	}
}





func (g *graph) initialize() {
	for id, node := range g.nodes {
		node.distance = infinity
		node.parent = ""
		g.nodes[id] = node
	}

}


// the breadth first search function sets the shortest path distance to all other nodes in g
func (g *graph) breadthFirstSearch(root string) {

	g.initialize()

	// create empty queue
	queue := make([]string, 0)

	r := g.nodes[root]
	r.distance = 0
	g.nodes[root] = r
	// enqueue root (push)
	queue = append(queue, root)

	for len(queue) != 0 {
		current := g.nodes[queue[0]]
		queue = queue[1:]
		for i := 0; i < len(current.edges); i ++ {

			// neighbor is a node
			neighbor := g.nodes[current.edges[i].toNode]
			if g.nodes[neighbor.id].distance == infinity {

				neighbor.distance = current.distance + 1
				neighbor.parent = current.id
				g.nodes[neighbor.id] = neighbor

				queue = append(queue, neighbor.id)
			}
		}
	}
}

// calculate the normalized closeness centrality for all nodes in g
func (g *graph) closenessCentrality() {
	totalNodes := float64(len(g.nodes))

	for nodeID, node := range g.nodes {
		g.breadthFirstSearch(nodeID)
		var sum float64 = 0.0
		for _, n := range g.nodes {
			sum += float64(n.distance)
		}

		sum = sum / (totalNodes - 1)
		node.closenessCentrality = math.Pow(sum, -1)
		g.nodes[nodeID] = node

	}
}


