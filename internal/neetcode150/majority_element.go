package neetcode150

// MajorityElement finds the majority element in an array.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
//
// Time Complexity: O(n)
// Space Complexity: O(1) - optimal goal (O(n) acceptable for first approach)
//
// Pattern: Boyer-Moore Voting Algorithm (optimal) or Hash Map (simpler)
//
// Two approaches:
// 1. Hash Map: Count frequencies, return element with count > n/2
// 2. Boyer-Moore: Track candidate and count, adjust based on current element

func MajorityElement(nums []int) int {
	// TODO: Implement majority element finding

	// Approach (Boyer-Moore - O(1) space):
	// - Initialize candidate and count
	// - For each number:
	//   - If count == 0: set current as candidate
	//   - If current == candidate: increment count
	//   - Else: decrement count
	// - Return candidate

	candidate := nums[0]
	count := 0

	for _, v := range nums {
		if count == 0 {
			candidate = v
			count++
			continue
		}

		if v == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate

}

func MajorityElementHash(nums []int) int {
	// Approach  (Hash Map - O(n) space):
	// - Create a frequency map
	// - Count occurrences of each element
	// - Return element with count > len(nums)/2

	// Hash Map Approach - O(n) time, O(n) space
	frequencies := make(map[int]int)

	// Count frequency of each element
	for _, num := range nums {
		frequencies[num]++
	}

	// Find the element with frequency > n/2
	threshold := len(nums) / 2
	for num, freq := range frequencies {
		if freq > threshold {
			return num
		}
	}

	// This should never be reached given problem constraints
	return 0
}
