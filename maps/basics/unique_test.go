package problems

import (
	"reflect"
	"testing"
)

func TestUniqueValues(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{name: "empty slice", in: []int{}, out: []int{}},
		{name: "all unique", in: []int{1, 2, 3}, out: []int{1, 2, 3}},
		{name: "with duplicates", in: []int{1, 2, 2, 3, 1}, out: []int{1, 2, 3}},
		{name: "repeated pattern", in: []int{4, 4, 4, 4}, out: []int{4}},
		{name: "negatives and zeros", in: []int{0, -1, -1, 0, 2}, out: []int{0, -1, 2}},
		{name: "mixed order", in: []int{5, 3, 5, 2, 3, 1}, out: []int{5, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueValues(tt.in); !reflect.DeepEqual(got, tt.out) {
				t.Errorf("UniqueValues(%v) = %v, want %v", tt.in, got, tt.out)
			}
		})
	}
}
