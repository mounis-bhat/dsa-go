package neetcode150_test

import (
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

func TestContainsDuplicateII(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "Example 1: [1,2,3,1], k=3",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "Example 2: [1,0,1,1], k=1",
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			name: "Example 3: [1,2,3,1,2,3], k=2",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
		{
			name: "Single element",
			nums: []int{1},
			k:    1,
			want: false,
		},
		{
			name: "Two identical adjacent elements, k=1",
			nums: []int{1, 1},
			k:    1,
			want: true,
		},
		{
			name: "Two identical elements, k=0",
			nums: []int{1, 1},
			k:    0,
			want: false,
		},
		{
			name: "No duplicates",
			nums: []int{1, 2, 3, 4, 5},
			k:    2,
			want: false,
		},
		{
			name: "All same elements, k=1",
			nums: []int{1, 1, 1, 1},
			k:    1,
			want: true,
		},
		{
			name: "All same elements, k=0",
			nums: []int{1, 1, 1, 1},
			k:    0,
			want: false,
		},
		{
			name: "Duplicates too far apart",
			nums: []int{1, 2, 3, 4, 1},
			k:    3,
			want: false,
		},
		{
			name: "Duplicates exactly k distance",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "Multiple pairs, one valid",
			nums: []int{1, 2, 1, 3, 2},
			k:    2,
			want: true,
		},
		{
			name: "Multiple pairs, none valid",
			nums: []int{1, 2, 3, 1, 2},
			k:    2,
			want: false,
		},
		{
			name: "Negative numbers",
			nums: []int{-1, -2, -3, -1},
			k:    3,
			want: true,
		},
		{
			name: "Mixed positive and negative",
			nums: []int{-1, 1, -1, 1},
			k:    1,
			want: false,
		},
		{
			name: "Zero values",
			nums: []int{0, 1, 0, 2},
			k:    2,
			want: true,
		},
		{
			name: "Large k value",
			nums: []int{1, 2, 3, 1},
			k:    100,
			want: true,
		},
		{
			name: "k equals array length",
			nums: []int{1, 2, 3, 1},
			k:    4,
			want: true,
		},
		{
			name: "First and last elements same, within k",
			nums: []int{5, 1, 2, 5},
			k:    4,
			want: true,
		},
		{
			name: "First and last elements same, outside k",
			nums: []int{5, 1, 2, 5},
			k:    2,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := neetcode150.ContainsDuplicateII(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("ContainsDuplicateII(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

// BenchmarkContainsDuplicateII tests the performance of the ContainsDuplicateII function
func BenchmarkContainsDuplicateII(t *testing.B) {
	// Large array with duplicates at the end
	nums := make([]int, 10000)
	for i := 0; i < 9999; i++ {
		nums[i] = i
	}
	nums[9999] = 5000 // Duplicate of nums[5000]
	k := 5000

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		neetcode150.ContainsDuplicateII(nums, k)
	}
}

// TestContainsDuplicateIIEdgeCases tests edge cases
func TestContainsDuplicateIIEdgeCases(t *testing.T) {
	// Test with minimum array size (exactly 2 elements)
	t.Run("Minimum size array - valid", func(t *testing.T) {
		nums := []int{1, 1}
		k := 1
		got := neetcode150.ContainsDuplicateII(nums, k)
		if !got {
			t.Errorf("ContainsDuplicateII(%v, %d) should return true for minimum valid case", nums, k)
		}
	})

	// Test with minimum array size - invalid
	t.Run("Minimum size array - invalid", func(t *testing.T) {
		nums := []int{1, 2}
		k := 1
		got := neetcode150.ContainsDuplicateII(nums, k)
		if got {
			t.Errorf("ContainsDuplicateII(%v, %d) should return false for no duplicates", nums, k)
		}
	})

	// Test with large k and small array
	t.Run("Large k, small array", func(t *testing.T) {
		nums := []int{1, 2, 1}
		k := 1000000
		got := neetcode150.ContainsDuplicateII(nums, k)
		if !got {
			t.Errorf("ContainsDuplicateII(%v, %d) should return true when k is very large", nums, k)
		}
	})

	// Test with all identical elements
	t.Run("All identical elements", func(t *testing.T) {
		nums := []int{7, 7, 7, 7, 7}
		k := 1
		got := neetcode150.ContainsDuplicateII(nums, k)
		if !got {
			t.Errorf("ContainsDuplicateII(%v, %d) should return true for all identical elements", nums, k)
		}
	})

	// Test boundary condition where difference equals k
	t.Run("Difference exactly equals k", func(t *testing.T) {
		nums := []int{1, 0, 0, 1}
		k := 3
		got := neetcode150.ContainsDuplicateII(nums, k)
		if !got {
			t.Errorf("ContainsDuplicateII(%v, %d) should return true when difference equals k", nums, k)
		}
	})

	// Test boundary condition where difference is k+1
	t.Run("Difference is k+1", func(t *testing.T) {
		nums := []int{1, 0, 0, 0, 1}
		k := 3
		got := neetcode150.ContainsDuplicateII(nums, k)
		if got {
			t.Errorf("ContainsDuplicateII(%v, %d) should return false when difference is k+1", nums, k)
		}
	})

	// Test with multiple duplicates of same number
	t.Run("Multiple duplicates, first valid", func(t *testing.T) {
		nums := []int{1, 2, 1, 2, 1}
		k := 2
		got := neetcode150.ContainsDuplicateII(nums, k)
		if !got {
			t.Errorf("ContainsDuplicateII(%v, %d) should return true for first valid duplicate pair", nums, k)
		}
	})

	// Test performance with worst case scenario
	t.Run("Worst case performance", func(t *testing.T) {
		// Create array where duplicate is at the very end
		nums := make([]int, 1000)
		for i := 0; i < 999; i++ {
			nums[i] = i
		}
		nums[999] = 500 // Duplicate of nums[500]
		k := 500

		got := neetcode150.ContainsDuplicateII(nums, k)
		if !got {
			t.Errorf("ContainsDuplicateII should handle large arrays efficiently")
		}
	})
}
