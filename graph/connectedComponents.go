package graph

import "errors"

type ConnectedComponets[T any] struct {
	marked []bool
	id     []int
	count  int
}

func NewConnectedComponents[T any](g *Graph[T]) *ConnectedComponets[T] {
	marked := make([]bool, g.Vertices())
	id := make([]int, g.Vertices())
	for i := 0; i < len(id); i++ {
		id[i] = -1
	}
	count := 0
	cc := &ConnectedComponets[T]{marked: marked, id: id, count: count}

	for v := 0; v < g.Vertices(); v++ {
		if !cc.marked[v] {
			cc.dfs(g, v)
			cc.count++
		}
	}

	return cc
}

func (cc *ConnectedComponets[T]) dfs(g *Graph[T], v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	neighbors := g.Adj(v)
	for w := range neighbors {
		if !cc.marked[w] {
			cc.dfs(g, w)
		}
	}
}

func (cc *ConnectedComponets[T]) Count() int {
	return cc.count
}

func (cc *ConnectedComponets[T]) AssertFullyConnected() error {
	if cc.Count() != 1 {
		return errors.New("graph is not fully connected")
	}
	return nil
}
