package graph

import (
	"math"

	"github.com/mtanzim/algo-go/queue"
)

type BellmanFordSP struct {
	edgeTo  []*DirectedEdge
	distTo  []float64
	onQueue []bool
	queue   *queue.Queue
	cost    int
	cycle   *DirectedCycle[*DirectedEdge]
}

// TODO: this isn't working
func NewBellmanFordSP(g *EdgeWeightedDigraph, s int) (*BellmanFordSP, error) {
	edgeTo := make([]*DirectedEdge, g.g.Vertices())
	onQueue := make([]bool, g.g.Vertices())
	queue := queue.NewQueue(s)
	onQueue[s] = true

	distTo := make([]float64, g.g.Vertices())
	for i := 0; i < g.g.Vertices(); i++ {
		distTo[i] = math.Inf(1)
	}
	distTo[s] = 0.0

	sp := &BellmanFordSP{edgeTo: edgeTo, distTo: distTo, onQueue: onQueue, queue: queue}

	for !sp.queue.IsEmpty() && !sp.hasNegativeCycle() {
		curVertex, err := sp.queue.Dequeue()
		if err != nil {
			return nil, err
		}
		onQueue[curVertex] = false
		sp.relax(g, curVertex)

	}

	return sp, nil
}

func (sp *BellmanFordSP) hasNegativeCycle() bool {
	return sp.cycle != nil
}

func (sp *BellmanFordSP) findNegativeCycle() error {
	numVertices := len(sp.edgeTo)
	spt, err := NewEdgeWeightedDigraph(numVertices)
	if err != nil {
		return nil
	}
	for i := 0; i < numVertices; i++ {
		if sp.edgeTo[i] != nil {
			spt.AddEdge(sp.edgeTo[i])
		}
	}
	sp.cycle = NewDirectedCycle[*DirectedEdge](spt.g)
	return nil

}

func (sp *BellmanFordSP) NegativeCycle() ([]int, error) {
	cycle, err := sp.cycle.Cycle()
	if err != nil {
		return nil, err
	}
	return cycle, nil
}

func (sp *BellmanFordSP) relax(g *EdgeWeightedDigraph, v int) {
	adjEdges := g.Adj(v)
	for _, edge := range adjEdges {
		w := edge.To()
		// found a shorter path to w from s
		if sp.distTo[w] > sp.distTo[v]+edge.Weight() {
			sp.distTo[w] = sp.distTo[v] + edge.Weight()
			sp.edgeTo[w] = edge
			if !sp.onQueue[w] {
				sp.queue.Enqueue(w)
				sp.onQueue[w] = true
			}
		}
		sp.cost++
		if sp.cost%g.V() == 0 {
			sp.findNegativeCycle()
		}

	}
}

func (sp *BellmanFordSP) DistTo(v int) float64 {
	return sp.distTo[v]
}

func (sp *BellmanFordSP) HasPathTo(v int) bool {
	return sp.distTo[v] < math.Inf(1)
}

func (sp *BellmanFordSP) PathTo(v int) []*DirectedEdge {
	stack := []*DirectedEdge{}
	curEdge := sp.edgeTo[v]
	for curEdge != nil {
		stack = append(stack, curEdge)
		curEdge = sp.edgeTo[curEdge.From()]
	}
	return stack
}
