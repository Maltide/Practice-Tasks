package problems

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	advancedsort "github.com/Maltide/Practice-Tasks/sprints/sprint-06-sorts/homework/advanced"
	basicsort "github.com/Maltide/Practice-Tasks/sprints/sprint-06-sorts/homework/basic"
	specializedsort "github.com/Maltide/Practice-Tasks/sprints/sprint-06-sorts/homework/specialized"
	sortutils "github.com/Maltide/Practice-Tasks/sprints/sprint-06-sorts/homework/utils"
)

// Test cases for all sorting algorithms
var testCases = []struct {
	name string
	in   []int
	out  []int
}{
	{name: "empty slice", in: []int{}, out: []int{}},
	{name: "single element", in: []int{42}, out: []int{42}},
	{name: "two elements sorted", in: []int{1, 2}, out: []int{1, 2}},
	{name: "two elements unsorted", in: []int{2, 1}, out: []int{1, 2}},
	{name: "three elements", in: []int{3, 1, 2}, out: []int{1, 2, 3}},
	{name: "duplicates", in: []int{4, 2, 4, 1, 2}, out: []int{1, 2, 2, 4, 4}},
	{name: "negative numbers", in: []int{-1, -3, -2, 0}, out: []int{-3, -2, -1, 0}},
	{name: "mixed positive and negative", in: []int{5, -2, 0, -1, 3}, out: []int{-2, -1, 0, 3, 5}},
	{name: "already sorted", in: []int{1, 2, 3, 4, 5}, out: []int{1, 2, 3, 4, 5}},
	{name: "reverse sorted", in: []int{5, 4, 3, 2, 1}, out: []int{1, 2, 3, 4, 5}},
	{name: "large numbers", in: []int{1000000, 999999, 1000001}, out: []int{999999, 1000000, 1000001}},
}

// Stability test case - pairs of (value, original_index)
var stabilityTestCases = []struct {
	name string
	in   [][2]int // [value, original_index]
	out  [][2]int // [value, original_index]
}{
	{
		name: "stability test",
		in:   [][2]int{{3, 0}, {1, 1}, {3, 2}, {2, 3}, {3, 4}},
		out:  [][2]int{{1, 1}, {2, 3}, {3, 0}, {3, 2}, {3, 4}},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.in))
			copy(inputCopy, tt.in)
			got := basicsort.BubbleSort(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("BubbleSort(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.in))
			copy(inputCopy, tt.in)
			got := basicsort.SelectionSort(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("SelectionSort(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.in))
			copy(inputCopy, tt.in)
			got := basicsort.InsertionSort(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("InsertionSort(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.in))
			copy(inputCopy, tt.in)
			got := advancedsort.MergeSort(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("MergeSort(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.in))
			copy(inputCopy, tt.in)
			got := advancedsort.QuickSort(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("QuickSort(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.in))
			copy(inputCopy, tt.in)
			got := advancedsort.HeapSort(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("HeapSort(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}

func TestCountingSort(t *testing.T) {
	// Special test cases for counting sort (works with non-negative integers)
	countingTestCases := []struct {
		name string
		in   []int
		out  []int
	}{
		{name: "empty slice", in: []int{}, out: []int{}},
		{name: "single element", in: []int{5}, out: []int{5}},
		{name: "two elements sorted", in: []int{1, 2}, out: []int{1, 2}},
		{name: "two elements unsorted", in: []int{2, 1}, out: []int{1, 2}},
		{name: "with zeros", in: []int{3, 0, 1, 2}, out: []int{0, 1, 2, 3}},
		{name: "duplicates", in: []int{4, 1, 3, 1, 3}, out: []int{1, 1, 3, 3, 4}},
		{name: "larger range", in: []int{10, 5, 0, 7, 2}, out: []int{0, 2, 5, 7, 10}},
		{name: "with negative numbers (should return original)", in: []int{-1, 2, 3}, out: []int{-1, 2, 3}},
	}

	for _, tt := range countingTestCases {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.in))
			copy(inputCopy, tt.in)
			got := specializedsort.CountingSort(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("CountingSort(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}

func TestRadixSort(t *testing.T) {
	// Special test cases for radix sort (works with non-negative integers)
	radixTestCases := []struct {
		name string
		in   []int
		out  []int
	}{
		{name: "empty slice", in: []int{}, out: []int{}},
		{name: "single element", in: []int{42}, out: []int{42}},
		{name: "two elements sorted", in: []int{12, 23}, out: []int{12, 23}},
		{name: "two elements unsorted", in: []int{23, 12}, out: []int{12, 23}},
		{name: "with zeros", in: []int{102, 0, 21, 12}, out: []int{0, 12, 21, 102}},
		{name: "duplicates", in: []int{123, 23, 123, 2}, out: []int{2, 23, 123, 123}},
		{name: "multi-digit numbers", in: []int{329, 457, 657, 839, 436, 720, 355}, out: []int{329, 355, 436, 457, 657, 720, 839}},
		{name: "with negative numbers (should return original)", in: []int{-1, 2, 3}, out: []int{-1, 2, 3}},
	}

	for _, tt := range radixTestCases {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.in))
			copy(inputCopy, tt.in)
			got := specializedsort.RadixSort(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("RadixSort(%v) = %v, want %v", inputCopy, got, tt.out)
			}
		})
	}
}

// Stability tests for stable sorting algorithms
func TestStability(t *testing.T) {
	// Test bubble sort stability
	t.Run("BubbleSort Stability", func(t *testing.T) {
		// Создаем тестовые данные с парами (значение, оригинальный индекс)
		// Равные значения имеют разные индексы, что позволяет проверить стабильность
		original := [][2]int{{3, 0}, {1, 1}, {3, 2}, {2, 3}, {1, 4}, {3, 5}}
		expected := [][2]int{{1, 1}, {1, 4}, {2, 3}, {3, 0}, {3, 2}, {3, 5}}

		// Извлекаем значения для сортировки
		values := make([]int, len(original))
		for i, pair := range original {
			values[i] = pair[0]
		}

		// Создаем индексы для отслеживания стабильности
		indices := make([]int, len(original))
		for i := range indices {
			indices[i] = i
		}

		// Сортируем индексы с помощью пузырьковой сортировки, сравнивая значения
		sortutils.BubbleSortStability(values, indices)

		// Создаем отсортированный массив пар
		sorted := make([][2]int, len(original))
		for i, idx := range indices {
			sorted[i] = original[idx]
		}

		// Проверяем, что результат соответствует ожидаемому
		for i, pair := range sorted {
			if pair[0] != expected[i][0] || pair[1] != expected[i][1] {
				t.Errorf("BubbleSort stability test failed. Expected %v, got %v", expected[i], pair)
			}
		}
	})

	// Test insertion sort stability
	t.Run("InsertionSort Stability", func(t *testing.T) {
		// Создаем тестовые данные с парами (значение, оригинальный индекс)
		// Равные значения имеют разные индексы, что позволяет проверить стабильность
		original := [][2]int{{3, 0}, {1, 1}, {3, 2}, {2, 3}, {1, 4}, {3, 5}}
		expected := [][2]int{{1, 1}, {1, 4}, {2, 3}, {3, 0}, {3, 2}, {3, 5}}

		// Извлекаем значения для сортировки
		values := make([]int, len(original))
		for i, pair := range original {
			values[i] = pair[0]
		}

		// Создаем индексы для отслеживания стабильности
		indices := make([]int, len(original))
		for i := range indices {
			indices[i] = i
		}

		// Сортируем индексы с помощью сортировки вставками, сравнивая значения
		sortutils.InsertionSortStability(values, indices)

		// Создаем отсортированный массив пар
		sorted := make([][2]int, len(original))
		for i, idx := range indices {
			sorted[i] = original[idx]
		}

		// Проверяем, что результат соответствует ожидаемому
		for i, pair := range sorted {
			if pair[0] != expected[i][0] || pair[1] != expected[i][1] {
				t.Errorf("InsertionSort stability test failed. Expected %v, got %v", expected[i], pair)
			}
		}
	})

	// Test merge sort stability
	t.Run("MergeSort Stability", func(t *testing.T) {
		// Создаем тестовые данные с парами (значение, оригинальный индекс)
		// Равные значения имеют разные индексы, что позволяет проверить стабильность
		original := [][2]int{{3, 0}, {1, 1}, {3, 2}, {2, 3}, {1, 4}, {3, 5}}
		expected := [][2]int{{1, 1}, {1, 4}, {2, 3}, {3, 0}, {3, 2}, {3, 5}}

		// Извлекаем значения для сортировки
		values := make([]int, len(original))
		for i, pair := range original {
			values[i] = pair[0]
		}

		// Создаем индексы для отслеживания стабильности
		indices := make([]int, len(original))
		for i := range indices {
			indices[i] = i
		}

		// Сортируем индексы с помощью сортировки слиянием, сравнивая значения
		sortedIndices := sortutils.MergeSortStability(values, indices)

		// Создаем отсортированный массив пар
		sorted := make([][2]int, len(original))
		for i, idx := range sortedIndices {
			sorted[i] = original[idx]
		}

		// Проверяем, что результат соответствует ожидаемому
		for i, pair := range sorted {
			if pair[0] != expected[i][0] || pair[1] != expected[i][1] {
				t.Errorf("MergeSort stability test failed. Expected %v, got %v", expected[i], pair)
			}
		}
	})
}

// Benchmark tests
func benchmarkSort(b *testing.B, sortFunc func([]int) []int) {
	// Generate random data
	rand.Seed(time.Now().UnixNano())
	data := make([]int, 1000)
	for i := range data {
		data[i] = rand.Intn(1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		input := make([]int, len(data))
		copy(input, data)
		sortFunc(input)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	benchmarkSort(b, basicsort.BubbleSort)
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmarkSort(b, basicsort.SelectionSort)
}

func BenchmarkInsertionSort(b *testing.B) {
	benchmarkSort(b, basicsort.InsertionSort)
}

func BenchmarkMergeSort(b *testing.B) {
	benchmarkSort(b, advancedsort.MergeSort)
}

func BenchmarkQuickSort(b *testing.B) {
	benchmarkSort(b, advancedsort.QuickSort)
}

func BenchmarkHeapSort(b *testing.B) {
	benchmarkSort(b, advancedsort.HeapSort)
}

func BenchmarkCountingSort(b *testing.B) {
	// Special benchmark for counting sort with non-negative integers
	rand.Seed(time.Now().UnixNano())
	data := make([]int, 1000)
	for i := range data {
		data[i] = rand.Intn(1000) // Only non-negative values
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([]int, len(data))
		copy(input, data)
		specializedsort.CountingSort(input)
	}
}

func BenchmarkRadixSort(b *testing.B) {
	// Special benchmark for radix sort with non-negative integers
	rand.Seed(time.Now().UnixNano())
	data := make([]int, 1000)
	for i := range data {
		data[i] = rand.Intn(1000) // Only non-negative values
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([]int, len(data))
		copy(input, data)
		specializedsort.RadixSort(input)
	}
}

func BenchmarkGoStandardSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, 1000)
	for i := range data {
		data[i] = rand.Intn(1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([]int, len(data))
		copy(input, data)
		sort.Ints(input)
	}
}
