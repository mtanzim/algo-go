package graph

import (
	"fmt"
)

type DirectedEdge struct {
	v      int
	w      int
	weight float64
}

func NewDirectedEdge(v, w int, weight float64) *DirectedEdge {
	return &DirectedEdge{v, w, weight}
}

func (e *DirectedEdge) Weight() float64 {
	return e.weight
}

func (e *DirectedEdge) From() int {
	return e.v
}

func (e *DirectedEdge) To() int {
	return e.w
}

func (e *DirectedEdge) String() string {
	return fmt.Sprintf("%d -> %d %.2f", e.From(), e.To(), e.Weight())
}
