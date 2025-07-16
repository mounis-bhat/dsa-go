package neetcode150

// ProductOfArrayExceptSelf returns an array where output[i] is the product of all elements except nums[i]
// Constraints:
// - Do not use division
// - Time Complexity: O(n)
// - Space Complexity: O(1) extra space (ignore output array)
func ProductOfArrayExceptSelf(nums []int) []int {
	// Hint: Use two passes - one for left products, one for right products
	// You can store left products in the result array first,
	// then multiply with right products in the second pass

	// Brute force solution O(n^2)
	// prod := make([]int, 0, len(nums))

	// for i := range nums {
	// 	currentProduct := 1
	// 	for j, v := range nums {
	//		if i != j {
	//			currentProduct = currentProduct * v
	//		}
	//	}
	//	prod = append(prod, currentProduct)
	// }

	left := make([]int, len(nums))
	right := make([]int, len(nums))

	left[0] = 1
	right[len(nums)-1] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	result := make([]int, len(nums))
	for i := range result {
		result[i] = left[i] * right[i]
	}

	return result
}
