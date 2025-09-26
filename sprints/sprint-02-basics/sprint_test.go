package problems

import (
	"math"
	"reflect"
	"testing"

	// Import functions from subdirectories
	basicsmap "example.com/practice-tasks/sprints/sprint-02-basics/basics-map"
	basicsslice "example.com/practice-tasks/sprints/sprint-02-basics/basics-slice"
	basicsstrings "example.com/practice-tasks/sprints/sprint-02-basics/basics-strings"
	misc "example.com/practice-tasks/sprints/sprint-02-basics/misc"
	star "example.com/practice-tasks/sprints/sprint-02-basics/the-one-with-the-star"
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
			if got := basicsmap.CountFrequencies(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountFrequencies(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

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
			if got := basicsmap.UniqueValues(tt.in); !reflect.DeepEqual(got, tt.out) {
				t.Errorf("UniqueValues(%v) = %v, want %v", tt.in, got, tt.out)
			}
		})
	}
}

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
			got := basicsslice.RemoveAtIndex(tt.nums, tt.idx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAtIndex(%v, %d) = %v, want %v", numsCopy, tt.idx, got, tt.want)
			}
		})
	}
}

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
			got := basicsslice.ReverseSlice(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("ReverseSlice(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}

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
			if got := basicsslice.SecondLargest(tt.in); got != tt.out {
				t.Errorf("SecondLargest(%v) = %d, want %d", tt.in, got, tt.out)
			}
		})
	}
}

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
			sum, avg := basicsslice.SumAndAverage(tt.input)
			if sum != tt.wantSum {
				t.Errorf("SumAndAverage(%v) sum = %d, want %d", tt.input, sum, tt.wantSum)
			}
			if math.Abs(avg-tt.wantAvg) > 1e-9 {
				t.Errorf("SumAndAverage(%v) avg = %f, want %f", tt.input, avg, tt.wantAvg)
			}
		})
	}
}

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
			if got := basicsstrings.AreAnagrams(tt.a, tt.b); got != tt.out {
				t.Errorf("AreAnagrams(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.out)
			}
		})
	}
}

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
			if got := basicsstrings.IsPalindrome(tt.in); got != tt.out {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.out)
			}
		})
	}
}

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
			if got := misc.Fibonacci(tt.n); !reflect.DeepEqual(got, tt.out) {
				t.Errorf("Fibonacci(%d) = %v, want %v", tt.n, got, tt.out)
			}
		})
	}
}

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
			if got := star.IsBracketSequenceBalanced(tt.in); got != tt.out {
				t.Errorf("IsBracketSequenceBalanced(%q) = %v, want %v", tt.in, got, tt.out)
			}
		})
	}
}
