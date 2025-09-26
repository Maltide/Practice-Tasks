package problems

import (
	"testing"
	"time"

	// Import functions from subdirectories
	findsmallestmap "example.com/practice-tasks/sprints/sprint-01-find-smallest/map"
	findsmallestslice "example.com/practice-tasks/sprints/sprint-01-find-smallest/slice"
)

func TestFindSmallestIntInMap(t *testing.T) {
	tests := []struct {
		name string
		m    map[int]int
		want int
	}{
		{
			name: "simple case",
			m:    map[int]int{1: 3, 2: 1, 3: 4, 4: 1, 5: 5},
			want: 1,
		},
		{
			name: "with negative numbers",
			m:    map[int]int{1: -1, 2: -5, 3: 0, 4: 5, 5: -10},
			want: -10,
		},
		{
			name: "empty map",
			m:    map[int]int{},
			want: 0, // Or handle as an error, for now, we expect 0
		},
		{
			name: "map with one element",
			m:    map[int]int{1: 42},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findsmallestmap.FindSmallestIntInMap(tt.m); got != tt.want {
				t.Errorf("FindSmallestIntInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSmallestIntInMapMap(t *testing.T) {
	tests := []struct {
		name string
		m    map[int]map[int]int
		want int
	}{
		{
			name: "simple case",
			m: map[int]map[int]int{
				1: {1: 10, 2: 5},
				2: {1: 12, 2: 1},
				3: {1: 8, 2: 9},
			},
			want: 1,
		},
		{
			name: "with negative numbers",
			m: map[int]map[int]int{
				1: {1: -1, 2: -5},
				2: {1: 0, 2: -10},
			},
			want: -10,
		},
		{
			name: "empty map",
			m:    map[int]map[int]int{},
			want: 0, // Or handle as an error
		},
		{
			name: "map with one element",
			m:    map[int]map[int]int{1: {1: 42}},
			want: 42,
		},
		{
			name: "map with empty inner map",
			m: map[int]map[int]int{
				1: {},
				2: {1: 100},
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findsmallestmap.FindSmallestIntInMapMap(tt.m); got != tt.want {
				t.Errorf("FindSmallestIntInMapMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSmallestInt(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{
			name:  "simple case",
			slice: []int{3, 1, 4, 1, 5, 9, 2, 6},
			want:  1,
		},
		{
			name:  "with negative numbers",
			slice: []int{-1, -5, 0, 5, -10},
			want:  -10,
		},
		{
			name:  "empty slice",
			slice: []int{},
			want:  0, // Or handle as an error, for now, we expect 0 for an empty slice
		},
		{
			name:  "slice with one element",
			slice: []int{42},
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findsmallestslice.FindSmallestInt(tt.slice); got != tt.want {
				t.Errorf("FindSmallestInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSmallestTime(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name  string
		slice []time.Time
		want  time.Time
	}{
		{
			name: "simple case",
			slice: []time.Time{
				now.Add(3 * time.Hour),
				now.Add(1 * time.Hour),
				now.Add(4 * time.Hour),
			},
			want: now.Add(1 * time.Hour),
		},
		{
			name:  "empty slice",
			slice: []time.Time{},
			want:  time.Time{},
		},
		{
			name:  "slice with one element",
			slice: []time.Time{now},
			want:  now,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findsmallestslice.FindSmallestTime(tt.slice); !got.Equal(tt.want) {
				t.Errorf("FindSmallestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
