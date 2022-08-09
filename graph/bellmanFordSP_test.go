package graph

import (
	"math"
	"reflect"
	"testing"
)

func TestNewBellmanFordSP(t *testing.T) {
	gTiny, err := WeightedDigraphFromFile("./fixtures/tinyEWDnc.txt")
	gTinyDistToWant := map[int]float64{0: math.Inf(1), 1: 0.32, 2: math.Inf(1), 3: 0.10, 4: -0.66, 5: -0.31, 6: math.Inf(1), 7: -0.29}
	gTinyCycleWant := []int{5, 4, 5}
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
		{name: "tiny", args: args{g: gTiny, s: 5}, wantCycle: gTinyCycleWant, wantDisto: gTinyDistToWant, wantErr: false},
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
