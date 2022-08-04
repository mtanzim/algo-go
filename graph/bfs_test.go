package graph

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

// TODO: this is not deterministic
// Update the snapshot tests to something better
func TestNewBFS(t *testing.T) {

	tests := []struct {
		name       string
		filename   string
		source     int
		isDirected bool
	}{
		{"basic", "./fixtures/graph.txt", 0, false},
		{"tiny", "./fixtures/tinyCG.txt", 3, false},
		{"tiny start 0", "./fixtures/tinyCG.txt", 0, false},
		{"tiny no connections", "./fixtures/tinyCG.txt", 11, false},
		{"directed", "./fixtures/tinyDG.txt", 6, true},
		{"directed no connections", "./fixtures/tinyDG.txt", 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var g *Graph[UnweightedEdgeType]
			var err error
			if tt.isDirected {
				g, err = DigraphFromFile(tt.filename)
			} else {
				g, err = GraphFromFile(tt.filename)
			}
			if err != nil {
				t.Error(err)
			}
			bfs := NewBFS(g, tt.source)
			cupaloy.SnapshotT(t, bfs.String())
		})
	}

}
