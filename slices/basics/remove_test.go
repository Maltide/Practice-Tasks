package problems

import (
	"reflect"
	"testing"
)

func TestRemoveAtIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		idx  int
		want []int
	}{
		{name: "remove from start", nums: []int{1, 2, 3}, idx: 0, want: []int{2, 3}},
		{name: "remove from middle", nums: []int{10, 20, 30, 40}, idx: 2, want: []int{10, 20, 40}},
		{name: "remove from end", nums: []int{5, 6, 7}, idx: 2, want: []int{5, 6}},
		{name: "single element", nums: []int{42}, idx: 0, want: []int{}},
		{name: "repeated values", nums: []int{1, 2, 2, 3}, idx: 1, want: []int{1, 2, 3}},
		{name: "zero index in larger slice", nums: []int{0, 0, 1, 0}, idx: 1, want: []int{0, 1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := append([]int(nil), tt.nums...)
			got := RemoveAtIndex(tt.nums, tt.idx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAtIndex(%v, %d) = %v, want %v", numsCopy, tt.idx, got, tt.want)
			}
		})
	}
}
