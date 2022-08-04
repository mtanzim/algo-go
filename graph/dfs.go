package graph

import (
	"fmt"
	"strings"

	"github.com/mtanzim/algo-go/stack"
)

type DFS struct {
	marked []bool
	edgeTo []int
	count  int
	graph  *Graph[UnweightedEdgeType]
	source int
}

func NewDFS(g *Graph[UnweightedEdgeType], source int) *DFS {
	marked := make([]bool, g.Vertices())
	edgeTo := make([]int, g.Vertices())
	// since 0 is a valid vertex key, let's initialize all edge to members to -1
	for i := 0; i < g.Vertices(); i++ {
		edgeTo[i] = -1
	}
	dfs := &DFS{marked: marked, graph: g, source: source, edgeTo: edgeTo}
	dfs.dfs(g, source)
	return dfs
}

func (dfs *DFS) dfs(g *Graph[UnweightedEdgeType], v int) {
	dfs.marked[v] = true
	dfs.count++
	adjVertices := make([]int, len(g.adj[v]))
	i := 0
	for k := range g.adj[v] {
		adjVertices[i] = k
		i++
	}
	for _, neighbor := range adjVertices {
		if !dfs.marked[neighbor] {
			dfs.edgeTo[neighbor] = v
			dfs.dfs(g, neighbor)
		}
	}
}

func (dfs *DFS) HasPathTo(v int) bool { return dfs.marked[v] }
func (dfs *DFS) Count(v int) int      { return dfs.count }

func (dfs *DFS) pathTo(v int) *stack.Stack {
	if !dfs.HasPathTo(v) {
		return nil
	}
	s := stack.NewStack(v)
	next := dfs.edgeTo[v]
	for next != dfs.source && next != -1 {
		s.Push(next)
		next = dfs.edgeTo[next]
	}
	return s
}

func (dfs *DFS) String() string {

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("source: %d, connected: ", dfs.source))
	for i, isMarked := range dfs.marked {
		if isMarked && i != dfs.source {
			sb.WriteString(fmt.Sprintf("%d,", i))
		}
	}
	sb.WriteString("\n")
	sb.WriteString("paths\n")
	for i := 0; i < dfs.graph.Vertices(); i++ {
		if dfs.HasPathTo(i) && i != dfs.source {
			sb.WriteString(fmt.Sprintf("%d -> %s\n", dfs.source, dfs.pathTo(i)))
		}
	}

	return sb.String()
}
