package problems

import (
	"reflect"
	"sort"
	"testing"

	// Import functions from subdirectories
	mapsprint "example.com/practice-tasks/sprints/sprint-03-advanced/map"
	mixed "example.com/practice-tasks/sprints/sprint-03-advanced/mixed"
	slice "example.com/practice-tasks/sprints/sprint-03-advanced/slice"
	stringsprint "example.com/practice-tasks/sprints/sprint-03-advanced/string"
)

// Test 1: Intersection of Two Slices
func TestIntersectionUnique(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want []int
	}{
		{"basic intersection", []int{1, 2, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{"no overlap", []int{1, 5}, []int{2, 3}, nil},
		{"empty a", nil, []int{1, 2}, nil},
		{"empty b", []int{1, 2}, nil, nil},
		{"both empty", nil, nil, nil},
		{"single element match", []int{1}, []int{1}, []int{1}},
		{"single element no match", []int{1}, []int{2}, nil},
		{"duplicates in both", []int{1, 1, 2, 2}, []int{2, 2, 3, 3}, []int{2}},
		{"all same elements", []int{5, 5, 5}, []int{5, 5}, []int{5}},
		{"negative numbers", []int{-1, -2, 0}, []int{-2, 1, 0}, []int{-2, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.IntersectionUnique(tt.a, tt.b)
			// Sort both slices for comparison since order doesn't matter
			if got != nil {
				sort.Ints(got)
			}
			if tt.want != nil {
				sort.Ints(tt.want)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectionUnique(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// Test 2: Remove Duplicates from Sorted Slice
func TestDedupeSortedInPlace(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantLen  int
		wantNums []int // first wantLen elements
	}{
		{"basic case", []int{1, 1, 2}, 2, []int{1, 2}},
		{"no duplicates", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"all same", []int{1, 1, 1, 1}, 1, []int{1}},
		{"empty slice", []int{}, 0, []int{}},
		{"single element", []int{42}, 1, []int{42}},
		{"long sequence", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{"negative numbers", []int{-3, -3, -1, -1, 0, 0, 2}, 4, []int{-3, -1, 0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			gotLen := slice.DedupeSortedInPlace(numsCopy)
			if gotLen != tt.wantLen {
				t.Errorf("DedupeSortedInPlace(%v) length = %d, want %d", tt.nums, gotLen, tt.wantLen)
			}
			if gotLen > 0 && !reflect.DeepEqual(numsCopy[:gotLen], tt.wantNums) {
				t.Errorf("DedupeSortedInPlace(%v) result = %v, want %v", tt.nums, numsCopy[:gotLen], tt.wantNums)
			}
		})
	}
}

// Test 3: Rotate Slice Right
func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"basic rotation", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"k larger than length", []int{1, 2}, 3, []int{2, 1}},
		{"k equals length", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"k is zero", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"empty slice", []int{}, 1, []int{}},
		{"single element", []int{42}, 5, []int{42}},
		{"negative numbers", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{"large k", []int{1, 2, 3, 4, 5, 6}, 8, []int{5, 6, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.RotateRight(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateRight(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

// Test 4: Maximum Sum Subarray (fixed k)
func TestMaxSumFixedK(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		k       int
		wantSum int
		wantOk  bool
	}{
		{"basic case", []int{1, 4, 2, 10, 23, 3, 1, 0, 20}, 4, 39, true},
		{"k larger than length", []int{2, 3}, 4, 0, false},
		{"k equals length", []int{1, 2, 3}, 3, 6, true},
		{"k is zero", []int{1, 2, 3}, 0, 0, false},
		{"negative k", []int{1, 2, 3}, -1, 0, false},
		{"empty slice", []int{}, 1, 0, false},
		{"single element", []int{5}, 1, 5, true},
		{"all negative", []int{-1, -2, -3, -4}, 2, -3, true},
		{"mixed numbers", []int{2, 1, 5, 1, 3, 2}, 3, 9, true},
		{"zeros", []int{0, 0, 0, 0}, 2, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, gotOk := slice.MaxSumFixedK(tt.nums, tt.k)
			if gotSum != tt.wantSum || gotOk != tt.wantOk {
				t.Errorf("MaxSumFixedK(%v, %d) = (%d, %v), want (%d, %v)", tt.nums, tt.k, gotSum, gotOk, tt.wantSum, tt.wantOk)
			}
		})
	}
}

// Test 5: Merge Two Sorted Slices
func TestMergeSorted(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want []int
	}{
		{"basic merge", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"no overlap", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"empty a", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"empty b", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"both empty", []int{}, []int{}, []int{}},
		{"with duplicates", []int{1, 1, 2}, []int{1, 3, 3}, []int{1, 1, 1, 2, 3, 3}},
		{"negative numbers", []int{-2, 0, 3}, []int{-1, 1, 4}, []int{-2, -1, 0, 1, 3, 4}},
		{"single elements", []int{1}, []int{2}, []int{1, 2}},
		{"a longer than b", []int{1, 3, 5, 7, 9}, []int{2, 4}, []int{1, 2, 3, 4, 5, 7, 9}},
		{"b longer than a", []int{1, 5}, []int{2, 3, 4, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.MergeSorted(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSorted(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// Test 6: Longest Common Prefix
func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{"basic case", []string{"flower", "flow", "flight"}, "fl"},
		{"no common prefix", []string{"dog", "racecar", "car"}, ""},
		{"identical words", []string{"prefix", "prefix"}, "prefix"},
		{"single word", []string{"alone"}, "alone"},
		{"empty slice", []string{}, ""},
		{"empty string in slice", []string{"", "abc"}, ""},
		{"one empty string", []string{"abc", "", "ab"}, ""},
		{"long common prefix", []string{"interspecies", "interstellar", "interstate"}, "inters"},
		{"single character", []string{"a", "a", "a"}, "a"},
		{"no match after first char", []string{"abc", "axy", "azz"}, "a"},
		{"different lengths", []string{"ab", "abc", "abcd"}, "ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringsprint.LongestCommonPrefix(tt.words)
			if got != tt.want {
				t.Errorf("LongestCommonPrefix(%v) = %q, want %q", tt.words, got, tt.want)
			}
		})
	}
}

// Test 7: First Non-Repeating Character
func TestFirstUniqueCharIndex(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"basic case", "leetcode", 0},
		{"unique in middle", "loveleetcode", 2},
		{"no unique chars", "aabb", -1},
		{"empty string", "", -1},
		{"single char", "a", 0},
		{"all same chars", "aaaa", -1},
		{"unique at end", "abccba", -1},
		{"first char repeated later", "abcabc", -1},
		{"complex case", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy", 25},
		{"palindrome", "racecar", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringsprint.FirstUniqueCharIndex(tt.s)
			if got != tt.want {
				t.Errorf("FirstUniqueCharIndex(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

// Test 8: Run-Length Encoding
func TestRLEncode(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"basic case", "aabcccccaaa", "a2bc5a3"},
		{"no repeats", "abcdef", "abcdef"},
		{"all same", "aaaa", "a4"},
		{"empty string", "", ""},
		{"single char", "a", "a"},
		{"alternating", "abab", "abab"},
		{"mixed repeats", "aabbcc", "a2b2c2"},
		{"complex pattern", "aabaa", "a2ba2"},
		{"with spaces", "aa  bb", "a2 2b2"},
		{"numbers", "112233", "1223"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringsprint.RLEncode(tt.s)
			if got != tt.want {
				t.Errorf("RLEncode(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

// Test 9: Run-Length Decoding
func TestRLDecode(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    string
		wantErr bool
	}{
		{"basic case", "a2bc5a3", "aabcccccaaa", false},
		{"no numbers", "abcdef", "abcdef", false},
		{"empty string", "", "", false},
		{"single char", "a", "a", false},
		{"zero count", "a0", "", false},
		{"large count", "a12", "aaaaaaaaaaaa", false},
		{"mixed", "a2b3c", "aabbbc", false},
		{"invalid format - number first", "2a", "", true},
		{"invalid format - negative", "a-1", "", true},
		{"complex valid", "x3y2z", "xxxyyz", false},
		{"single digits", "a1b2c3", "abbccc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringsprint.RLDecode(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("RLDecode(%q) error = %v, wantErr %v", tt.s, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RLDecode(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

// Test 10: Invert Map with Unique Values
func TestInvertUnique(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want map[int]string
	}{
		{
			"basic case",
			map[string]int{"a": 1, "b": 2, "c": 3},
			map[int]string{1: "a", 2: "b", 3: "c"},
		},
		{
			"empty map",
			map[string]int{},
			map[int]string{},
		},
		{
			"single element",
			map[string]int{"hello": 42},
			map[int]string{42: "hello"},
		},
		{
			"negative values",
			map[string]int{"zero": 0, "negative": -5, "positive": 10},
			map[int]string{0: "zero", -5: "negative", 10: "positive"},
		},
		{
			"large values",
			map[string]int{"small": 1, "large": 1000000},
			map[int]string{1: "small", 1000000: "large"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapsprint.InvertUnique(tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertUnique(%v) = %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}

// Test 11: Two Sum
func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want1  int
		want2  int
	}{
		{"basic case", []int{2, 7, 11, 15}, 9, 0, 1},
		{"different order", []int{3, 2, 4}, 6, 1, 2},
		{"duplicates", []int{3, 3}, 6, 0, 1},
		{"negative numbers", []int{-1, -2, -3, -4, -5}, -8, 2, 4},
		{"with zero", []int{0, 4, 3, 0}, 0, 0, 3},
		{"no solution", []int{1, 2}, 4, -1, -1},
		{"large numbers", []int{1000000, 2000000, 3000000}, 5000000, 1, 2},
		{"mixed positive negative", []int{-3, 4, 3, 90}, 0, 0, 2},
		{"solution at end", []int{1, 2, 3, 4, 5}, 9, 3, 4},
		{"first and last", []int{5, 1, 2, 3, 4}, 9, 0, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := mixed.TwoSum(tt.nums, tt.target)
			// Check if we got the expected result (order doesn't matter)
			if (got1 == tt.want1 && got2 == tt.want2) || (got1 == tt.want2 && got2 == tt.want1) {
				return // Test passed
			}
			// Check if both are -1 (no solution case)
			if got1 == -1 && got2 == -1 && tt.want1 == -1 && tt.want2 == -1 {
				return // Test passed
			}
			t.Errorf("TwoSum(%v, %d) = (%d, %d), want (%d, %d)", tt.nums, tt.target, got1, got2, tt.want1, tt.want2)
		})
	}
}
