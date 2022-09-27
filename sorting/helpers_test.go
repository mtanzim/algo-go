package sorting

import (
	"math/rand"
	"sort"
	"time"
)

type person struct {
	Name string
	Age  int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func makePeople(population int) byAge {
	rand.Seed(time.Now().UnixNano())
	size := population
	people := make(byAge, size)
	for i := range people {
		person := person{}
		person.Name = randSeq(5)
		person.Age = rand.Intn(95)
		people[i] = person
	}
	return people

}

func makeArr(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(95)
	}
	return arr

}

func makeTestArrs(size int) (arr, arrSorted []int) {
	arr = makeArr(size)
	arrSorted = make([]int, size)
	for i := range arrSorted {
		arrSorted[i] = arr[i]
	}
	sort.Ints(arrSorted)
	return
}
