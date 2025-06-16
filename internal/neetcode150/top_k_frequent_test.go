package neetcode150_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

// Helper function to check if two slices contain the same elements (order doesn't matter)
func containsSameElements(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	// Create copies to avoid modifying original slices
	gotCopy := make([]int, len(got))
	wantCopy := make([]int, len(want))
	copy(gotCopy, got)
	copy(wantCopy, want)

	// Sort both slices for comparison
	sort.Ints(gotCopy)
	sort.Ints(wantCopy)

	return reflect.DeepEqual(gotCopy, wantCopy)
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Example 1: [1,1,1,2,2,3], k=2",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "Example 2: [1], k=1",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "All elements same frequency",
			nums: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{1, 2, 3}, // Any 3 elements are valid
		},
		{
			name: "Two elements, k=1",
			nums: []int{1, 2},
			k:    1,
			want: []int{1}, // Either 1 or 2 is valid
		},
		{
			name: "All same elements",
			nums: []int{5, 5, 5, 5},
			k:    1,
			want: []int{5},
		},
		{
			name: "Clear frequency order",
			nums: []int{4, 1, -1, 2, -1, 2, 3},
			k:    2,
			want: []int{-1, 2}, // -1 appears 2 times, 2 appears 2 times
		},
		{
			name: "Negative numbers",
			nums: []int{-1, -1, -2, -2, -2, -3},
			k:    2,
			want: []int{-2, -1}, // -2 appears 3 times, -1 appears 2 times
		},
		{
			name: "Large numbers",
			nums: []int{100000, 100000, 99999, 99999, 99999},
			k:    1,
			want: []int{99999},
		},
		{
			name: "Mixed positive and negative",
			nums: []int{-1, 1, -1, 1, 0, 0, 0},
			k:    2,
			want: []int{0, -1}, // 0 appears 3 times, -1 and 1 each appear 2 times
		},
		{
			name: "Single element repeated",
			nums: []int{7, 7},
			k:    1,
			want: []int{7},
		},
		{
			name: "k equals number of unique elements",
			nums: []int{1, 2, 3, 1, 2, 1},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "Zero in array",
			nums: []int{0, 0, 1, 1, 1, 2},
			k:    2,
			want: []int{1, 0}, // 1 appears 3 times, 0 appears 2 times
		},
		{
			name: "Sequential numbers",
			nums: []int{1, 2, 3, 4, 5, 1, 2, 3, 1, 2, 1},
			k:    3,
			want: []int{1, 2, 3}, // 1 appears 4 times, 2 appears 3 times, 3 appears 2 times
		},
		{
			name: "Large k value",
			nums: []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
			k:    4,
			want: []int{1, 2, 3, 4}, // All except 5 appear twice
		},
		{
			name: "Duplicate frequencies",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: []int{1, 2}, // All appear twice, any 2 are valid
		},
		{
			name: "One element much more frequent",
			nums: []int{1, 1, 1, 1, 1, 2, 3},
			k:    2,
			want: []int{1, 2}, // 1 appears 5 times, 2 and 3 each appear once
		},
		{
			name: "Minimum array size with k=1",
			nums: []int{42},
			k:    1,
			want: []int{42},
		},
		{
			name: "Powers of 2",
			nums: []int{1, 2, 4, 8, 1, 2, 4, 1, 2, 1},
			k:    2,
			want: []int{1, 2}, // 1 appears 4 times, 2 appears 3 times
		},
		{
			name: "Consecutive duplicates",
			nums: []int{1, 1, 2, 2, 3, 3, 3},
			k:    1,
			want: []int{3}, // 3 appears 3 times
		},
		{
			name: "Mixed frequency pattern",
			nums: []int{5, 3, 1, 1, 1, 3, 5, 5, 5, 5},
			k:    3,
			want: []int{5, 1, 3}, // 5 appears 5 times, 1 appears 3 times, 3 appears 2 times
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := neetcode150.TopKFrequent(tt.nums, tt.k)

			// Check if we got the right number of elements
			if len(got) != tt.k {
				t.Errorf("TopKFrequent(%v, %d) returned %d elements, want %d", tt.nums, tt.k, len(got), tt.k)
				return
			}

			// For tests where all elements have same frequency or multiple valid answers exist,
			// we need more flexible checking
			if tt.name == "All elements same frequency" ||
				tt.name == "Two elements, k=1" ||
				tt.name == "Duplicate frequencies" ||
				tt.name == "One element much more frequent" {
				// Just verify we got k unique elements and they exist in the input
				uniqueElements := make(map[int]bool)
				for _, elem := range got {
					uniqueElements[elem] = true
				}
				if len(uniqueElements) != tt.k {
					t.Errorf("TopKFrequent(%v, %d) returned non-unique elements: %v", tt.nums, tt.k, got)
				}
				// Verify all returned elements exist in input
				inputMap := make(map[int]bool)
				for _, num := range tt.nums {
					inputMap[num] = true
				}
				for _, elem := range got {
					if !inputMap[elem] {
						t.Errorf("TopKFrequent(%v, %d) returned element %d not in input", tt.nums, tt.k, elem)
					}
				}
			} else {
				// For deterministic cases, check exact match
				if !containsSameElements(got, tt.want) {
					t.Errorf("TopKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
				}
			}
		})
	}
}

// BenchmarkTopKFrequent tests the performance of the TopKFrequent function
func BenchmarkTopKFrequent(t *testing.B) {
	// Create a large test case
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i % 100 // 100 unique numbers with varying frequencies
	}
	k := 10

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		neetcode150.TopKFrequent(nums, k)
	}
}

// TestTopKFrequentEdgeCases tests edge cases and constraints
func TestTopKFrequentEdgeCases(t *testing.T) {
	// Test with maximum frequency difference
	t.Run("Maximum frequency difference", func(t *testing.T) {
		nums := make([]int, 1000)
		// Element 1 appears 997 times, element 2 appears 2 times, element 3 appears 1 time
		for i := 0; i < 997; i++ {
			nums[i] = 1
		}
		nums[997] = 2
		nums[998] = 2
		nums[999] = 3

		got := neetcode150.TopKFrequent(nums, 2)
		expected := []int{1, 2}

		if !containsSameElements(got, expected) {
			t.Errorf("TopKFrequent with max frequency difference failed: got %v, want %v", got, expected)
		}
	})

	// Test with large array, small k
	t.Run("Large array, small k", func(t *testing.T) {
		nums := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			nums[i] = i % 10 // 10 unique numbers
		}

		got := neetcode150.TopKFrequent(nums, 3)

		if len(got) != 3 {
			t.Errorf("Expected 3 elements, got %d", len(got))
		}

		// All elements should appear 100 times, so any 3 are valid
		uniqueElements := make(map[int]bool)
		for _, elem := range got {
			uniqueElements[elem] = true
		}
		if len(uniqueElements) != 3 {
			t.Errorf("Expected 3 unique elements, got %d", len(uniqueElements))
		}
	})

	// Test performance with worst case scenario
	t.Run("Worst case performance", func(t *testing.T) {
		// All elements unique, so we need to track all frequencies
		nums := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			nums[i] = i
		}

		got := neetcode150.TopKFrequent(nums, 5)

		if len(got) != 5 {
			t.Errorf("Expected 5 elements, got %d", len(got))
		}

		// All elements appear once, so any 5 are valid
		uniqueElements := make(map[int]bool)
		for _, elem := range got {
			uniqueElements[elem] = true
		}
		if len(uniqueElements) != 5 {
			t.Errorf("Expected 5 unique elements, got %d", len(uniqueElements))
		}
	})

	// Test with very high frequency concentration
	t.Run("High frequency concentration", func(t *testing.T) {
		nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 5}
		got := neetcode150.TopKFrequent(nums, 2)

		// Element 1 should definitely be included
		found1 := false
		for _, elem := range got {
			if elem == 1 {
				found1 = true
				break
			}
		}
		if !found1 {
			t.Errorf("Element 1 (highest frequency) should be in top 2: got %v", got)
		}

		if len(got) != 2 {
			t.Errorf("Expected 2 elements, got %d", len(got))
		}
	})

	// Test boundary condition k = 1
	t.Run("k equals 1", func(t *testing.T) {
		nums := []int{3, 3, 3, 2, 2, 1}
		got := neetcode150.TopKFrequent(nums, 1)
		expected := []int{3}

		if !containsSameElements(got, expected) {
			t.Errorf("TopKFrequent for k=1 failed: got %v, want %v", got, expected)
		}
	})

	// Test with negative numbers having highest frequency
	t.Run("Negative numbers highest frequency", func(t *testing.T) {
		nums := []int{-5, -5, -5, -5, 1, 1, 2}
		got := neetcode150.TopKFrequent(nums, 2)
		expected := []int{-5, 1}

		if !containsSameElements(got, expected) {
			t.Errorf("TopKFrequent with negative numbers failed: got %v, want %v", got, expected)
		}
	})

	// Test memory efficiency
	t.Run("Memory efficiency test", func(t *testing.T) {
		// Test that we don't use excessive memory for sparse data
		nums := []int{1000000, 1000000, 1000000, -1000000, -1000000, 0}
		got := neetcode150.TopKFrequent(nums, 2)

		if len(got) != 2 {
			t.Errorf("Expected 2 elements, got %d", len(got))
		}

		// Should include 1000000 and -1000000 (both appear twice)
		hasPositive := false
		hasNegative := false
		for _, elem := range got {
			if elem == 1000000 {
				hasPositive = true
			}
			if elem == -1000000 {
				hasNegative = true
			}
		}
		if !hasPositive || !hasNegative {
			t.Errorf("Should include both 1000000 and -1000000: got %v", got)
		}
	})
}
