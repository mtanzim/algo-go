package stack

import (
	"reflect"
	"testing"

	"github.com/mtanzim/algo-go/linkedList"
)

func TestNewStack(t *testing.T) {

	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want *Stack
	}{
		{"push", args{44}, &Stack{linkedList.NewLinkedList(44)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {

	s := NewStack(4)
	s.Push(33)
	s.Push(36)

	want := "36 -> 33 -> 4"

	t.Run("push", func(t *testing.T) {
		got := s.String()
		if got != want {
			t.Errorf("Stack.Push() = %v, want %v", got, want)
		}
	})
}

func TestStack_Pop(t *testing.T) {
	s := NewStack(4)
	s.Push(33)
	s.Push(36)
	gotV, _ := s.Pop()

	want := "33 -> 4"
	wantV := 36

	t.Run("push", func(t *testing.T) {
		got := s.String()
		if got != want {
			t.Errorf("Stack.Pop() resulting stack = %v, want %v", got, want)
		}
		if gotV != wantV {
			t.Errorf("Stack.Pop() = %v, want %v", got, want)
		}
	})
}

func TestStack_Peek(t *testing.T) {
	s := NewStack(4)
	s.Push(33)
	s.Push(36)

	wantV := 36

	t.Run("peek", func(t *testing.T) {
		got := s.String()
		gotD, _ := s.Peek()
		if gotD != wantV {
			t.Errorf("Stack.Peek() = %v, want %v", got, wantV)
		}
	})
}

func TestStack_IsEmpty(t *testing.T) {
	s := NewStack(4)
	s.Push(33)
	s.Pop()
	s.Pop()

	t.Run("peek", func(t *testing.T) {
		got := s.String()
		if s.IsEmpty() != true {
			t.Errorf("Stack.IsEmpty() = %v, want %v", got, true)
		}
	})
}
