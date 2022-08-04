package graph

import (
	"errors"
	"fmt"
)

type WeightedEdge struct {
	v      int
	w      int
	weight float64
}

func NewEdge(v int, w int, weight float64) *WeightedEdge {
	return &WeightedEdge{v: v, w: w, weight: weight}
}

func (e *WeightedEdge) Weight() float64 {
	return e.weight
}

func (e *WeightedEdge) Either() int {
	return e.v
}

func (e *WeightedEdge) Other(vertex int) (int, error) {
	if e.v == vertex {
		return e.w, nil

	}
	if e.w == vertex {
		return e.v, nil
	}

	return -1, errors.New("invalid vertex sent")
}

func (e *WeightedEdge) String() string {
	return fmt.Sprintf("%d -> %d %f", e.v, e.w, e.weight)
}
