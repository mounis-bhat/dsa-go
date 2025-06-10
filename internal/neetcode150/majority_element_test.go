package neetcode150_test

import (
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1: [3,2,3]",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "Example 2: [2,2,1,1,1,2,2]",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "Single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "All same elements",
			nums: []int{5, 5, 5, 5, 5},
			want: 5,
		},
		{
			name: "Majority at beginning",
			nums: []int{1, 1, 1, 2, 3},
			want: 1,
		},
		{
			name: "Majority at end",
			nums: []int{1, 2, 3, 3, 3},
			want: 3,
		},
		{
			name: "Majority scattered",
			nums: []int{6, 5, 5, 6, 6, 6, 7},
			want: 6,
		},
		{
			name: "Two elements - first majority",
			nums: []int{1, 2, 1},
			want: 1,
		},
		{
			name: "Two elements - second majority",
			nums: []int{1, 2, 2},
			want: 2,
		},
		{
			name: "Large array with clear majority",
			nums: []int{1, 2, 3, 2, 2, 2, 5, 4, 2},
			want: 2,
		},
		{
			name: "Negative numbers",
			nums: []int{-1, -1, -2, -1, -3},
			want: -1,
		},
		{
			name: "Mixed positive and negative",
			nums: []int{-1, 1, -1, 1, -1},
			want: -1,
		},
		{
			name: "Zero as majority",
			nums: []int{0, 1, 0, 2, 0},
			want: 0,
		},
		{
			name: "Large numbers",
			nums: []int{1000000, 999999, 1000000, 999998, 1000000},
			want: 1000000,
		},
		{
			name: "Exact majority (n/2 + 1)",
			nums: []int{1, 2, 1, 2, 1}, // 3 out of 5
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := neetcode150.MajorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("MajorityElement(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// TestMajorityElementEdgeCases tests edge cases for MajorityElement
func TestMajorityElementEdgeCases(t *testing.T) {
	// Test with large array where majority appears exactly n/2 + 1 times
	t.Run("Large array with minimal majority", func(t *testing.T) {
		// Create array of size 1001, majority element appears 501 times
		nums := make([]int, 1001)
		for i := 0; i < 501; i++ {
			nums[i] = 42 // majority element
		}
		for i := 501; i < 1001; i++ {
			nums[i] = i // different elements
		}
		want := 42

		got := neetcode150.MajorityElement(nums)
		if got != want {
			t.Errorf("MajorityElement with large minimal majority failed: got %v, want %v", got, want)
		}
	})

	// Test with alternating pattern where one clearly dominates
	t.Run("Alternating pattern with clear majority", func(t *testing.T) {
		nums := []int{1, 2, 1, 3, 1, 4, 1, 5, 1, 6, 1}
		want := 1

		got := neetcode150.MajorityElement(nums)
		if got != want {
			t.Errorf("MajorityElement with alternating pattern failed: got %v, want %v", got, want)
		}
	})

	// Test with maximum integer values
	t.Run("Maximum integer values", func(t *testing.T) {
		maxInt := 2147483647
		nums := []int{maxInt, maxInt - 1, maxInt, maxInt - 2, maxInt}
		want := maxInt

		got := neetcode150.MajorityElement(nums)
		if got != want {
			t.Errorf("MajorityElement with max values failed: got %v, want %v", got, want)
		}
	})

	// Test with minimum integer values
	t.Run("Minimum integer values", func(t *testing.T) {
		minInt := -2147483648
		nums := []int{minInt, minInt + 1, minInt, minInt + 2, minInt}
		want := minInt

		got := neetcode150.MajorityElement(nums)
		if got != want {
			t.Errorf("MajorityElement with min values failed: got %v, want %v", got, want)
		}
	})
}

// BenchmarkMajorityElement tests the performance of the MajorityElement function
func BenchmarkMajorityElement(b *testing.B) {
	// Create a large array with majority element
	nums := make([]int, 10000)
	for i := 0; i < 6000; i++ {
		nums[i] = 42 // majority element (60% of array)
	}
	for i := 6000; i < 10000; i++ {
		nums[i] = i // different elements
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		neetcode150.MajorityElement(nums)
	}
}

// BenchmarkMajorityElementWorstCase benchmarks worst case scenario
func BenchmarkMajorityElementWorstCase(b *testing.B) {
	// Worst case: majority element appears at the end
	nums := make([]int, 10000)
	for i := 0; i < 4999; i++ {
		nums[i] = i // different elements
	}
	for i := 4999; i < 10000; i++ {
		nums[i] = 99999 // majority element appears in second half
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		neetcode150.MajorityElement(nums)
	}
}

// BenchmarkMajorityElementBestCase benchmarks best case scenario
func BenchmarkMajorityElementBestCase(b *testing.B) {
	// Best case: all elements are the same
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = 42
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		neetcode150.MajorityElement(nums)
	}
}
