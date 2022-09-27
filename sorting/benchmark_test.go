package sorting

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type sortFns = map[string]func(a sort.Interface) sort.Interface

func mergeSortWrapped(a sort.Interface) sort.Interface {
	switch v := a.(type) {
	case sort.IntSlice:
		rv := MergeSort[int](v)
		return sort.IntSlice(rv)
	default:
		panic("invalid setup")
	}

}

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sizes := []int{5, 50, 500, 5_000}
	sortFns := sortFns{
		"insertion": InsertionSort,
		"selection": SelectionSort,
		"merge":     mergeSortWrapped,
	}
	for fnName, sortFn := range sortFns {
		for _, v := range sizes {
			arr := makeArr(v)
			arrPrime := sort.IntSlice(arr)
			b.Run(fmt.Sprintf("fn_%s_size_%d", fnName, v), func(b *testing.B) {
				sortFn(arrPrime)
			})
		}
	}
}
