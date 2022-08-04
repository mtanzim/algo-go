package graph

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Graph[T any] struct {
	v   int
	e   int
	adj []map[int]T
}

func (g *Graph[T]) checkVertex(v int) error {
	if v < 0 || v > (g.v-1) {
		return errors.New("invalid vertex")
	}
	return nil
}

func New[T any](v int) (*Graph[T], error) {
	if v < 0 {
		return nil, errors.New("cannot have negative number of vertices")
	}
	g := &Graph[T]{v: v, e: 0}
	g.adj = make([]map[int]T, v)
	for i := 0; i < v; i++ {
		g.adj[i] = make(map[int]T)
	}
	return g, nil
}

type UnweightedEdgeType = struct{}

type edge struct {
	from int
	to   int
}
type graphData struct {
	numVertices int
	edges       []edge
}

func getGraphFromScanner(in *bufio.Scanner) (*graphData, error) {
	var numVertices int
	edges := []edge{}
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

		edges = append(edges, edge{from, to})
		if err != nil {
			return nil, err
		}
		if err := in.Err(); err != nil {
			return nil, err
		}

	}
	return &graphData{
		numVertices: numVertices,
		edges:       edges,
	}, nil
}

func NewUndirectedFromScanner(in *bufio.Scanner) (*Graph[UnweightedEdgeType], error) {
	var g *Graph[UnweightedEdgeType]
	gd, err := getGraphFromScanner(in)
	if err != nil {
		return nil, err
	}
	g, err = New[UnweightedEdgeType](gd.numVertices)
	if err != nil {
		return nil, err
	}
	for _, edge := range gd.edges {
		err = g.AddEdgeUndirected(edge.from, edge.to, struct{}{})
		if err != nil {
			return nil, err
		}
	}
	return g, nil
}

func NewDigraphFromScanner(in *bufio.Scanner) (*Graph[UnweightedEdgeType], error) {
	var g *Graph[UnweightedEdgeType]
	gd, err := getGraphFromScanner(in)
	if err != nil {
		return nil, err
	}
	g, err = New[UnweightedEdgeType](gd.numVertices)
	if err != nil {
		return nil, err
	}
	for _, edge := range gd.edges {
		err = g.AddEdge(edge.from, edge.to, struct{}{})
		if err != nil {
			return nil, err
		}
	}
	return g, nil
}

func (g *Graph[T]) addEdge(from, to int, edgeVal T) error {
	err := g.checkVertex(from)
	if err != nil {
		return err
	}
	err = g.checkVertex(to)
	if err != nil {
		return err
	}

	curEdge := g.adj[from]
	curEdge[to] = edgeVal
	g.adj[from] = curEdge
	return nil
}

func (g *Graph[T]) AddEdge(from, to int, edgeVal T) error {
	err := g.addEdge(from, to, edgeVal)
	if err != nil {
		return err
	}
	g.e++
	return nil
}

func (g *Graph[T]) AddEdgeUndirected(from, to int, edgeVal T) error {
	err := g.addEdge(to, from, edgeVal)
	if err != nil {
		return err
	}

	err = g.addEdge(from, to, edgeVal)
	if err != nil {
		return err
	}
	g.e++
	return nil
}

func (g *Graph[T]) Vertices() int {
	return g.v
}

func (g *Graph[T]) Edges() int {
	return g.e
}

func (g *Graph[T]) Adj(v int) map[int]T {
	return g.adj[v]
}

func (g *Graph[T]) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("%d vertices, ", g.v))
	s.WriteString(fmt.Sprintf("%d edges\n", g.e))

	for i, v := range g.adj {
		s.WriteString(fmt.Sprintf("%d: %v\n", i, v))
	}

	return s.String()
}
