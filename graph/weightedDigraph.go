package graph

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type EdgeWeightedDigraph struct {
	g *Graph[*DirectedEdge]
}

func NewEdgeWeightedDigraph(v int) (*EdgeWeightedDigraph, error) {
	if v < 0 {
		return nil, errors.New("cannot have negative number of vertices")
	}
	g := &Graph[*DirectedEdge]{v: v, e: 0}
	g.adj = make([]map[int]*DirectedEdge, v)
	for i := 0; i < v; i++ {
		g.adj[i] = make(map[int]*DirectedEdge)
	}
	return &EdgeWeightedDigraph{g}, nil
}

func (wdg *EdgeWeightedDigraph) AddEdge(edge *DirectedEdge) {
	v := edge.From()
	w := edge.To()
	wdg.g.AddEdge(v, w, edge)
}

func (wdg *EdgeWeightedDigraph) V() int {
	return wdg.g.Vertices()
}

func (wdg *EdgeWeightedDigraph) E() int {
	return wdg.g.Edges()
}

func (wdg *EdgeWeightedDigraph) Adj(v int) []*DirectedEdge {

	simplifiedEdges := []*DirectedEdge{}
	detailedEdge := wdg.g.adj[v]
	for _, v := range detailedEdge {
		simplifiedEdges = append(simplifiedEdges, v)
	}
	return simplifiedEdges
}

func (wdg *EdgeWeightedDigraph) Edges() []*DirectedEdge {
	edges := []*DirectedEdge{}
	for v := 0; v < wdg.V(); v++ {
		curAdj := wdg.g.adj[v]
		for _, e := range curAdj {
			edges = append(edges, e)
		}
	}
	return edges
}

func getWeightedDiraphFromScanner(in *bufio.Scanner) (*EdgeWeightedDigraph, error) {
	var numVertices int
	edges := []*DirectedEdge{}
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

		newEdge := NewDirectedEdge(from, to, weight)

		edges = append(edges, newEdge)
		if err != nil {
			return nil, err
		}
		if err := in.Err(); err != nil {
			return nil, err
		}

	}

	g, err := NewEdgeWeightedDigraph(numVertices)
	if err != nil {
		return nil, err
	}
	for _, edge := range edges {
		g.AddEdge(edge)
	}
	return g, nil
}

func NewWeightedDigraphFromScanner(in *bufio.Scanner) (*EdgeWeightedDigraph, error) {
	g, err := getWeightedDiraphFromScanner(in)
	return g, err
}

func (wdg *EdgeWeightedDigraph) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("%d vertices, ", wdg.g.Vertices()))
	s.WriteString(fmt.Sprintf("%d edges\n", wdg.g.Edges()))

	for v := 0; v < wdg.g.Vertices(); v++ {
		s.WriteString(fmt.Sprintf("%d: ", v))
		curEdges := wdg.Adj(v)
		for _, edge := range curEdges {
			s.WriteString(fmt.Sprintf("%v\t", edge))
		}
		s.WriteString("\n")
	}

	return s.String()
}
