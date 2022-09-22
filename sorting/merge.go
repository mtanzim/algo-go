package sorting

import (
	"golang.org/x/exp/constraints"
)

func merge[T constraints.Ordered](a []T, lo, mid, hi int) {
	aux := make([]T, hi-lo)
	for k := lo; k <= hi; k++ {
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
