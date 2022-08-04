package queue

import (
	"errors"
	"fmt"

	"github.com/mtanzim/algo-go/linkedList"
)

type Queue struct {
	ll *linkedList.LinkedList
}

// TODO: argh this is a bad api; should allow empty on construction
func NewQueue(d int) *Queue {
	return &Queue{linkedList.NewLinkedList(d)}
}

func NewQueueEmpty() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(d int) {

	if q.IsEmpty() {
		q.ll = linkedList.NewLinkedList(d)
		return
	}
	q.ll = q.ll.AddToTail(d)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue, cannot dequeue")
	}
	d, updatedLl := q.ll.RemoveHead()
	q.ll = updatedLl
	return d, nil
}

func (q *Queue) IsEmpty() bool {
	return q.ll.IsEmpty()
}

func (q *Queue) String() string {
	return fmt.Sprint(q.ll)
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue, cannot peek")
	}
	return q.ll.Head(), nil
}
