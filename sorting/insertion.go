package sorting

import "sort"

func InsertionSort(a sort.Interface) sort.Interface {
	n := a.Len()
	for i := 0; i < n; i++ {
		for j := i; j > 0 && a.Less(j, j-1); j-- {
			a.Swap(j, j-1)
		}
	}
	return a
}
