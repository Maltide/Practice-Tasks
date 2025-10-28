package problems

// InsertionSort sorts a slice of integers in ascending order using the insertion sort algorithm.
//
// Algorithm:
// Insertion sort builds the final sorted array (or list) one item at a time.
// It is much less efficient on large lists than more advanced algorithms such as
// quicksort, heapsort, or merge sort. However, insertion sort provides several advantages:
// simple implementation, efficient for small data sets, adaptive (efficient for data sets
// that are already substantially sorted), stable (does not change the relative order of
// elements with equal keys), in-place (requires only O(1) additional memory space),
// and online (can sort a list as it receives it).
//
// Function signature:
//
//	func InsertionSort(arr []int) []int
//
// Input:
//   - arr []int: unsorted slice of integers
//
// Output:
//   - []int: new sorted slice
//
// Examples:
//
//	InsertionSort([]int{64, 34, 25, 12, 22, 11, 90}) → []int{11, 12, 22, 25, 34, 64, 90}
//	InsertionSort([]int{}) → []int{}
//	InsertionSort([]int{5}) → []int{5}
//	InsertionSort([]int{3, 1, 2}) → []int{1, 2, 3}
//
// Complexity:
//   - Time:
//   - Best case: O(n) - when array is already sorted
//   - Average case: O(n²)
//   - Worst case: O(n²) - when array is sorted in reverse order
//   - Space: O(1) - in-place sorting
//
// Features:
//   - Stable sort
//   - Adaptive - efficient for partially sorted arrays
//   - Efficient for small data sets
//   - In-place - requires minimal additional memory
func InsertionSort(arr []int) []int {
	panic("implement me")
}
