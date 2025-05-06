package leetcode

import (
	"sort"
	"strings"
)

func ContainsDuplicateBF(nums []int) bool {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func ContainsDuplicateSort(nums []int) bool {
	sort.Ints(nums)

	for i := range len(nums) - 1 {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func ContainsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool)

	for i := range nums {
		if numsMap[nums[i]] {
			return true
		}
		numsMap[nums[i]] = true
	}

	return false
}

func IsAnagram(s string, t string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	t = strings.ToLower(strings.ReplaceAll(t, " ", ""))

	if len(s) != len(t) {
		return false
	}

	charCount := make(map[rune]int)

	for i := range len(s) {
		charCount[rune(s[i])]++
		charCount[rune(t[i])]--
	}

	// for _, ch := range s {
	// 	charCount[ch]++
	// }

	// for _, ch := range t {
	// 	charCount[ch]--
	// }

	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}
