package graph

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

// TODO: this is not deterministic
// Update the snapshot tests to something better
func TestNewDirectedDFS(t *testing.T) {

	tests := []struct {
		name     string
		filename string
		sources  []int
	}{
		{"from 1", "./fixtures/tinyDG.txt", []int{1}},
		{"from 2", "./fixtures/tinyDG.txt", []int{2}},
		{"from 1,2,6", "./fixtures/tinyDG.txt", []int{1, 2, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := DigraphFromFile(tt.filename)
			if err != nil {
				t.Error(err)
			}
			ddfs := NewDirectedDFS(g, tt.sources)
			cupaloy.SnapshotT(t, ddfs.String())
		})
	}

}
