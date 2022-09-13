package sorting

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		a byAge
	}

	rand.Seed(time.Now().UnixNano())
	people := makePeople(500)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "integers",
			args: args{a: people},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertionSort(tt.args.a)

			b := make(byAge, len(tt.args.a))
			copy(b, tt.args.a)
			sort.SliceStable(b, func(i, j int) bool {
				return b[i].Age < b[j].Age
			})

			if !reflect.DeepEqual(got, b) {
				t.Errorf("want: %v, got: %v", b, got)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	sizes := []int{5, 50, 500, 5_000}
	for _, v := range sizes {
		people := makePeople(v)
		b.Run(fmt.Sprintf("size_%d", v), func(b *testing.B) {
			InsertionSort(people)
		})
	}
}
