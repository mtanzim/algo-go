package graph

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

// TODO: tests are not deterministic
func TestNewDFS(t *testing.T) {

	tests := []struct {
		name     string
		filename string
		source   int
	}{
		{"basic", "./fixtures/graph.txt", 0},
		{"tiny", "./fixtures/tinyCG.txt", 3},
		{"tiny start 0", "./fixtures/tinyCG.txt", 0},
		{"tiny no connections", "./fixtures/tinyCG.txt", 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := GraphFromFile(tt.filename)
			if err != nil {
				t.Error(err)
			}
			dfs := NewDFS(g, tt.source)
			cupaloy.SnapshotT(t, dfs.String())
		})
	}

}
