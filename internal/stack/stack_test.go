package stack_test

import (
	"reflect"
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/stack"
)

func TestStack(t *testing.T) {
	s := stack.NewStack()

	// Helper function to check if the stack content matches expected
	checkStack := func(want []int, message string) {
		if !reflect.DeepEqual(s.Display(), want) {
			t.Errorf("%s: got %v, want %v", message, s.Display(), want)
		}
	}

	// Test Push
	s.Push(1)
	s.Push(2)
	s.Push(3)
	checkStack([]int{1, 2, 3}, "After Push")

	// Test Pop
	if element, _ := s.Pop(); element != 3 {
		t.Errorf("Pop(): got %v, want %v", element, 3)
	}
	checkStack([]int{1, 2}, "After Pop")

	// Test IsEmpty
	if s.IsEmpty() {
		t.Error("IsEmpty(): expected false, got true")
	}
}
