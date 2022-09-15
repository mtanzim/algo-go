package sorting

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type sortFns = map[string]func(a sort.Interface) sort.Interface

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sizes := []int{5, 50, 500, 5_000}
	sortFns := sortFns{
		"insertion": InsertionSort,
		"selection": SelectionSort,
	}
	for fnName, fn := range sortFns {
		for _, v := range sizes {
			people := makePeople(v)
			b.Run(fmt.Sprintf("fn_%s_size_%d", fnName, v), func(b *testing.B) {
				fn(people)
			})
		}
	}
}
