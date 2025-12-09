package problems

// SelectionSort sorts a slice of integers in ascending order using the selection sort algorithm.
//
// Algorithm:
// Selection sort divides the input list into two parts: a sorted sublist of items
// which is built from left to right at the front of the list, and a sublist of the
// remaining unsorted items. Initially, the sorted sublist is empty and the unsorted
// sublist is the entire input list. The algorithm finds the smallest (or largest,
// depending on sorting order) element in the unsorted sublist, swaps it with the
// leftmost unsorted element (putting it in sorted order), and moves the sublist
// boundaries one element to the right.
//
// Function signature:
//
//	func SelectionSort(arr []int) []int
//
// Input:
//   - arr []int: unsorted slice of integers
//
// Output:
//   - []int: new sorted slice
//
// Examples:
//
//	SelectionSort([]int{64, 34, 25, 12, 22, 11, 90}) → []int{11, 12, 22, 25, 34, 64, 90}
//	SelectionSort([]int{}) → []int{}
//	SelectionSort([]int{5}) → []int{5}
//	SelectionSort([]int{3, 1, 2}) → []int{1, 2, 3}
//
// Complexity:
//   - Time: O(n²) in all cases (best, average, worst)
//   - Space: O(1) - in-place sorting
//
// Features:
//   - Unstable sort (may change relative order of equal elements)
//   - Performs minimum number of swaps (at most n-1)
//   - Not adaptive - always performs same number of comparisons
func SelectionSort(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}
