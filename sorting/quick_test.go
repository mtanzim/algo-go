package sorting

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		size int
	}{{1}, {45}, {90000}}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("size-%d", tt.size), func(t *testing.T) {
			arr, arrSorted := makeTestArrs(tt.size)
			arrInterface := sort.IntSlice(arr)
			want := sort.IntSlice(arrSorted)

			got := QuickSort(arrInterface)
			if !reflect.DeepEqual(want, got) {
				t.Errorf("failed, got: %v, want: %v", got, want)
			}
		})
	}
}
