package pq

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type PQ[T constraints.Ordered] struct {
	keys    []*T
	size    int
	maxSize int
	lessFn  func(keys []*T, i, j int) bool
}

func maxPQless[T constraints.Ordered](keys []*T, i, j int) bool {
	return *keys[i] < *keys[j]
}

func minPQless[T constraints.Ordered](keys []*T, i, j int) bool {
	return *keys[i] > *keys[j]
}

func NewMaxPQ[T constraints.Ordered](maxSize int) *PQ[T] {
	keys := make([]*T, maxSize+1)
	return &PQ[T]{keys: keys, size: 0, maxSize: maxSize, lessFn: maxPQless[T]}
}

func NewMinPQ[T constraints.Ordered](maxSize int) *PQ[T] {
	keys := make([]*T, maxSize+1)
	return &PQ[T]{keys: keys, size: 0, maxSize: maxSize, lessFn: minPQless[T]}
}

func (pq *PQ[T]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PQ[T]) Size() int {
	return pq.size
}

func (pq *PQ[T]) Insert(key T) {
	pq.size++
	pq.keys[pq.size] = &key
	pq.swim(pq.size)
}

func (pq *PQ[T]) swim(k int) {
	for k > 1 && pq.lessFn(pq.keys, k/2, k) {
		pq.exchange(k/2, k)
		k = k / 2
	}
}

func (pq *PQ[T]) DelTop() (*T, error) {
	if pq.IsEmpty() {
		return nil, errors.New("pq is empty")
	}
	top := *pq.keys[1]
	pq.exchange(1, pq.size)
	pq.size--
	pq.keys[pq.size+1] = nil
	pq.sink(1)
	return &top, nil
}

func (pq *PQ[T]) sink(k int) {
	for 2*k <= pq.size {
		leftChild := 2 * k
		rightChild := leftChild + 1
		chosenChild := leftChild
		if leftChild < pq.size && pq.lessFn(pq.keys, leftChild, rightChild) {
			chosenChild = rightChild
		}
		if !pq.lessFn(pq.keys, k, chosenChild) {
			break
		}
		pq.exchange(k, chosenChild)
		k = chosenChild
	}

}

func (pq *PQ[T]) exchange(i, j int) {
	pq.keys[i], pq.keys[j] = pq.keys[j], pq.keys[i]
}

func (pq *PQ[T]) String() string {
	var sb strings.Builder
	for _, v := range pq.keys {
		if v != nil {
			sb.WriteString(fmt.Sprintf("%v ->", *v))
		}
	}
	return sb.String()
}
