package selectionsort_test

import (
	"reflect"
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/selectionsort"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{4, 5, 4, 3, 4, 1}, expected: []int{1, 3, 4, 4, 4, 5}},
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{3, 3, 3, 3, 3}, expected: []int{3, 3, 3, 3, 3}},
	}

	for _, test := range tests {
		arr := make([]int, len(test.input))
		copy(arr, test.input)

		selectionsort.SelectionSort(&arr)
		if !reflect.DeepEqual(arr, test.expected) {
			t.Errorf("selectionSort(%v) = %v; expected %v", test.input, arr, test.expected)
		}
	}
}
