package neetcode150_test

import (
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Example 1: anagram and nagaram",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "Example 2: rat and car",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "Single character - same",
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			name: "Single character - different",
			s:    "a",
			t:    "b",
			want: false,
		},
		{
			name: "Empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "Different lengths",
			s:    "abc",
			t:    "ab",
			want: false,
		},
		{
			name: "Same string",
			s:    "listen",
			t:    "listen",
			want: true,
		},
		{
			name: "Classic anagram: listen and silent",
			s:    "listen",
			t:    "silent",
			want: true,
		},
		{
			name: "Repeated characters - valid anagram",
			s:    "aab",
			t:    "aba",
			want: true,
		},
		{
			name: "Repeated characters - invalid anagram",
			s:    "aab",
			t:    "aaa",
			want: false,
		},
		{
			name: "All same characters",
			s:    "aaaa",
			t:    "aaaa",
			want: true,
		},
		{
			name: "Long strings - valid anagram",
			s:    "abcdefghijklmnopqrstuvwxyz",
			t:    "zyxwvutsrqponmlkjihgfedcba",
			want: true,
		},
		{
			name: "Long strings - invalid anagram",
			s:    "abcdefghijklmnopqrstuvwxyz",
			t:    "abcdefghijklmnopqrstuvwxyy",
			want: false,
		},
		{
			name: "Multiple repeated chars",
			s:    "aabbcc",
			t:    "abcabc",
			want: true,
		},
		{
			name: "One character off",
			s:    "abc",
			t:    "abd",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := neetcode150.ValidAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("ValidAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

// BenchmarkValidAnagram tests the performance of the ValidAnagram function
func BenchmarkValidAnagram(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		neetcode150.ValidAnagram(s, t)
	}
}

// TestValidAnagramEdgeCases tests edge cases
func TestValidAnagramEdgeCases(t *testing.T) {
	// Test with minimum string size (exactly 1 character)
	t.Run("Minimum size strings", func(t *testing.T) {
		s := "a"
		target := "a"
		got := neetcode150.ValidAnagram(s, target)
		if !got {
			t.Errorf("ValidAnagram(%q, %q) failed for minimum size strings", s, target)
		}
	})

	// Test with large identical strings
	t.Run("Large identical strings", func(t *testing.T) {
		s := make([]byte, 1000)
		for i := range s {
			s[i] = 'a' + byte(i%26)
		}
		str1 := string(s)
		str2 := string(s)

		got := neetcode150.ValidAnagram(str1, str2)
		if !got {
			t.Errorf("ValidAnagram with large identical strings should return true")
		}
	})

	// Test with large strings that are anagrams
	t.Run("Large anagram strings", func(t *testing.T) {
		s := make([]byte, 1000)
		for i := range s {
			s[i] = 'a' + byte(i%26)
		}
		str1 := string(s)

		// Reverse the string to create an anagram
		reversed := make([]byte, len(s))
		for i := range s {
			reversed[len(s)-1-i] = s[i]
		}
		str2 := string(reversed)

		got := neetcode150.ValidAnagram(str1, str2)
		if !got {
			t.Errorf("ValidAnagram with large anagram strings should return true")
		}
	})
}
