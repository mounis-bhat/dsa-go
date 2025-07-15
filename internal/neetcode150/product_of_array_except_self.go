package neetcode150

// ProductOfArrayExceptSelf returns an array where output[i] is the product of all elements except nums[i]
// Constraints:
// - Do not use division
// - Time Complexity: O(n)
// - Space Complexity: O(1) extra space (ignore output array)
func ProductOfArrayExceptSelf(nums []int) []int {
	// TODO: Implement the solution
	// Hint: Use two passes - one for left products, one for right products
	// You can store left products in the result array first,
	// then multiply with right products in the second pass

	prod := make([]int, 0, len(nums))

	for i := range nums {
		currentProduct := 1
		for j, v := range nums {
			if i != j {
				currentProduct = currentProduct * v
			}
		}
		prod = append(prod, currentProduct)
	}

	return prod
}
