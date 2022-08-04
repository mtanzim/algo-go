package graph

import (
	"errors"

	"github.com/mtanzim/algo-go/stack"
)

func GetTopologicalSort[T any](g *Graph[T]) (*stack.Stack, error) {
	dfsCycle := NewDirectedCycle(g)
	if dfsCycle.hasCycle() {
		return nil, errors.New("cannot find order, has cycle")
	}
	dfsOrder := NewDepthFirstOrder(g)
	return dfsOrder.reversePost, nil

}
