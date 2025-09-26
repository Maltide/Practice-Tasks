package problems

import "testing"

func TestIsBracketSequenceBalanced(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  bool
	}{
		{name: "empty string", in: "", out: true},
		{name: "single pair", in: "()", out: true},
		{name: "nested mixed", in: "([]){}", out: true},
		{name: "crossed brackets", in: "([)]", out: false},
		{name: "extra closing", in: "()]", out: false},
		{name: "only opening", in: "((", out: false},
		{name: "deep nesting", in: "{[()()]}", out: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBracketSequenceBalanced(tt.in); got != tt.out {
				t.Errorf("IsBracketSequenceBalanced(%q) = %v, want %v", tt.in, got, tt.out)
			}
		})
	}
}
