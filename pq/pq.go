package pq

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type PQ[T constraints.Ordered] struct {
	keys    []T
	size    int
	maxSize int
}

func NewMaxPQ[T constraints.Ordered](maxSize int) *PQ[T] {
	keys := make([]T, maxSize+1)
	return &PQ[T]{keys: keys, size: 0, maxSize: maxSize}
}

func (pq *PQ[T]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PQ[T]) Size() int {
	return pq.size
}

func (pq *PQ[T]) Insert(key T) {
	pq.size++
	pq.keys[pq.size] = key
	pq.swim(pq.size)
}

func (pq *PQ[T]) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.exchange(k/2, k)
		k = k / 2
	}
}

func (pq *PQ[T]) less(i, j int) bool {
	return pq.keys[i] < pq.keys[j]
}

func (pq *PQ[T]) exchange(i, j int) {
	pq.keys[i], pq.keys[j] = pq.keys[j], pq.keys[i]
}

func (pq *PQ[T]) String() string {
	return fmt.Sprintf("%v", pq.keys)
}
