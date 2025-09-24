package problems

import (
	"math"
	"testing"
)

func TestSumAndAverage(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		wantSum int
		wantAvg float64
	}{
		{name: "empty slice", input: []int{}, wantSum: 0, wantAvg: 0},
		{name: "single element", input: []int{5}, wantSum: 5, wantAvg: 5},
		{name: "positive numbers", input: []int{1, 2, 3, 4}, wantSum: 10, wantAvg: 2.5},
		{name: "negative numbers", input: []int{-2, -4, -6}, wantSum: -12, wantAvg: -4},
		{name: "mixed numbers", input: []int{-1, 0, 1, 2}, wantSum: 2, wantAvg: 0.5},
		{name: "with zeros", input: []int{0, 0, 0}, wantSum: 0, wantAvg: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, avg := SumAndAverage(tt.input)
			if sum != tt.wantSum {
				t.Errorf("SumAndAverage(%v) sum = %d, want %d", tt.input, sum, tt.wantSum)
			}
			if math.Abs(avg-tt.wantAvg) > 1e-9 {
				t.Errorf("SumAndAverage(%v) avg = %f, want %f", tt.input, avg, tt.wantAvg)
			}
		})
	}
}
