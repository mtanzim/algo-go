package graph

import (
	"fmt"
	"strings"

	"github.com/mtanzim/algo-go/stack"
)

type DirectedCycle[T any] struct {
	marked  []bool
	onStack []bool
	edgeTo  []int
	cycle   *stack.Stack
}

func NewDirectedCycle[T any](g *Graph[T]) *DirectedCycle[T] {
	marked := make([]bool, g.Vertices())
	onStack := make([]bool, g.Vertices())
	edgeTo := make([]int, g.Vertices())
	// since 0 is a valid vertex key, let's initialize all edge to members to -1
	for i := 0; i < g.Vertices(); i++ {
		edgeTo[i] = -1
	}

	dc := &DirectedCycle[T]{
		marked:  marked,
		edgeTo:  edgeTo,
		cycle:   nil,
		onStack: onStack,
	}

	for v := 0; v < g.Vertices(); v++ {
		if !dc.marked[v] {
			dc.dfs(g, v)
		}
	}
	return dc

}

func (dc *DirectedCycle[T]) hasCycle() bool {
	return dc.cycle != nil
}

func (dc *DirectedCycle[T]) dfs(g *Graph[T], v int) {
	dc.marked[v] = true
	dc.onStack[v] = true
	adjVertices := make([]int, len(g.adj[v]))
	i := 0
	for k := range g.adj[v] {
		adjVertices[i] = k
		i++
	}
	for _, neighbor := range adjVertices {
		if dc.hasCycle() {
			return
		}
		if !dc.marked[neighbor] {
			dc.edgeTo[neighbor] = v
			dc.dfs(g, neighbor)
		} else if dc.onStack[neighbor] {
			// cycle found, indentify the cycle path
			dc.cycle = stack.NewStack(v)
			next := dc.edgeTo[v]
			for next != neighbor && next != -1 {
				dc.cycle.Push(next)
				next = dc.edgeTo[next]
			}
			dc.cycle.Push(neighbor)
			dc.cycle.Push(v)

		}
	}
	dc.onStack[v] = false

}

func (dc *DirectedCycle[T]) Cycle() ([]int, error) {
	cycleCopy := *dc.cycle
	cycleSlice := []int{}
	for !cycleCopy.IsEmpty() {
		curVal, err := cycleCopy.Pop()
		if err != nil {
			return nil, err
		}
		cycleSlice = append(cycleSlice, curVal)
	}
	return cycleSlice, nil
}
func (dc *DirectedCycle[T]) String() string {

	var sb strings.Builder
	sb.WriteString("cycle: \n")
	sb.WriteString(fmt.Sprintf("%s,", dc.cycle))

	return sb.String()
}
