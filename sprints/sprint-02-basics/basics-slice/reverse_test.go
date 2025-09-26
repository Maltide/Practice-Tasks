package problems

import (
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{name: "empty slice", in: []int{}, out: []int{}},
		{name: "single element", in: []int{42}, out: []int{42}},
		{name: "even number of elements", in: []int{1, 2, 3, 4}, out: []int{4, 3, 2, 1}},
		{name: "odd number of elements", in: []int{1, 2, 3, 4, 5}, out: []int{5, 4, 3, 2, 1}},
		{name: "with negatives", in: []int{-1, 0, 1}, out: []int{1, 0, -1}},
		{name: "repeated values", in: []int{1, 1, 2, 2}, out: []int{2, 2, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := append([]int(nil), tt.in...)
			got := ReverseSlice(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("ReverseSlice(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}
