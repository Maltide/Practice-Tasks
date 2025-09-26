package problems

import "testing"

func TestAreAnagrams(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		out  bool
	}{
		{name: "both empty", a: "", b: "", out: true},
		{name: "simple anagram", a: "listen", b: "silent", out: true},
		{name: "different lengths", a: "rat", b: "carp", out: false},
		{name: "same letters different counts", a: "aabb", b: "ab", out: false},
		{name: "ignoring spaces and case", a: "Dormitory", b: "Dirty room", out: true},
		{name: "numbers included", a: "123", b: "312", out: true},
		{name: "not an anagram", a: "foo", b: "bar", out: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreAnagrams(tt.a, tt.b); got != tt.out {
				t.Errorf("AreAnagrams(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.out)
			}
		})
	}
}
