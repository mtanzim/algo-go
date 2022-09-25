package sorting

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func makeArr(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(95)
	}
	return arr

}

func makeTestArrs(size int) (arr, arrSorted []int) {
	arr = makeArr(size)
	arrSorted = make([]int, size)
	for i := range arrSorted {
		arrSorted[i] = arr[i]
	}
	sort.Ints(arrSorted)
	return
}

func TestMerge(t *testing.T) {

	tests := []struct {
		size int
	}{
		{1}, {45}, {9999}, {0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("size-%d", tt.size), func(t *testing.T) {
			arr, arrSorted := makeTestArrs(tt.size)
			Merge(arr)
			want := arrSorted
			if !reflect.DeepEqual(want, arr) {
				t.Errorf("failed, got: %v, want: %v", arr, want)
			}
		})
	}
}
