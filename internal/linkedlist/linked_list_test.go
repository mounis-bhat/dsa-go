package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/mounis-bhat/dsa-go/internal/linkedlist"
)

func TestLinkedList(t *testing.T) {
	ll := linkedlist.NewLinkedList()

	// Helper function to check if the list content matches expected
	checkList := func(want []int, message string) {
		got := ll.Display()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v, want %v", message, got, want)
		}
	}

	// Test Append
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	checkList([]int{1, 2, 3}, "After Append")

	// Test Prepend
	ll.Prepend(0)
	checkList([]int{0, 1, 2, 3}, "After Prepend")

	// Test Insert
	ll.Insert(4, 4)
	checkList([]int{0, 1, 2, 3, 4}, "After Insert")

	// Test Delete using value
	ll.Delete(4)
	checkList([]int{0, 1, 2, 3}, "After Delete")

	// Test Get using index
	if element := ll.Get(2); element != 2 {
		t.Errorf("Get(2): got %v, want %v", element, 2)
	}

	// Test Search using value
	if element := ll.Search(3); element != 3 {
		t.Errorf("Search(3): got %v, want %v", element, 3)
	}

	// Test Size
	if size := ll.Size(); size != 4 {
		t.Errorf("Size(): got %v, want %v", size, 4)
	}

	// Test IsEmpty
	if ll.IsEmpty() {
		t.Error("IsEmpty(): expected false, got true")
	}

	// Test Reverse
	ll.Reverse()
	checkList([]int{3, 2, 1, 0}, "After Reverse")

	ll.Pop()
	checkList([]int{3, 2, 1}, "After Pop")

	// Test Clear
	ll.Clear()
	got := ll.Display()
	if len(got) != 0 {
		t.Errorf("After Clear: list is not empty, got %v", got)
	}

	// Test IsEmpty after clearing
	if !ll.IsEmpty() {
		t.Error("IsEmpty() after Clear: expected true, got false")
	}
}
