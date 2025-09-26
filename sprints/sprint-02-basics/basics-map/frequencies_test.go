package problems

import (
	"reflect"
	"testing"
)

func TestCountFrequencies(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want map[string]int
	}{
		{name: "empty slice", in: []string{}, want: map[string]int{}},
		{name: "single item", in: []string{"a"}, want: map[string]int{"a": 1}},
		{name: "repeated items", in: []string{"a", "b", "a", "b", "a"}, want: map[string]int{"a": 3, "b": 2}},
		{name: "case sensitive", in: []string{"Go", "go", "GO"}, want: map[string]int{"Go": 1, "go": 1, "GO": 1}},
		{name: "with empty strings", in: []string{"", "", "value"}, want: map[string]int{"": 2, "value": 1}},
		{name: "multiple uniques", in: []string{"x", "y", "z"}, want: map[string]int{"x": 1, "y": 1, "z": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountFrequencies(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountFrequencies(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
