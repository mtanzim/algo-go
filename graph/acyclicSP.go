package graph

import (
	"math"
)

type AcyclicSP struct {
	edgeTo []*DirectedEdge
	distTo []float64
}

func NewAcyclicLP(g *EdgeWeightedDigraph, s int) (*AcyclicSP, error) {
	gPrime, err := NewEdgeWeightedDigraph(g.V())
	if err != nil {
		return nil, err
	}
	for v := 0; v < g.V(); v++ {
		for _, edge := range g.Adj(v) {
			edgePrime := NewDirectedEdge(edge.From(), edge.To(), -1*edge.Weight())
			gPrime.AddEdge(edgePrime)
		}
	}
	sp, err := NewAcyclicSP(gPrime, s)
	if err != nil {
		return nil, err
	}
	edgeToPrime := make([]*DirectedEdge, len(sp.edgeTo))
	distToPrime := make([]float64, len(sp.distTo))

	for i, edge := range sp.edgeTo {
		if edge == nil {
			continue
		}
		edgeToPrime[i] = NewDirectedEdge(edge.From(), edge.To(), -1*edge.Weight())
		distToPrime[i] = -1 * sp.distTo[i]
	}

	return &AcyclicSP{edgeTo: edgeToPrime, distTo: distToPrime}, nil

}

func NewAcyclicSP(g *EdgeWeightedDigraph, s int) (*AcyclicSP, error) {
	edgeTo := make([]*DirectedEdge, g.g.Vertices())
	distTo := make([]float64, g.g.Vertices())
	for i := 0; i < g.g.Vertices(); i++ {
		distTo[i] = math.Inf(1)
	}
	distTo[s] = 0.0
	sp := &AcyclicSP{edgeTo: edgeTo, distTo: distTo}
	topoOrder, err := GetTopologicalSort[*DirectedEdge](g.g)
	if err != nil {
		return nil, err
	}
	for !topoOrder.IsEmpty() {
		w, err := topoOrder.Pop()
		if err != nil {
			return nil, err
		}
		sp.relax(g, w)

	}
	return sp, nil
}

func (sp *AcyclicSP) relax(g *EdgeWeightedDigraph, v int) {
	adjEdges := g.Adj(v)
	for _, edge := range adjEdges {
		w := edge.To()
		// found a shorter path to w from s
		if sp.distTo[w] > sp.distTo[v]+edge.Weight() {
			sp.distTo[w] = sp.distTo[v] + edge.Weight()
			sp.edgeTo[w] = edge
		}
	}
}

func (sp *AcyclicSP) DistTo(v int) float64 {
	return sp.distTo[v]
}

func (sp *AcyclicSP) HasPathTo(v int) bool {
	return sp.distTo[v] < math.Inf(1)
}

func (sp *AcyclicSP) PathTo(v int) []*DirectedEdge {
	stack := []*DirectedEdge{}
	curEdge := sp.edgeTo[v]
	for curEdge != nil {
		stack = append(stack, curEdge)
		curEdge = sp.edgeTo[curEdge.From()]
	}
	return stack
}
