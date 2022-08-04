package stack

import (
	"errors"
	"fmt"

	"github.com/mtanzim/algo-go/linkedList"
)

type Stack struct {
	ll *linkedList.LinkedList
}

// TODO: argh this is a bad api; should allow empty stack on construction
func NewStack(d int) *Stack {
	return &Stack{linkedList.NewLinkedList(d)}
}

func (s *Stack) Push(d int) {
	s.ll = s.ll.AddToHead(d)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty, cannot pop")
	}
	var d int
	d, s.ll = s.ll.RemoveHead()
	return d, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty, cannot peek")
	}
	return s.ll.Head(), nil
}

func (s *Stack) IsEmpty() bool {
	return s.ll.IsEmpty()
}

func (s *Stack) String() string {
	return fmt.Sprintf("%s", s.ll)
}
