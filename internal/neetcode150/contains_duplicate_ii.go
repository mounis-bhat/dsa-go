package neetcode150

// ContainsDuplicateII checks if there are two distinct indices i and j
// such that nums[i] == nums[j] and abs(i - j) <= k
// Time Complexity: O(n)
// Space Complexity: O(n)

func ContainsDuplicateII(nums []int, k int) bool {
	countMap := make(map[int]int)

	for i, v := range nums {
		savedIndex, exists := countMap[v]
		if exists && i-savedIndex <= k {
			return true
		}
		countMap[v] = i
	}

	return false
}

func ContainsDuplicateIIBurteForce(nums []int, k int) bool {
	start := 0
	end := min(k, len(nums)-1)

	for end < len(nums) {
		countMap := make(map[int]int)
		for i := start; i <= end; i++ {
			countMap[nums[i]]++
		}

		for _, v := range countMap {
			if v > 1 {
				return true
			}
		}

		start++
		end++
	}

	return false
}
