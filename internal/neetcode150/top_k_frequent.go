package neetcode150

import "fmt"

// TopKFrequent returns the k most frequent elements from the given array.
// The answer can be returned in any order.
//
// Time Complexity: O(n log k) using hash map + min heap approach
// Space Complexity: O(n) for the frequency map and heap
//
// Example:
//
//	Input:  nums = [1,1,1,2,2,3], k = 2
//	Output: [1,2] (or [2,1])
//
//	Input:  nums = [1], k = 1
//	Output: [1]
//
// Approach:
// 1. Count frequencies using a map[int]int
// 2. Build a min-heap of size k based on frequency
// 3. Keep top k frequent elements in the heap
// 4. Extract elements from the heap into the result
//
// Note: Go uses container/heap package which requires implementing heap.Interface
func TopKFrequent(nums []int, k int) []int {
	// TODO: Implement top k frequent elements logic

	frequencyMap := make(map[int]int)

	for _, v := range nums {
		count := frequencyMap[v]
		frequencyMap[v] = count + 1
	}

	fmt.Println(frequencyMap)

	return nil
}
