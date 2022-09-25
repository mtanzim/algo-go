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
		{name: "basic", args: args{a: []int{4, 3, 1, 2, 5}}, want: []int{1, 2, 3, 4, 5}},
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
