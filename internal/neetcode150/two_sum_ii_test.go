package neetcode150_test

import (
	"reflect"
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "Example 1: [2,7,11,15], target 9",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "Example 2: [2,3,4], target 6",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "Minimum array size",
			numbers: []int{1, 2},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "Negative numbers",
			numbers: []int{-1, 0, 1, 2},
			target:  0,
			want:    []int{1, 3},
		},
		{
			name:    "All negative numbers",
			numbers: []int{-5, -3, -1},
			target:  -4,
			want:    []int{2, 3},
		},
		{
			name:    "Target is sum of first and last",
			numbers: []int{1, 2, 3, 4, 5},
			target:  6,
			want:    []int{1, 5},
		},
		{
			name:    "Target is sum of adjacent elements",
			numbers: []int{1, 2, 3, 4, 5},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "Duplicate numbers - valid pair",
			numbers: []int{1, 2, 2, 3},
			target:  5,
			want:    []int{2, 4},
		},
		{
			name:    "Large numbers",
			numbers: []int{1000, 2000, 3000, 4000},
			target:  7000,
			want:    []int{3, 4},
		},
		{
			name:    "Zero in array",
			numbers: []int{-1, 0, 1, 2},
			target:  3,
			want:    []int{3, 4},
		},
		{
			name:    "Target at beginning of array",
			numbers: []int{1, 2, 3, 4, 5},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "Target at end of array",
			numbers: []int{1, 2, 3, 4, 5},
			target:  9,
			want:    []int{4, 5},
		},
		{
			name:    "Middle elements sum to target",
			numbers: []int{1, 2, 3, 4, 5, 6, 7},
			target:  13,
			want:    []int{6, 7},
		},
		{
			name:    "Same difference from center",
			numbers: []int{2, 7, 11, 15},
			target:  22,
			want:    []int{2, 4},
		},
		{
			name:    "Consecutive numbers",
			numbers: []int{1, 2, 3, 4, 5},
			target:  8,
			want:    []int{3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := neetcode150.TwoSumII(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumII(%v, %d) = %v, want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}

// TestTwoSumIIEdgeCases tests edge cases for TwoSumII
func TestTwoSumIIEdgeCases(t *testing.T) {
	// Test with large array
	t.Run("Large array", func(t *testing.T) {
		numbers := make([]int, 1000)
		for i := range numbers {
			numbers[i] = i + 1
		}
		target := 1999 // sum of first and last element (1 + 1000 = 1001, so use 999 + 1000)
		want := []int{999, 1000}

		got := neetcode150.TwoSumII(numbers, target)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TwoSumII with large array failed: got %v, want %v", got, want)
		}
	})

	// Test with maximum values
	t.Run("Maximum integer values", func(t *testing.T) {
		numbers := []int{1000000000, 1000000001}
		target := 2000000001
		want := []int{1, 2}

		got := neetcode150.TwoSumII(numbers, target)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TwoSumII with max values failed: got %v, want %v", got, want)
		}
	})

	// Test with very small target
	t.Run("Very small target", func(t *testing.T) {
		numbers := []int{-1000, -500, -100, 100}
		target := -1100
		want := []int{1, 3}

		got := neetcode150.TwoSumII(numbers, target)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TwoSumII with small target failed: got %v, want %v", got, want)
		}
	})

	// Test with identical adjacent numbers
	t.Run("Identical adjacent numbers", func(t *testing.T) {
		numbers := []int{1, 3, 3, 6}
		target := 6
		want := []int{2, 3}

		got := neetcode150.TwoSumII(numbers, target)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TwoSumII with identical adjacent numbers failed: got %v, want %v", got, want)
		}
	})
}

// BenchmarkTwoSumII tests the performance of the TwoSumII function
func BenchmarkTwoSumII(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i + 1
	}
	target := 1999

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		neetcode150.TwoSumII(numbers, target)
	}
}

// BenchmarkTwoSumIIWorstCase benchmarks worst case scenario
func BenchmarkTwoSumIIWorstCase(b *testing.B) {
	// Worst case: target requires scanning most of the array
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i + 1
	}
	target := numbers[498] + numbers[499] // middle elements

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		neetcode150.TwoSumII(numbers, target)
	}
}

// BenchmarkTwoSumIIBestCase benchmarks best case scenario
func BenchmarkTwoSumIIBestCase(b *testing.B) {
	// Best case: target is sum of first and last elements
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i + 1
	}
	target := numbers[0] + numbers[999] // first + last

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		neetcode150.TwoSumII(numbers, target)
	}
}
