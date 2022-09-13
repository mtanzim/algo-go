package sorting

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestSelectionSort(t *testing.T) {
	type args struct {
		a ByAge
	}

	rand.Seed(time.Now().UnixNano())
	size := 500
	people := make(ByAge, size)
	for i := range people {
		person := Person{}
		person.Name = randSeq(5)
		person.Age = rand.Intn(95)
		people[i] = person
	}

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
			got := SelectionSort(tt.args.a)

			b := make(ByAge, len(tt.args.a))
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

func BenchmarkSelectionSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	sizes := []int{5, 50, 500, 5_000}
	for _, v := range sizes {
		people := make(ByAge, v)
		for i := range people {
			person := Person{}
			person.Name = randSeq(5)
			person.Age = rand.Intn(95)
			people[i] = person
		}
		b.Run(fmt.Sprintf("size_%d", v), func(b *testing.B) {
			SelectionSort(people)
		})
	}
}
