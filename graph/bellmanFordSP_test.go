package graph

import (
	"math"
	"reflect"
	"testing"
)

func TestNewBellmanFordSP(t *testing.T) {
	gCyclic, err := WeightedDigraphFromFile("./fixtures/tinyEWDnc.txt")
	if err != nil {
		t.Error("cannot open ./fixtures/tinyEWDnc.txt")
		return
	}
	gCyclicDistToWant := map[int]float64{0: math.Inf(1), 1: 0.32, 2: math.Inf(1), 3: 0.10, 4: -0.66, 5: -0.31, 6: math.Inf(1), 7: -0.29}
	gCyclicCycleWant := []int{5, 4, 5}

	gAcyclic, err := WeightedDigraphFromFile("./fixtures/tinyEWDn.txt")
	if err != nil {
		t.Error("cannot open ./fixtures/tinyEWDn.txt")
		return
	}
	gAcyclicDistToWant := map[int]float64{0: 0.0, 1: 0.93, 2: 0.26, 3: 0.99, 4: 0.26, 5: 0.61, 6: 1.51, 7: 0.60}
	gAcyclicCycleWant := []int{}

	epsilon := 0.00001

	if err != nil {
		t.Error(err)
		return
	}
	if err != nil {
		t.Error(err)
	}
	type args struct {
		g *EdgeWeightedDigraph
		s int
	}
	tests := []struct {
		name      string
		args      args
		wantCycle []int
		wantDisto map[int]float64
		wantErr   bool
	}{
		{name: "cyclic", args: args{g: gCyclic, s: 5}, wantCycle: gCyclicCycleWant, wantDisto: gCyclicDistToWant, wantErr: false},
		{name: "acyclic with path", args: args{g: gAcyclic, s: 0}, wantCycle: gAcyclicCycleWant, wantDisto: gAcyclicDistToWant, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBellmanFordSP(tt.args.g, tt.args.s)
			if err != nil && tt.wantErr == false {
				t.Errorf("NewBellmanFordSP() Error= %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotCycle, err := got.NegativeCycle()
			if err != nil && tt.wantErr == false {
				t.Errorf("NewBellmanFordSP().NegativeCycle() Error= %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCycle, tt.wantCycle) {
				t.Errorf("want cycle = %v, got = %v ", tt.wantCycle, gotCycle)
				return
			}

			gotMap := make(map[int]float64)
			for k := range tt.wantDisto {
				if got.HasPathTo(k) {
					gotMap[k] = got.DistTo(k)
				}
			}

			for k, v := range gotMap {
				want := tt.wantDisto[k]
				isCorrect := want-epsilon < v && want+epsilon > v
				if !isCorrect {
					t.Errorf("vertex = %d, want = %.2f, got = %.2f ", k, want, v)
				}
			}

		})
	}
}
