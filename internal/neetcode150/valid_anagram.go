package neetcode150

// ValidAnagram checks if two strings are anagrams of each other
// An anagram is a word or phrase formed by rearranging the letters of another,
// using all original letters exactly once.
// Time Complexity: O(n) where n is the length of the strings
// Space Complexity: O(1) (O(k) for alphabet size, but k is fixed for lowercase letters → O(1))
func ValidAnagram(s, t string) bool {
	// TODO: Implement the ValidAnagram function
	// Hint: You can use character frequency counting
	// Consider edge cases like different lengths

	if len(s) != len(t) {
		return false
	}

	return false
}
