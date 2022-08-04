package graph

import (
	"testing"
)

func TestNewLazyPrimMST(t *testing.T) {
	gTiny, err := WeightedGraphFromFile("./fixtures/tinyEWG.txt")
	if err != nil {
		t.Error(err)
		return
	}
	gMedium, err := WeightedGraphFromFile("./fixtures/mediumEWG.txt")
	if err != nil {
		t.Error(err)
		return
	}
	epsilon := 0.00001
	if err != nil {
		t.Error(err)
	}
	type args struct {
		g *EdgeWeightedGraph
	}
	tests := []struct {
		name       string
		args       args
		wantWeight float64
		wantErr    bool
	}{
		{name: "tiny", args: args{g: gTiny}, wantWeight: 1.81, wantErr: false},
		{name: "medium", args: args{g: gMedium}, wantWeight: 10.46351, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLazyPrimMST(tt.args.g)
			if err != nil && tt.wantErr == false {
				t.Errorf("NewLazyPrimMST() Error= %v, wantErr %v", err, tt.wantErr)
				return
			}
			isCorrect := tt.wantWeight-epsilon < got.Weight() && tt.wantWeight+epsilon > got.Weight()
			if !isCorrect {
				t.Errorf("NewLazyPrimMST() = %v, want %v", got.Weight(), tt.wantWeight)
			}

		})
	}
}
