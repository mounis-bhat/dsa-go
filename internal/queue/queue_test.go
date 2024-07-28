package queue_test

import (
	"reflect"
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/queue"
)

func TestQueue(t *testing.T) {
	q := queue.NewQueue()

	// Helper function to check if the queue content matches expected
	checkQueue := func(want []int, message string) {
		if len(want) == 0 && q.IsEmpty() {
			return
		}
		if !reflect.DeepEqual(q.Display(), want) {
			t.Errorf("%s: got %v, want %v", message, q.Display(), want)
		}
	}

	// Test Enqueue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	checkQueue([]int{1, 2, 3}, "After Enqueue")

	// Test Dequeue
	if element, _ := q.Dequeue(); element != 1 {
		t.Errorf("Dequeue(): got %v, want %v", element, 1)
	}
	checkQueue([]int{2, 3}, "After Dequeue")

	// Test Peek
	if element, _ := q.Peek(); element != 2 {
		t.Errorf("Peek(): got %v, want %v", element, 2)
	}

	// Test IsEmpty
	if q.IsEmpty() {
		t.Error("IsEmpty(): expected false, got true")
	}
}
