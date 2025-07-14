package neetcode150

import (
	"container/heap"
)

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

// Implementing a heap

// Struct to save the number and it's frequency
type Node struct {
	num  int
	freq int
}

// FreqHeap is a slice of Freq structs.
// Each Freq struct holds a number and its frequency.
// Example visualization:
//
//	FreqHeap: [
//	  {num: 1, freq: 3},
//	  {num: 2, freq: 2},
//	  {num: 3, freq: 1},
//	]
type FreqHeap []Node

func (f FreqHeap) Len() int {
	return len(f)
}

func (f FreqHeap) Less(i, j int) bool {
	return f[i].freq > f[j].freq // Max heap: higher frequency has higher priority
}

func (f FreqHeap) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *FreqHeap) Push(x any) {
	*f = append(*f, x.(Node))
}

func (f *FreqHeap) Pop() any {
	old := *f
	n := len(old)
	val := old[n-1]
	*f = old[:n-1]
	return val
}

func TopKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int)

	for _, v := range nums {
		count := frequencyMap[v]
		frequencyMap[v] = count + 1
	}

	h := &FreqHeap{}
	heap.Init(h)

	for num, freq := range frequencyMap {
		heap.Push(h, Node{num: num, freq: freq})
	}

	topK := make([]int, 0, k)

	for range k {
		popped := heap.Pop(h).(Node)
		topK = append(topK, popped.num)
	}

	return topK
}
