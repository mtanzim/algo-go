package linkedList

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLinkedList_AddToTail(t *testing.T) {

	ll := NewLinkedList(4)
	ll.AddToTail(5)
	ll.AddToTail(6)
	ll.AddToTail(7)

	want := &LinkedList{Data: 4, Next: &LinkedList{Data: 5, Next: &LinkedList{Data: 6, Next: &LinkedList{Data: 7, Next: nil}}}}

	if !reflect.DeepEqual(want, ll) {
		t.Errorf("LinkedList.AddToTail() = %v, want %v", ll, want)
	}
}

func TestLinkedList_RemoveTail(t *testing.T) {

	t1 := func(t *testing.T) {
		ll := NewLinkedList(4)
		ll.AddToTail(5)
		ll.AddToTail(6)
		ll.AddToTail(7)
		gotD, gotLL := ll.RemoveTail()

		want := &LinkedList{Data: 4, Next: &LinkedList{Data: 5, Next: &LinkedList{Data: 6, Next: nil}}}
		wantD := 7
		if !reflect.DeepEqual(want, ll) {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", gotLL, want)
		}
		if gotD != wantD {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", gotD, wantD)
		}
	}

	t2 := func(t *testing.T) {
		ll := NewLinkedList(4)
		ll.AddToTail(5)
		gotD, gotLL := ll.RemoveTail()

		want := &LinkedList{Data: 4, Next: nil}
		wantD := 5
		if !reflect.DeepEqual(want, ll) {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", gotLL, want)
		}
		if gotD != wantD {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", gotD, wantD)
		}
	}

	t3 := func(t *testing.T) {
		ll := NewLinkedList(4)
		gotD, gotLL := ll.RemoveTail()
		wantD := 4
		if gotLL != nil {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", gotLL, nil)
		}
		if gotD != wantD {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", gotD, wantD)
		}
	}

	tests := map[string]func(*testing.T){
		"remove tail from n > 2 nodes": t1,
		"remove tail from n = 2 nodes": t2,
		"remove tail from n = 1 nodes": t3,
	}

	for name, tt := range tests {
		t.Run(name, tt)
	}

}

func TestLinkedList_AddToHead(t *testing.T) {

	ll := NewLinkedList(4)
	ll.AddToTail(5)
	ll.AddToTail(6)
	newLL := ll.AddToHead(77)

	want := &LinkedList{Data: 77, Next: &LinkedList{Data: 4, Next: &LinkedList{Data: 5, Next: &LinkedList{Data: 6, Next: nil}}}}

	if !reflect.DeepEqual(want, newLL) {
		t.Errorf("LinkedList.AddToTail() = %v, want %v", newLL, want)
	}
}

func TestLinkedList_Remove(t *testing.T) {

	t1 := func(t *testing.T) {
		ll := NewLinkedList(4)
		ll.AddToTail(5)
		ll.AddToTail(6)
		ll.AddToTail(7)

		updatedLL := ll.Remove(5)

		want := &LinkedList{Data: 4, Next: &LinkedList{Data: 6, Next: &LinkedList{Data: 7, Next: nil}}}

		if !reflect.DeepEqual(want, updatedLL) {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", ll, want)
		}
	}

	t2 := func(t *testing.T) {
		ll := NewLinkedList(4)
		ll.AddToTail(5)
		ll.AddToTail(6)
		ll.AddToTail(7)

		updatedLL := ll.Remove(4)

		want := &LinkedList{Data: 5, Next: &LinkedList{Data: 6, Next: &LinkedList{Data: 7, Next: nil}}}

		if !reflect.DeepEqual(want, updatedLL) {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", ll, want)
		}
	}

	t3 := func(t *testing.T) {
		ll := NewLinkedList(4)
		ll.AddToTail(5)
		ll.AddToTail(6)
		ll.AddToTail(7)

		updatedLL := ll.Remove(7)

		want := &LinkedList{Data: 4, Next: &LinkedList{Data: 5, Next: &LinkedList{Data: 6, Next: nil}}}

		if !reflect.DeepEqual(want, updatedLL) {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", ll, want)
		}
	}

	t4 := func(t *testing.T) {
		ll := NewLinkedList(4)
		ll.AddToTail(5)
		ll.AddToTail(6)
		ll.AddToTail(7)

		updatedLL := ll.Remove(79)

		want := &LinkedList{Data: 4, Next: &LinkedList{Data: 5, Next: &LinkedList{Data: 6, Next: &LinkedList{Data: 7, Next: nil}}}}

		if !reflect.DeepEqual(want, updatedLL) {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", ll, want)
		}
	}

	t5 := func(t *testing.T) {
		ll := NewLinkedList(4)
		ll.AddToTail(5)
		ll.AddToTail(6)
		ll.AddToTail(7)

		updatedLL := ll.Remove(7)
		updatedLL = updatedLL.Remove(6)
		updatedLL = updatedLL.Remove(5)

		want := &LinkedList{Data: 4, Next: nil}

		if !reflect.DeepEqual(want, updatedLL) {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", ll, want)
		}
	}

	t6 := func(t *testing.T) {
		ll := NewLinkedList(4)

		updatedLL := ll.Remove(4)

		if updatedLL != nil {
			t.Errorf("LinkedList.AddToTail() = %v, want %v", ll, nil)
		}
	}

	tests := map[string]func(*testing.T){
		"remove from middle":   t1,
		"remove from head":     t2,
		"remove from tail":     t3,
		"remove non element":   t4,
		"remove multiple":      t5,
		"remove to make empty": t6,
	}

	for name, tt := range tests {
		t.Run(name, tt)
	}

}

func TestLinkedList_Tail(t *testing.T) {
	ll := NewLinkedList(4)
	ll.AddToTail(5)
	ll.AddToTail(6)
	ll.AddToTail(7)

	got := ll.Tail()
	want := 7

	if got != want {
		t.Errorf("LinkedList.Tail() = %v, want %v", got, want)
	}
}

func TestLinkedList_String(t *testing.T) {

	ll := NewLinkedList(4)
	ll.AddToTail(5)
	ll.AddToTail(6)
	ll.AddToTail(7)

	want := "4 -> 5 -> 6 -> 7"
	got := fmt.Sprint(ll)

	if want != got {
		t.Errorf("LinkedList.AddToTail() = `%v`, want `%v`", got, want)
	}
}
