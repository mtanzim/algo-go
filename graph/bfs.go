package graph

import (
	"fmt"
	"log"
	"strings"

	"github.com/mtanzim/algo-go/queue"
	"github.com/mtanzim/algo-go/stack"
)

type BFS struct {
	marked []bool
	edgeTo []int
	count  int
	queue  *queue.Queue
	graph  *Graph[UnweightedEdgeType]
	source int
}

func NewBFS(g *Graph[UnweightedEdgeType], source int) *BFS {
	marked := make([]bool, g.Vertices())
	edgeTo := make([]int, g.Vertices())
	// since 0 is a valid vertex key, let's initialize all edge to members to -1
	for i := 0; i < g.Vertices(); i++ {
		edgeTo[i] = -1
	}
	q := queue.NewQueue(source)
	marked[source] = true
	bfs := &BFS{marked: marked, graph: g, source: source, edgeTo: edgeTo, queue: q}
	bfs.bfs(g)
	return bfs
}

func (bfs *BFS) bfs(g *Graph[UnweightedEdgeType]) {

	for !bfs.queue.IsEmpty() {

		v, err := bfs.queue.Dequeue()
		if err != nil {
			log.Println("invalid case encountered")
			return
		}

		bfs.count++
		adjVertices := make([]int, len(g.adj[v]))
		i := 0
		for k := range g.adj[v] {
			adjVertices[i] = k
			i++
		}
		for _, neighbor := range adjVertices {
			if !bfs.marked[neighbor] {
				bfs.queue.Enqueue(neighbor)
				bfs.marked[neighbor] = true
				bfs.edgeTo[neighbor] = v
			}
		}
	}
}

func (bfs *BFS) HasPathTo(v int) bool { return bfs.marked[v] }
func (bfs *BFS) Count(v int) int      { return bfs.count }

func (bfs *BFS) pathTo(v int) *stack.Stack {
	if !bfs.HasPathTo(v) {
		return nil
	}
	s := stack.NewStack(v)
	next := bfs.edgeTo[v]
	for next != bfs.source {
		if next == -1 {
			break
		}
		s.Push(next)
		next = bfs.edgeTo[next]
	}
	return s
}

func (bfs *BFS) String() string {

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("source: %d, connected: ", bfs.source))
	for i, isMarked := range bfs.marked {
		if isMarked && i != bfs.source {
			sb.WriteString(fmt.Sprintf("%d,", i))
		}
	}
	sb.WriteString("\n")
	sb.WriteString("paths\n")
	for i := 0; i < bfs.graph.Vertices(); i++ {
		if bfs.HasPathTo(i) && i != bfs.source {
			sb.WriteString(fmt.Sprintf("%d -> %s\n", bfs.source, bfs.pathTo(i)))
		}
	}

	return sb.String()
}
