package sorting

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "basic", args: args{a: []int{44, -33, 45}}, want: []int{-33, 44, 45}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.args.a)
			if !reflect.DeepEqual(tt.want, tt.args.a) {
				t.Errorf("failed, got: %v, want: %v", tt.args.a, tt.want)
			}
		})
	}
}
