package sorting

import "sort"

func QuickSort(a sort.Interface) sort.Interface {
	quickSort(a, 0, a.Len()-1)
	return a
}

func quickSort(a sort.Interface, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	quickSort(a, lo, j-1)
	quickSort(a, j+1, hi)
}

func partition(a sort.Interface, lo, hi int) int {
	i := lo
	j := hi + 1

	for {
		for {
			i++
			if a.Less(i, lo) {
				break
			}
			if i == hi {
				break
			}
		}

		for {
			j--
			if a.Less(lo, j) {
				break
			}
			if j == lo {
				break
			}
		}
		a.Swap(lo, j)
		return j

	}

}
