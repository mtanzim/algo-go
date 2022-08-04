package linkedList

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	Data int
	Next *LinkedList
}

func (ll *LinkedList) String() string {
	var sb strings.Builder
	sb.Write([]byte(fmt.Sprintf("%d", ll.Data)))
	for ptr := ll.Next; ptr != nil; ptr = ptr.Next {
		sb.Write([]byte(fmt.Sprintf(" -> %d", ptr.Data)))
	}
	return sb.String()
}

// TODO: argh this is a bad api; should allow empty on construction
func NewLinkedList(d int) *LinkedList {
	return &LinkedList{d, nil}
}

func (ll *LinkedList) AddToTail(d int) *LinkedList {
	if ll == nil {
		return NewLinkedList(d)
	}
	ptr := ll
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = NewLinkedList(d)
	return ll
}

func (ll *LinkedList) AddToHead(d int) *LinkedList {
	head := NewLinkedList(d)
	head.Next = ll
	return head
}

func (ll *LinkedList) Head() int {
	return ll.Data
}

func (ll *LinkedList) RemoveHead() (int, *LinkedList) {
	d := ll.Data
	updatedLL := ll.Next
	return d, updatedLL
}

// TODO: does Remove make sense? Will anything use it?
// Nothing gurantees that d is unique

// Remove needs to return a *LinkedList since there is a possiblity that the head will change
func (ll *LinkedList) Remove(d int) *LinkedList {

	// move head over
	if d == ll.Data {
		_, updatedLL := ll.RemoveHead()
		return updatedLL
	}

	var prev *LinkedList
	for ptr := ll; ptr != nil; ptr = ptr.Next {
		if ptr.Data == d {
			prev.Next = ptr.Next
			return ll
		}
		prev = ptr
	}
	return ll
}

func (ll *LinkedList) Len() int {
	var i int
	ptr := ll
	for ptr != nil {
		ptr = ptr.Next
		i++
	}
	return i
}

func (ll *LinkedList) Tail() int {
	var i int
	ptr := ll
	for ptr.Next != nil {
		ptr = ptr.Next
		i++
	}
	return ptr.Data
}

func (ll *LinkedList) RemoveTail() (int, *LinkedList) {

	if ll.Next == nil {
		return ll.RemoveHead()
	}

	var i int
	ptr := ll
	for ptr.Next.Next != nil {
		ptr = ptr.Next
		i++
	}
	d := ptr.Next.Data
	ptr.Next = nil
	return d, ll
}

func (ll *LinkedList) IsEmpty() bool {
	return ll == nil
}
