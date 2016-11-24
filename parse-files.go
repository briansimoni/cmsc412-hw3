// This files provides utilities for parsing the text files that contain the YouTube data
// it will use simple regex to split the values and then ultimately join them together into nodes
// then it will take all of the nodes and create a graph

package main

import (
	"os"
	"bufio"
	"strings"
	"errors"
)

func createGraphFromFile(fileName string) graph {
	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)


	g := NewGraph()

	scanner.Scan()

	for scanner.Scan() {
		err, n := createNodeFromText(scanner.Text())
		if err != nil {
			continue
		}
		g.InsertNode(n)

		undirected := createUndirectedEdge(n)
		g.InsertNode(undirected)
	}

	return g
}

func createNodeFromText(line string) (error, node) {
	values := strings.Split(line, " ")

	if len(values) < 2 {
		e := errors.New("incomplete data")
		return e, node{}
	}

	id := values[0]
	toNode := values[1]


	edges := make([]edge, 0)

	e := edge{fromNode:id, toNode:toNode}
	edges = append(edges, e)

	n := node{id:id, edges:edges}

	return nil, n
}


// ultimately will create a graph equivalent to an undirected graph
func createUndirectedEdge(n node) node {
	id := n.edges[0].toNode
	e := edge{toNode:n.edges[0].fromNode, fromNode:n.edges[0].toNode}
	edges := make([]edge, 0)
	edges = append(edges, e)
	undirected := node{id:id, edges:edges}
	return undirected
}

