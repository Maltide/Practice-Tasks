package problems

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  bool
	}{
		{name: "empty string", in: "", out: true},
		{name: "single character", in: "A", out: true},
		{name: "simple even palindrome", in: "abba", out: true},
		{name: "mixed case and spaces", in: "A man a plan a canal Panama", out: true},
		{name: "not a palindrome", in: "abc", out: false},
		{name: "palindrome with spaces", in: "No lemon no melon", out: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.in); got != tt.out {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.out)
			}
		})
	}
}
