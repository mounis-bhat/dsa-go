package avltree_test

import (
	"reflect"
	"testing"

	avltree "github.com/mounis-bhat/dsa-go/internal/avltree"
)

func TestAVLTree(t *testing.T) {
	avl := avltree.NewAVLTree()

	checkTree := func(want []int, got []int, message string) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v, want %v", message, got, want)
		}
	}

	// Test Insert
	avl.Insert(2)
	avl.Insert(1)
	avl.Insert(3)

	// Test InOrder
	got := avl.InOrder()
	want := []int{1, 2, 3}
	checkTree(want, got, "InOrder")

	// Test PreOrder
	got = avl.PreOrder()
	want = []int{2, 1, 3}
	checkTree(want, got, "PreOrder")

	// Test PostOrder
	got = avl.PostOrder()
	want = []int{1, 3, 2}
	checkTree(want, got, "PostOrder")

	// Test LevelOrder
	got = avl.LevelOrder()
	want = []int{2, 1, 3}
	checkTree(want, got, "LevelOrder")

}
