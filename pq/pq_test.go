package pq

import "testing"

func TestPQ(t *testing.T) {

	pq := NewMaxPQ[int](10)
	pq.Insert(1)
	pq.Insert(10)
	pq.Insert(3)
	pq.Insert(2)
	pq.Insert(17)


    got := pq.String()
		want := ""
    if got != "" {
        t.Errorf("Abs(-1) = %s; want %s", got, want)
    }
}