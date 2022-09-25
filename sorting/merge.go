package sorting

import (
	"golang.org/x/exp/constraints"
)

func Merge[T constraints.Ordered](a []T) {
	var sort func(a []T, lo, hi int)
	sort = func(a []T, lo, hi int) {
		if hi <= lo {
			return
		}
		mid := lo + (hi-lo)/2
		sort(a, lo, mid)
		sort(a, mid+1, hi)
		merge(a, lo, mid, hi)
	}
	sort(a, 0, len(a)-1)
}

func merge[T constraints.Ordered](a []T, lo, mid, hi int) {
	aux := make([]T, hi-lo)
	for k := lo; k < hi; k++ {
		aux[k] = a[k]
	}
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			j++
			a[k] = aux[j]
		} else if j > hi {
			i++
			a[k] = aux[i]
		} else if aux[j] < aux[i] {
			j++
			a[k] = a[j]
		} else {
			i++
			a[k] = aux[i]
		}

	}
}
