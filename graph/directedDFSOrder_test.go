package graph

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestNewDirectedOrder(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{"from 1,2,6", "./fixtures/cycle.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := DigraphFromFile(tt.filename)
			if err != nil {
				t.Error(err)
			}
			dfso := NewDepthFirstOrder(g)
			cupaloy.SnapshotT(t, dfso.String())
		})
	}

}
