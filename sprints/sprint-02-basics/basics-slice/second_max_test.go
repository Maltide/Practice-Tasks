package problems

import "testing"

func TestSecondLargest(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{name: "simple ascending", in: []int{1, 2, 3}, out: 2},
		{name: "with duplicates of max", in: []int{5, 1, 5, 2}, out: 2},
		{name: "negative numbers", in: []int{-5, -2, -3, -1}, out: -2},
		{name: "mixed order", in: []int{10, 9, 8, 11}, out: 10},
		{name: "max at end", in: []int{4, 2, 3, 5}, out: 4},
		{name: "second max repeated", in: []int{7, 7, 6, 6, 5}, out: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecondLargest(tt.in); got != tt.out {
				t.Errorf("SecondLargest(%v) = %d, want %d", tt.in, got, tt.out)
			}
		})
	}
}
