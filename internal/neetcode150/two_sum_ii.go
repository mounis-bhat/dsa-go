package neetcode150

// TwoSumII finds two numbers in a sorted array that add up to the target.
// The array is 1-indexed and sorted in non-decreasing order.
// Returns the indices (1-based) of the two numbers.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
//
// Pattern: Two Pointers
// Since the array is sorted, we can use two pointers:
// - One at the start (left)
// - One at the end (right)
// - Move them inward based on the sum compared to target
func TwoSumII(numbers []int, target int) []int {
	// TODO: Implement the two pointers approach
	// Initialize left pointer at start (index 0)
	// Initialize right pointer at end (index len(numbers)-1)

	// While left < right:
	//   Calculate sum = numbers[left] + numbers[right]
	//   If sum == target: return [left+1, right+1] (convert to 1-based)
	//   If sum < target: move left pointer right (need larger sum)
	//   If sum > target: move right pointer left (need smaller sum)

	// Return empty slice (should never reach here per problem constraints)

	start := 0
	end := len(numbers) - 1

	for start < end {

		sum := numbers[start] + numbers[end]

		if sum == target {
			return []int{start + 1, end + 1}
		}

		if sum > target {
			end--
		} else {
			start++
		}

	}

	return []int{}
}
