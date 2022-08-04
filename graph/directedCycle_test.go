package graph

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestNewDirectedCycle(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{"cycle", "./fixtures/cycle.txt"},
		{"tinyDG", "./fixtures/tinyDG.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := DigraphFromFile(tt.filename)
			if err != nil {
				t.Error(err)
			}
			dc := NewDirectedCycle(g)
			cupaloy.SnapshotT(t, dc.String())
		})
	}

}
