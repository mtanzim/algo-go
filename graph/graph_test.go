package graph

import (
	"reflect"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

type edgeType = float64

func TestGraph_NewUndirectedFromScanner(t *testing.T) {
	g, err := GraphFromFile("./fixtures/graph.txt")
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, g.String())
}

func TestGraph_NewDigraphFromScanner(t *testing.T) {
	g, err := DigraphFromFile("./fixtures/graph.txt")
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, g.String())
}

func TestGraph_New(t *testing.T) {
	type args struct {
		v int
	}

	makeMaps := func(v int) []map[int]edgeType {
		adj := make([]map[int]edgeType, v)
		for i := 0; i < v; i++ {
			adj[i] = make(map[int]edgeType)
		}
		return adj
	}
	tests := []struct {
		name    string
		args    args
		want    *Graph[edgeType]
		wantErr bool
	}{
		{"1 vertex", args{1}, &Graph[edgeType]{v: 1, e: 0, adj: makeMaps(1)}, false},
		{"3 vertices", args{3}, &Graph[edgeType]{v: 3, e: 0, adj: makeMaps(3)}, false},
		{"0 vertices", args{0}, &Graph[edgeType]{v: 0, e: 0, adj: makeMaps(0)}, false},
		{"negative vertices", args{-4}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New[edgeType](tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {
	type edgesArg struct {
		from   int
		to     int
		weight edgeType
	}
	type args struct {
		v     int
		edges []edgesArg
	}
	tests := []struct {
		name    string
		args    args
		want    *Graph[edgeType]
		wantErr bool
	}{
		{
			name:    "3 vertices, add 1 edge",
			args:    args{3, []edgesArg{{from: 0, to: 1, weight: 3.45}}},
			want:    &Graph[edgeType]{v: 3, e: 1, adj: []map[int]edgeType{map[int]edgeType{1: 3.45}, map[int]edgeType{}, map[int]edgeType{}}},
			wantErr: false,
		},
		{
			name:    "3 vertices, add 1 edge in reverse",
			args:    args{3, []edgesArg{{from: 1, to: 0, weight: 3.46}}},
			want:    &Graph[edgeType]{v: 3, e: 1, adj: []map[int]edgeType{map[int]edgeType{}, map[int]edgeType{0: 3.46}, map[int]edgeType{}}},
			wantErr: false,
		},
		{
			name: "3 vertices, add all possible edges",
			args: args{3, []edgesArg{
				{from: 0, to: 1, weight: 44.25},
				{from: 0, to: 2, weight: 99.75},
				{from: 1, to: 0, weight: 3.46},
				{from: 1, to: 2, weight: 55.66},
				{from: 2, to: 1, weight: 23.34},
				{from: 2, to: 0, weight: 65.12},
			}},
			want:    &Graph[edgeType]{v: 3, e: 6, adj: []map[int]edgeType{map[int]edgeType{1: 44.25, 2: 99.75}, map[int]edgeType{0: 3.46, 2: 55.66}, map[int]edgeType{1: 23.34, 0: 65.12}}},
			wantErr: false,
		},
		{
			name:    "invalid vertex in edges",
			args:    args{3, []edgesArg{{from: 45, to: 0, weight: 3.46}}},
			want:    &Graph[edgeType]{v: 3, e: 0, adj: []map[int]edgeType{map[int]edgeType{}, map[int]edgeType{}, map[int]edgeType{}}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		g, err := New[edgeType](tt.args.v)
		for _, edge := range tt.args.edges {
			err = g.AddEdge(edge.from, edge.to, edge.weight)
		}
		got := g
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.AddEdge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.AddEgde() = %v, want %v", got, tt.want)
			}
		})
	}
}
