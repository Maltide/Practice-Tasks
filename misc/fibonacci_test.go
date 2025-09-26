package problems

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		out  []int
	}{
		{name: "n equals 1", n: 1, out: []int{0}},
		{name: "n equals 2", n: 2, out: []int{0, 1}},
		{name: "first five", n: 5, out: []int{0, 1, 1, 2, 3}},
		{name: "first seven", n: 7, out: []int{0, 1, 1, 2, 3, 5, 8}},
		{name: "first ten", n: 10, out: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{name: "n equals 3", n: 3, out: []int{0, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.n); !reflect.DeepEqual(got, tt.out) {
				t.Errorf("Fibonacci(%d) = %v, want %v", tt.n, got, tt.out)
			}
		})
	}
}
