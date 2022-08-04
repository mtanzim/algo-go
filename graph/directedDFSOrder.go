package graph

import (
	"strings"

	"github.com/mtanzim/algo-go/queue"
	"github.com/mtanzim/algo-go/stack"
)

type DepthFirstOrder[T any] struct {
	marked      []bool
	pre         *queue.Queue
	post        *queue.Queue
	reversePost *stack.Stack
}

func NewDepthFirstOrder[T any](g *Graph[T]) *DepthFirstOrder[T] {
	marked := make([]bool, g.Vertices())
	dfso := &DepthFirstOrder[T]{marked: marked}
	for v := 0; v < g.Vertices(); v++ {
		if !dfso.marked[v] {
			dfso.dfs(g, v)
		}
	}
	return dfso
}

func (dfso *DepthFirstOrder[T]) dfs(g *Graph[T], v int) {

	if dfso.pre == nil {
		dfso.pre = queue.NewQueue(v)
	} else {
		dfso.pre.Enqueue(v)
	}
	dfso.marked[v] = true
	adjVertices := make([]int, len(g.adj[v]))
	i := 0
	for k := range g.adj[v] {
		adjVertices[i] = k
		i++
	}
	for _, neighbor := range adjVertices {
		if !dfso.marked[neighbor] {
			dfso.dfs(g, neighbor)
		}
	}
	if dfso.post == nil {
		dfso.post = queue.NewQueue(v)
	} else {
		dfso.post.Enqueue(v)
	}

	if dfso.reversePost == nil {
		dfso.reversePost = stack.NewStack(v)
	} else {
		dfso.reversePost.Push(v)
	}

}

func (dfso *DepthFirstOrder[T]) String() string {

	var sb strings.Builder
	sb.WriteString("pre: \n")
	sb.WriteString(dfso.pre.String())
	sb.WriteString("\npost: \n")
	sb.WriteString(dfso.post.String())
	sb.WriteString("\nreverse post: \n")
	sb.WriteString(dfso.reversePost.String())
	return sb.String()
}
