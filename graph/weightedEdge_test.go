package graph

import (
	"reflect"
	"testing"

	priorityQueue "github.com/jupp0r/go-priority-queue"
)

func TestNewEdge(t *testing.T) {
	type args struct {
		v      int
		w      int
		weight float64
	}
	tests := []struct {
		name string
		args args
		want *WeightedEdge
	}{
		{name: "simple", args: args{v: 1, w: 2, weight: 33.3}, want: &WeightedEdge{1, 2, 33.3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEdge(tt.args.v, tt.args.w, tt.args.weight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPQwithWeightedEdge(t *testing.T) {
	a := NewEdge(0, 1, 33.3)
	b := NewEdge(0, 1, 13.3)
	c := NewEdge(0, 1, 19.3)
	d := NewEdge(0, 1, 99.3)

	edges := []*WeightedEdge{a, b, c, d}
	pq := priorityQueue.New()
	for _, edge := range edges {
		pq.Insert(edge, edge.Weight())
	}
	sorted := []*WeightedEdge{}
	for {
		curEdge, err := pq.Pop()
		if err != nil {
			break
		}
		sorted = append(sorted, curEdge.(*WeightedEdge))
	}

	expected := []*WeightedEdge{b, c, a, d}

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("pq did not sort correctly; expected %v, got: %v", expected, sorted)
	}

}
