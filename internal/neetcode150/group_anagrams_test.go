package neetcode150_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

// Helper function to sort groups for comparison since order doesn't matter
func sortGroups(groups [][]string) {
	// Sort each group internally
	for i := range groups {
		sort.Strings(groups[i])
	}
	// Sort groups by their first element
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 || len(groups[j]) == 0 {
			return len(groups[i]) < len(groups[j])
		}
		return groups[i][0] < groups[j][0]
	})
}

// Helper function to check if two group results are equivalent
func areGroupsEqual(got, want [][]string) bool {
	if len(got) != len(want) {
		return false
	}

	// Create copies to avoid modifying original slices
	gotCopy := make([][]string, len(got))
	wantCopy := make([][]string, len(want))

	for i := range got {
		gotCopy[i] = make([]string, len(got[i]))
		copy(gotCopy[i], got[i])
	}

	for i := range want {
		wantCopy[i] = make([]string, len(want[i]))
		copy(wantCopy[i], want[i])
	}

	// Sort both for comparison
	sortGroups(gotCopy)
	sortGroups(wantCopy)

	return reflect.DeepEqual(gotCopy, wantCopy)
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "Example 1: Basic anagram groups",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name: "Example 2: Single character strings",
			strs: []string{""},
			want: [][]string{
				{""},
			},
		},
		{
			name: "Example 3: Single string",
			strs: []string{"a"},
			want: [][]string{
				{"a"},
			},
		},
		{
			name: "All same letters",
			strs: []string{"abc", "bca", "cab", "acb", "bac", "cba"},
			want: [][]string{
				{"abc", "bca", "cab", "acb", "bac", "cba"},
			},
		},
		{
			name: "No anagrams",
			strs: []string{"abc", "def", "ghi"},
			want: [][]string{
				{"abc"},
				{"def"},
				{"ghi"},
			},
		},
		{
			name: "Empty strings",
			strs: []string{"", "", ""},
			want: [][]string{
				{"", "", ""},
			},
		},
		{
			name: "Mixed empty and non-empty",
			strs: []string{"", "a", "", "b"},
			want: [][]string{
				{"", ""},
				{"a"},
				{"b"},
			},
		},
		{
			name: "Single character anagrams",
			strs: []string{"a", "b", "a", "c", "b"},
			want: [][]string{
				{"a", "a"},
				{"b", "b"},
				{"c"},
			},
		},
		{
			name: "Different length strings",
			strs: []string{"ab", "ba", "abc", "cab", "a"},
			want: [][]string{
				{"ab", "ba"},
				{"abc", "cab"},
				{"a"},
			},
		},
		{
			name: "Case sensitive",
			strs: []string{"Eat", "Tea", "eat", "tea"},
			want: [][]string{
				{"Eat"},
				{"Tea"},
				{"eat", "tea"},
			},
		},
		{
			name: "Numbers as strings",
			strs: []string{"123", "321", "132", "456"},
			want: [][]string{
				{"123", "321", "132"},
				{"456"},
			},
		},
		{
			name: "Special characters",
			strs: []string{"a!b", "b!a", "!ab", "c@d"},
			want: [][]string{
				{"a!b", "b!a", "!ab"},
				{"c@d"},
			},
		},
		{
			name: "Repeated characters",
			strs: []string{"aab", "aba", "baa", "aabb"},
			want: [][]string{
				{"aab", "aba", "baa"},
				{"aabb"},
			},
		},
		{
			name: "Long strings",
			strs: []string{"listen", "silent", "enlist", "hello"},
			want: [][]string{
				{"listen", "silent", "enlist"},
				{"hello"},
			},
		},
		{
			name: "Unicode characters",
			strs: []string{"café", "éfac", "face", "hello"},
			want: [][]string{
				{"café", "éfac"},
				{"face"},
				{"hello"},
			},
		},
		{
			name: "Large group",
			strs: []string{"abc", "bca", "cab", "acb", "bac", "cba", "xyz", "def"},
			want: [][]string{
				{"abc", "bca", "cab", "acb", "bac", "cba"},
				{"xyz"},
				{"def"},
			},
		},
		{
			name: "Palindromes",
			strs: []string{"aba", "bab", "aaa", "bbb"},
			want: [][]string{
				{"aba"},
				{"bab"},
				{"aaa"},
				{"bbb"},
			},
		},
		{
			name: "Single letter repeated",
			strs: []string{"a", "aa", "aaa", "aaaa"},
			want: [][]string{
				{"a"},
				{"aa"},
				{"aaa"},
				{"aaaa"},
			},
		},
		{
			name: "Mixed case anagrams",
			strs: []string{"Stop", "Tops", "pots", "spot"},
			want: [][]string{
				{"Stop"},
				{"Tops"},
				{"pots", "spot"},
			},
		},
		{
			name: "Whitespace strings",
			strs: []string{" ", "  ", " a", "a "},
			want: [][]string{
				{" "},
				{"  "},
				{" a", "a "},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := neetcode150.GroupAnagrams(tt.strs)
			if !areGroupsEqual(got, tt.want) {
				t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
			}
		})
	}
}

// BenchmarkGroupAnagrams tests the performance of the GroupAnagrams function
func BenchmarkGroupAnagrams(t *testing.B) {
	// Create a large test case with many anagrams
	strs := make([]string, 1000)
	baseWords := []string{"listen", "silent", "enlist", "hello", "world", "earth", "heart", "hater"}

	for i := 0; i < 1000; i++ {
		strs[i] = baseWords[i%len(baseWords)]
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		neetcode150.GroupAnagrams(strs)
	}
}

// TestGroupAnagramsEdgeCases tests edge cases
func TestGroupAnagramsEdgeCases(t *testing.T) {
	// Test with very long strings
	t.Run("Very long strings", func(t *testing.T) {
		longStr1 := "abcdefghijklmnopqrstuvwxyz"
		longStr2 := "zyxwvutsrqponmlkjihgfedcba"
		strs := []string{longStr1, longStr2, "short"}

		got := neetcode150.GroupAnagrams(strs)
		expected := [][]string{
			{longStr1, longStr2},
			{"short"},
		}

		if !areGroupsEqual(got, expected) {
			t.Errorf("GroupAnagrams with long strings failed")
		}
	})

	// Test with many duplicate strings
	t.Run("Many duplicates", func(t *testing.T) {
		strs := []string{"abc", "abc", "abc", "bca", "bca", "def"}
		got := neetcode150.GroupAnagrams(strs)
		expected := [][]string{
			{"abc", "abc", "abc", "bca", "bca"},
			{"def"},
		}

		if !areGroupsEqual(got, expected) {
			t.Errorf("GroupAnagrams with duplicates failed")
		}
	})

	// Test with single character repeated many times
	t.Run("Single character repeated", func(t *testing.T) {
		strs := []string{"aaaa", "aaaa", "aaaa"}
		got := neetcode150.GroupAnagrams(strs)
		expected := [][]string{
			{"aaaa", "aaaa", "aaaa"},
		}

		if !areGroupsEqual(got, expected) {
			t.Errorf("GroupAnagrams with repeated single chars failed")
		}
	})

	// Test performance with worst case scenario
	t.Run("Worst case performance", func(t *testing.T) {
		// All strings are anagrams of each other
		strs := []string{"abc", "bca", "cab", "acb", "bac", "cba"}
		got := neetcode150.GroupAnagrams(strs)

		// Should have exactly one group with all strings
		if len(got) != 1 {
			t.Errorf("Expected 1 group, got %d", len(got))
		}
		if len(got[0]) != 6 {
			t.Errorf("Expected 6 strings in group, got %d", len(got[0]))
		}
	})

	// Test with strings that have repeated characters in different orders
	t.Run("Repeated characters different orders", func(t *testing.T) {
		strs := []string{"aabbcc", "abcabc", "ccbbaa", "abccba"}
		got := neetcode150.GroupAnagrams(strs)
		expected := [][]string{
			{"aabbcc", "abcabc", "ccbbaa", "abccba"},
		}

		if !areGroupsEqual(got, expected) {
			t.Errorf("GroupAnagrams with repeated chars failed")
		}
	})

	// Test memory efficiency
	t.Run("Memory efficiency test", func(t *testing.T) {
		// Create strings that would create many different groups
		strs := make([]string, 100)
		for i := 0; i < 100; i++ {
			strs[i] = string(rune('a' + i%26))
		}

		got := neetcode150.GroupAnagrams(strs)

		// Each unique character should form its own group
		// We have 26 unique characters, each repeated
		if len(got) != 26 {
			t.Errorf("Expected 26 groups for unique characters, got %d", len(got))
		}
	})
}
