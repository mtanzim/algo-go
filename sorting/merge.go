package sorting

import (
	"golang.org/x/exp/constraints"
)

func Merge[T constraints.Ordered](a []T) {

	aux := make([]T, len(a))
	var sort func(a []T, lo, hi int)
	sort = func(a []T, lo, hi int) {

		if hi <= lo {
			return
		}
		mid := lo + (hi-lo)/2

		sort(a, lo, mid)
		sort(a, mid+1, hi)
		merge(a, aux, lo, mid, hi)
	}
	sort(a, 0, len(a)-1)
}

func merge[T constraints.Ordered](a, aux []T, lo, mid, hi int) {

	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = a[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}

	}
}
