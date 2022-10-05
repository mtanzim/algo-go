package sorting

import (
	"math/rand"
	"sort"
	"time"
)

func QuickSort(a sort.Interface) sort.Interface {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(a.Len(), a.Swap)
	quickSort(a, 0, a.Len()-1)
	return a
}

func quickSort(a sort.Interface, lo, hi int) sort.Interface {
	if hi <= lo {
		return a
	}
	var p int
	a, p = partition(a, lo, hi)
	a = quickSort(a, lo, p-1)
	a = quickSort(a, p+1, hi)
	return a
}
func partition(a sort.Interface, low, high int) (sort.Interface, int) {
	i := low
	pivot := high
	for j := low; j < high; j++ {
		if a.Less(j, pivot) {
			a.Swap(i, j)
			i++
		}
	}
	a.Swap(i, high)
	return a, i
}

// func partition(a sort.Interface, lo, hi int) int {
// 	i := lo
// 	j := hi + 1

// 	for {
// 		for {
// 			i++
// 			if a.Less(lo, i) {
// 				break
// 			}

// 		}

// 		if i == hi {
// 			break
// 		}

// 		for {
// 			j--
// 			if a.Less(j, lo) {
// 				break
// 			}
// 		}
// 		if j == lo {
// 			break
// 		}
// 		if i >= j {
// 			break
// 		}
// 		a.Swap(i, j)
// 	}
// 	a.Swap(lo, j)

// 	return j

// }
