package graph

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{"with cycle", "./fixtures/cycle.txt", true},
		{"with dag", "./fixtures/dag.txt", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := DigraphFromFile(tt.filename)
			if err != nil {
				t.Error(err)
			}
			res, err := GetTopologicalSort(g)
			if tt.wantErr && err != nil {
				cupaloy.SnapshotT(t, err.Error())
				return
			}
			cupaloy.SnapshotT(t, res.String())
		})
	}

}
