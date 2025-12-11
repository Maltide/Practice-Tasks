package problems

// CountingSort sorts a slice of integers in ascending order using the counting sort algorithm.
//
// Algorithm:
// Counting sort is a sorting algorithm that works by counting the number of objects
// having distinct key values. It then performs arithmetic on those counts to determine
// the positions of each key value in the output sequence. The algorithm is suitable
// for direct use when the range of keys is known and not significantly greater than
// the number of items. It works efficiently when the range of input values is not
// much larger than the number of values to be sorted.
//
// Function signature:
//
//	func CountingSort(arr []int) []int
//
// Input:
//   - arr []int: unsorted slice of integers
//
// Output:
//   - []int: new sorted slice
//
// Examples:
//
//	CountingSort([]int{4, 2, 2, 8, 3, 3, 1}) → []int{1, 2, 2, 3, 3, 4, 8}
//	CountingSort([]int{}) → []int{}
//	CountingSort([]int{5}) → []int{5}
//	CountingSort([]int{3, 1, 2}) → []int{1, 2, 3}
//
// Complexity:
//   - Time: O(n + k), where n is number of elements and k is range of values
//   - Space: O(k) - requires additional memory for count array
//
// Features:
//   - Stable sort
//   - Not an in-place sorting algorithm
//   - Efficient when range (k) is not much larger than number of elements (n)
//   - Works only with integers
//   - Can be adapted to work with negative numbers
func CountingSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	minInt := arr[0]

	maxint := arr[0]

	for _, val := range arr {
		if val < minInt {
			minInt = val
		}

		if val > maxint {
			maxint = val
		}
	}

	counterslice := make([]int, maxint-minInt+1, maxint-minInt+1)

	for i := range arr {
		counterslice[arr[i]-minInt]++
	}

	resultslice := make([]int, 0, len(arr))

	for i := range counterslice {
		for counterslice[i] > 0 {
			resultslice = append(resultslice, i+minInt)
			counterslice[i]--
		}
	}
	return resultslice
}

// CountingSortPositive сортирует слайс неотрицательных целых чисел по возрастанию
// Это упрощенная версия сортировки подсчетом для положительных чисел
func CountingSortPositive(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	maxint := arr[0]

	for i := range arr {
		if arr[i] > maxint {
			maxint = arr[i]
		}
	}

	counterslice := make([]int, maxint+1, maxint+1)

	for i := range arr {
		counterslice[arr[i]]++
	}

	resultslice := make([]int, 0, len(arr))

	for i := range counterslice {
		for counterslice[i] > 0 {
			resultslice = append(resultslice, i)
			counterslice[i]--
		}
	}

	return resultslice
}
