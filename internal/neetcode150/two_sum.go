package neetcode150

// TwoSum finds two numbers in the array that add up to the target
// Time Complexity: O(n)
// Space Complexity: O(n)
func TwoSum(nums []int, target int) []int {
	// Hash map to store number -> index mapping
	numToIndex := make(map[int]int)

	for i, num := range nums {
		// Calculate what number we need to find
		complement := target - num

		// Check if the complement exists in our map
		if index, exists := numToIndex[complement]; exists {
			// Found the pair! Return the indices
			return []int{index, i}
		}

		// Store current number and its index for future lookups
		numToIndex[num] = i
	}

	// No solution found (though problem guarantees one exists)
	return nil
}
