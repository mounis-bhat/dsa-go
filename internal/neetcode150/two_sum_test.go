package neetcode150_test

import (
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1: [2,7,11,15], target 9",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Example 2: [3,2,4], target 6",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Example 3: [3,3], target 6",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "Negative numbers: [-1,-2,-3,-4,-5], target -8",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "Mixed positive and negative: [-3,4,3,90], target 0",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
		{
			name:   "Large array with valid pair",
			nums:   []int{230, 863, 916, 585, 981, 404, 316, 785, 88, 12, 70, 435, 384, 778, 887, 755, 740, 337, 86, 92, 325, 422, 815, 650, 920, 125, 277, 336, 221, 847, 168, 23, 677, 61, 400, 136, 874, 363, 394, 199, 863},
			target: 634, // 230 + 404 = 634
			want:   []int{0, 5},
		},
		{
			name:   "Two elements at beginning: [1,2,3,4], target 3",
			nums:   []int{1, 2, 3, 4},
			target: 3,
			want:   []int{0, 1},
		},
		{
			name:   "Two elements at end: [1,2,3,4], target 7",
			nums:   []int{1, 2, 3, 4},
			target: 7,
			want:   []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := neetcode150.TwoSum(tt.nums, tt.target)

			// Check if the result is correct by verifying the sum
			if len(got) == 2 {
				if tt.nums[got[0]]+tt.nums[got[1]] != tt.target {
					t.Errorf("TwoSum() returned indices that don't sum to target. nums[%d] + nums[%d] = %d + %d = %d, want %d",
						got[0], got[1], tt.nums[got[0]], tt.nums[got[1]], tt.nums[got[0]]+tt.nums[got[1]], tt.target)
				}
			} else {
				t.Errorf("TwoSum() = %v, want valid indices that sum to %d", got, tt.target)
			}

			// For exact match testing (order might vary, so we check both possibilities)
			if len(got) == 2 && len(tt.want) == 2 {
				// Check if we got the expected indices (order might be different)
				if !((got[0] == tt.want[0] && got[1] == tt.want[1]) ||
					(got[0] == tt.want[1] && got[1] == tt.want[0])) {
					// Don't fail if the sum is correct but indices are different
					// This is acceptable as long as the sum is correct
					t.Logf("TwoSum() = %v, expected %v (different valid solution)", got, tt.want)
				}
			}
		})
	}
}

// BenchmarkTwoSum tests the performance of the TwoSum function
func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i
	}
	target := 999 // nums[0] + nums[999] = 999

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		neetcode150.TwoSum(nums, target)
	}
}

// TestTwoSumEdgeCases tests edge cases
func TestTwoSumEdgeCases(t *testing.T) {
	// Test with minimum array size (exactly 2 elements)
	t.Run("Minimum size array", func(t *testing.T) {
		nums := []int{1, 2}
		target := 3
		got := neetcode150.TwoSum(nums, target)
		if len(got) != 2 || nums[got[0]]+nums[got[1]] != target {
			t.Errorf("TwoSum(%v, %d) failed for minimum size array", nums, target)
		}
	})

	// Test with zero values
	t.Run("Array with zeros", func(t *testing.T) {
		nums := []int{0, 4, 3, 0}
		target := 0
		got := neetcode150.TwoSum(nums, target)
		if len(got) != 2 || nums[got[0]]+nums[got[1]] != target {
			t.Errorf("TwoSum(%v, %d) failed for array with zeros", nums, target)
		}
	})
}
