package neetcode150

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
	return nil
}
