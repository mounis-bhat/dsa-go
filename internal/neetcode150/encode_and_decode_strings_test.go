package neetcode150_test

import (
	"reflect"
	"testing"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

func TestEncodeDecode(t *testing.T) {
	codec := neetcode150.Constructor()

	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "Simple case with two strings",
			input: []string{"hello", "world"},
		},
		{
			name:  "List with an empty string",
			input: []string{"", "hello", "world"},
		},
		{
			name:  "List of empty strings",
			input: []string{"", "", ""},
		},
		{
			name:  "Strings with special characters",
			input: []string{"a:b", "c#d", "e;f"},
		},
		{
			name:  "Single string in the list",
			input: []string{"single"},
		},
		{
			name:  "Empty list",
			input: []string{},
		},
		{
			name:  "Long strings",
			input: []string{"a_very_long_string_that_goes_on_and_on", "another_long_string"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := codec.Encode(tt.input)
			decoded := codec.Decode(encoded)

			if !reflect.DeepEqual(tt.input, decoded) {
				t.Errorf("Encode/Decode failed. \nInput: %v, \nDecoded: %v", tt.input, decoded)
			}
		})
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	codec := neetcode150.Constructor()
	input := []string{"hello", "world", "this", "is", "a", "benchmark"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded := codec.Encode(input)
		codec.Decode(encoded)
	}
}
