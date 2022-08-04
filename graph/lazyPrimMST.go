package graph

import (
	priorityQueue "github.com/jupp0r/go-priority-queue"
)

type LazyPrimMST struct {
	marked []bool
	minPQ  priorityQueue.PriorityQueue
	queue  []*WeightedEdge
}

func NewLazyPrimMST(g *EdgeWeightedGraph) (*LazyPrimMST, error) {

	cc := NewConnectedComponents[*WeightedEdge](g.g)
	err := cc.AssertFullyConnected()
	if err != nil {
		return nil, err
	}
	minPQ := priorityQueue.New()
	marked := make([]bool, g.V())
	for i := 0; i < len(marked); i++ {
		marked[i] = false
	}
	queue := []*WeightedEdge{}
	mst := &LazyPrimMST{marked: marked, minPQ: minPQ, queue: queue}
	mst.visit(g, 0)
	for mst.minPQ.Len() != 0 {
		e, err := mst.minPQ.Pop()
		if err != nil {
			return nil, err
		}
		eCasted := e.(*WeightedEdge)
		v := eCasted.Either()
		w, err := eCasted.Other(v)
		if err != nil {
			return nil, err
		}
		if mst.marked[v] && mst.marked[w] {
			continue
		}
		// improve this; should not be an array
		mst.queue = append(mst.queue, eCasted)
		if !marked[v] {
			mst.visit(g, v)
		}
		if !marked[w] {
			mst.visit(g, w)
		}

	}
	return mst, nil
}

func (mst *LazyPrimMST) visit(g *EdgeWeightedGraph, v int) error {
	mst.marked[v] = true
	for _, edge := range g.Adj(v) {
		other, err := edge.Other(v)
		if err != nil {
			return err
		}
		if !mst.marked[other] {
			mst.minPQ.Insert(edge, edge.Weight())
		}
	}
	return nil
}

func (mst *LazyPrimMST) Weight() float64 {
	sum := 0.0
	for _, edge := range mst.queue {
		sum += edge.Weight()
	}
	return sum
}
