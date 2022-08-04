package graph

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestGraph_WeightedDigraphFromScanner(t *testing.T) {
	g, err := WeightedDigraphFromFile("./fixtures/tinyEWD.txt")
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, g.String())
}
