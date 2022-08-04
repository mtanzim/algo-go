package graph

import (
	"testing"
)

func TestNewDjikstraDigrap(t *testing.T) {
	gTiny, err := WeightedDigraphFromFile("./fixtures/tinyEWD.txt")
	gTinyDistToWant := map[int]float64{0: 0.0, 1: 1.05, 2: 0.26, 3: 0.99, 4: 0.38, 5: 0.73, 6: 1.51, 7: 0.60}
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
		{name: "tiny", args: args{g: gTiny, s: 0}, wantDisto: gTinyDistToWant, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDjikstraSP(tt.args.g, tt.args.s)
			if err != nil && tt.wantErr == false {
				t.Errorf("NewDjikstraSP() Error= %v, wantErr %v", err, tt.wantErr)
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
