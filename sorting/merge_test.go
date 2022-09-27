package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {

	tests := []struct {
		size int
	}{
		{1}, {45}, {9999}, {0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("size-%d", tt.size), func(t *testing.T) {
			arr, arrSorted := makeTestArrs(tt.size)
			MergeSort(arr)
			want := arrSorted
			if !reflect.DeepEqual(want, arr) {
				t.Errorf("failed, got: %v, want: %v", arr, want)
			}
		})
	}
}
