package neetcode150

import (
	"slices"
)

// GroupAnagrams groups anagrams together from an array of strings.
// An anagram is a word formed by rearranging the letters of another,
// using all original letters exactly once.
//
// Time Complexity: O(n * k * log k) where n = number of strings, k = max length of any string
// Space Complexity: O(n * k) for storing output + keys in the map
//
// Example:
//
//	Input:  ["eat", "tea", "tan", "ate", "nat", "bat"]
//	Output: [["bat"], ["nat","tan"], ["ate","eat","tea"]]
//
// Approach:
// Use a hash map where:
// - Key: sorted version of each word (like "aet" for "eat")
// - Value: list of original words that match the sorted key
func GroupAnagrams(strs []string) [][]string {
	// TODO: Implement group anagrams logic
	sortedStore := make(map[string][]string)

	for _, v := range strs {
		sorted := sortString(v)
		if current, ok := sortedStore[sorted]; ok {
			sortedStore[sorted] = append(current, v)
		} else {
			sortedStore[sorted] = []string{v}
		}
	}

	var group [][]string

	for _, v := range sortedStore {
		group = append(group, v)
	}

	return group
}

func sortString(s string) string {
	// Convert to slice of runes (to support Unicode characters)
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}
