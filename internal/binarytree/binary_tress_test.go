package binarytree_test

import (
	"reflect"
	"testing"

	binarytree "github.com/mounis-bhat/dsa-go/internal/binarytree"
)

func TestBinaryTree(t *testing.T) {
	bt := binarytree.NewBinaryTree()

	checkTree := func(want []int, got []int, message string) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v, want %v", message, got, want)
		}
	}

	// Test Insert
	bt.Insert(2)
	bt.Insert(1)
	bt.Insert(3)

	// Test InOrder
	got := bt.InOrder()
	want := []int{1, 2, 3}
	checkTree(want, got, "InOrder")

	// Test PreOrder
	got = bt.PreOrder()
	want = []int{2, 1, 3}
	checkTree(want, got, "PreOrder")

	// Test PostOrder
	got = bt.PostOrder()
	want = []int{1, 3, 2}
	checkTree(want, got, "PostOrder")

	// Test LevelOrder
	got = bt.LevelOrder()
	want = []int{2, 1, 3}
	checkTree(want, got, "LevelOrder")

	// Test Search
	if !bt.Search(3) {
		t.Errorf("Search(3): got %v, want %v", false, true)
	}
}
