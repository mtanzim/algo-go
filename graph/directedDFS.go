package graph

import (
	"fmt"
	"strings"
)

type DirectedDFS struct {
	marked []bool
}

func NewDirectedDFS(g *Graph[UnweightedEdgeType], sources []int) *DirectedDFS {
	marked := make([]bool, g.Vertices())
	ddfs := &DirectedDFS{marked: marked}
	for _, v := range sources {
		ddfs.dfs(g, v)
	}
	return ddfs
}

func (ddfs *DirectedDFS) dfs(g *Graph[UnweightedEdgeType], v int) {
	ddfs.marked[v] = true
	adjVertices := make([]int, len(g.adj[v]))
	i := 0
	for k := range g.adj[v] {
		adjVertices[i] = k
		i++
	}
	for _, neighbor := range adjVertices {
		if !ddfs.marked[neighbor] {
			ddfs.dfs(g, neighbor)
		}
	}
}

func (ddfs *DirectedDFS) String() string {

	var sb strings.Builder
	sb.WriteString("reachable nodes: \n")
	for i, isMarked := range ddfs.marked {
		if isMarked {
			sb.WriteString(fmt.Sprintf("%d,", i))
		}
	}

	return sb.String()
}
