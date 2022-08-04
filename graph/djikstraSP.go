package graph

import (
	"math"

	priorityQueue "github.com/jupp0r/go-priority-queue"
)

type DjikstraSP struct {
	edgeTo []*DirectedEdge
	distTo []float64
	minPQ  priorityQueue.PriorityQueue
}

func NewDjikstraSP(g *EdgeWeightedDigraph, s int) (*DjikstraSP, error) {
	edgeTo := make([]*DirectedEdge, g.g.Vertices())
	distTo := make([]float64, g.g.Vertices())
	minPQ := priorityQueue.New()
	for i := 0; i < g.g.Vertices(); i++ {
		distTo[i] = math.Inf(1)
	}
	distTo[s] = 0.0
	minPQ.Insert(s, 0.0)
	sp := &DjikstraSP{edgeTo: edgeTo, distTo: distTo, minPQ: minPQ}
	for sp.minPQ.Len() != 0 {
		minVertex, err := sp.minPQ.Pop()
		if err != nil {
			return nil, err
		}
		sp.relax(g, minVertex.(int))
	}
	return sp, nil
}

func (sp *DjikstraSP) relax(g *EdgeWeightedDigraph, v int) {
	adjEdges := g.Adj(v)
	for _, edge := range adjEdges {
		w := edge.To()
		// found a shorter path to w from s
		if sp.distTo[w] > sp.distTo[v]+edge.Weight() {
			sp.distTo[w] = sp.distTo[v] + edge.Weight()
			sp.edgeTo[w] = edge
			// no-op if item doesn't exist
			sp.minPQ.UpdatePriority(w, edge.weight)
			// no-op if item exists
			sp.minPQ.Insert(w, edge.weight)
		}
	}
}

func (sp *DjikstraSP) DistTo(v int) float64 {
	return sp.distTo[v]
}

func (sp *DjikstraSP) HasPathTo(v int) bool {
	return sp.distTo[v] < math.Inf(1)
}

func (sp *DjikstraSP) PathTo(v int) []*DirectedEdge {
	stack := []*DirectedEdge{}
	curEdge := sp.edgeTo[v]
	for curEdge != nil {
		stack = append(stack, curEdge)
		curEdge = sp.edgeTo[curEdge.From()]
	}
	return stack
}
