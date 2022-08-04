package graph

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestGraph_WeightedGraphFromScanner(t *testing.T) {
	g, err := WeightedGraphFromFile("./fixtures/tinyEWG.txt")
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, g.String())
}
