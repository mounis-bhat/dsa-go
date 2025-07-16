package neetcode150_test

import (
	"reflect"
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1: [1,2,3,4]",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "Example 2: [-1,1,0,-3,3]",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "Single element",
			nums: []int{5},
			want: []int{1},
		},
		{
			name: "Two elements",
			nums: []int{2, 3},
			want: []int{3, 2},
		},
		{
			name: "Contains zero at start",
			nums: []int{0, 1, 2, 3},
			want: []int{6, 0, 0, 0},
		},
		{
			name: "Contains zero at end",
			nums: []int{1, 2, 3, 0},
			want: []int{0, 0, 0, 6},
		},
		{
			name: "Contains zero in middle",
			nums: []int{1, 2, 0, 4},
			want: []int{0, 0, 8, 0},
		},
		{
			name: "Multiple zeros",
			nums: []int{0, 0, 2, 3},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "Negative numbers",
			nums: []int{-2, -3, 4, -5},
			want: []int{60, 40, -30, 24},
		},
		{
			name: "All ones",
			nums: []int{1, 1, 1, 1},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "Large numbers",
			nums: []int{10, 20, 30},
			want: []int{600, 300, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := neetcode150.ProductOfArrayExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductOfArrayExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// Benchmark test
func BenchmarkProductOfArrayExceptSelf(t *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i + 1
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		neetcode150.ProductOfArrayExceptSelf(nums)
	}
}
