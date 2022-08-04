package graph

import (
	"testing"
)

func TestNewAcyclicSP(t *testing.T) {
	gTiny, err := WeightedDigraphFromFile("./fixtures/tinyEWDAG.txt")
	gTinyDistToWant := map[int]float64{0: 0.73, 1: 0.32, 2: 0.62, 3: 0.61, 4: 0.35, 5: 0.00, 6: 1.13, 7: 0.28}
	if err != nil {
		t.Error(err)
		return
	}
	epsilon := 0.00001
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
		wantDisto map[int]float64
		wantErr   bool
	}{
		{name: "tiny", args: args{g: gTiny, s: 5}, wantDisto: gTinyDistToWant, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAcyclicSP(tt.args.g, tt.args.s)
			if err != nil && tt.wantErr == false {
				t.Errorf("NewAcyclicSP() Error= %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotMap := make(map[int]float64)
			for k := range tt.wantDisto {
				gotMap[k] = got.DistTo(k)
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
