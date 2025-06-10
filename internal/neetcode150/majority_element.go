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
	//
	// Approach 1 (Hash Map - O(n) space):
	// - Create a frequency map
	// - Count occurrences of each element
	// - Return element with count > len(nums)/2
	//
	// Approach 2 (Boyer-Moore - O(1) space):
	// - Initialize candidate and count
	// - For each number:
	//   - If count == 0: set current as candidate
	//   - If current == candidate: increment count
	//   - Else: decrement count
	// - Return candidate

	return 0
}
