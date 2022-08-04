package graph

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type EdgeWeightedGraph struct {
	g *Graph[*WeightedEdge]
}

func NewEdgeWeighted(v int) (*EdgeWeightedGraph, error) {
	if v < 0 {
		return nil, errors.New("cannot have negative number of vertices")
	}
	g := &Graph[*WeightedEdge]{v: v, e: 0}
	g.adj = make([]map[int]*WeightedEdge, v)
	for i := 0; i < v; i++ {
		g.adj[i] = make(map[int]*WeightedEdge)
	}
	return &EdgeWeightedGraph{g}, nil
}

func (wg *EdgeWeightedGraph) AddEdge(edge *WeightedEdge) {
	v := edge.Either()
	w, _ := edge.Other(v)
	wg.g.AddEdgeUndirected(v, w, edge)
	wg.g.e++
}

func (wg *EdgeWeightedGraph) V() int {
	return wg.g.Vertices()
}

func (wg *EdgeWeightedGraph) E() int {
	return wg.g.Edges()
}

func (wg *EdgeWeightedGraph) Adj(v int) []*WeightedEdge {

	simplifiedEdges := []*WeightedEdge{}
	detailedEdge := wg.g.adj[v]
	for _, v := range detailedEdge {
		simplifiedEdges = append(simplifiedEdges, v)
	}
	return simplifiedEdges
}

func (wg *EdgeWeightedGraph) Edges(edge *WeightedEdge) []*WeightedEdge {
	edges := []*WeightedEdge{}
	for v := 0; v < wg.V(); v++ {
		curAdj := wg.g.adj[v]
		for _, e := range curAdj {
			other, _ := e.Other(v)
			if other > v {
				edges = append(edges, e)
			}
		}
	}
	return edges
}

func (wg *EdgeWeightedGraph) String() string {
	return wg.g.String()
}

func getWeightedGraphFromScanner(in *bufio.Scanner) (*EdgeWeightedGraph, error) {
	var numVertices int
	edges := []*WeightedEdge{}
	for i := 0; in.Scan(); i++ {
		line := in.Text()
		if i == 0 {
			v, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}
			numVertices = v
			continue
		}
		if i == 1 {
			continue
		}
		words := strings.Fields(line)
		if len(words) < 2 {
			return nil, fmt.Errorf("error in input on line: %d", i)
		}
		from, err := strconv.Atoi(words[0])
		if err != nil {
			return nil, err
		}
		to, err := strconv.Atoi(words[1])
		if err != nil {
			return nil, err
		}

		weight, err := strconv.ParseFloat(words[2], 64)
		if err != nil {
			return nil, err
		}

		newEdge := NewEdge(from, to, weight)

		edges = append(edges, newEdge)
		if err != nil {
			return nil, err
		}
		if err := in.Err(); err != nil {
			return nil, err
		}

	}

	g, err := NewEdgeWeighted(numVertices)
	if err != nil {
		return nil, err
	}
	for _, edge := range edges {
		g.AddEdge(edge)
	}
	return g, nil
}

func NewWeightedFromScanner(in *bufio.Scanner) (*EdgeWeightedGraph, error) {
	g, err := getWeightedGraphFromScanner(in)
	return g, err
}
