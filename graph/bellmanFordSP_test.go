package graph

import (
	"reflect"
	"testing"
)

func TestNewBellmanFordSP(t *testing.T) {
	gTiny, err := WeightedDigraphFromFile("./fixtures/tinyEWDnc.txt")
	gTinyCycleWant := []int{5, 4, 5}
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
		wantErr   bool
	}{
		{name: "tiny", args: args{g: gTiny, s: 5}, wantCycle: gTinyCycleWant, wantErr: false},
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
		})
	}
}
