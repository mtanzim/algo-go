package pq

import (
	"reflect"
	"testing"
)

func TestPQ_Insert(t *testing.T) {

	pq := NewMaxPQ[int](10)
	pq.Insert(1)
	pq.Insert(10)
	pq.Insert(3)
	pq.Insert(2)
	pq.Insert(17)
	// pq.Insert(17)

	got := pq.String()
	want := "17 ->10 ->3 ->1 ->2 ->"
	if got != want {
		t.Errorf("want = %s; got %s", want, got)
	}
}

func TestPQ_DelMax(t *testing.T) {

	pq := NewMaxPQ[int](10)
	pq.Insert(1)
	pq.Insert(10)
	pq.Insert(3)
	pq.Insert(2)
	pq.Insert(17)

	var err error
	var v *int
	vals := []int{}
	for err == nil {
		v, err = pq.DelTop()
		if v != nil {
			vals = append(vals, *v)
		}
	}
	want := []int{17, 10, 3, 2, 1}
	if !reflect.DeepEqual(vals, want) {
		t.Errorf("want = %v; got %v", want, vals)
	}

}
