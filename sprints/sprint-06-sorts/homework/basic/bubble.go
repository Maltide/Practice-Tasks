package problems

// BubbleSort sorts a slice of integers in ascending order using the bubble sort algorithm.
//
// Algorithm:
// Bubble sort repeatedly steps through the list, compares adjacent elements and
// swaps them if they are in the wrong order. The process is repeated until
// the list is sorted.
//
// Function signature:
//
//	func BubbleSort(arr []int) []int
//
// Input:
//   - arr []int: unsorted slice of integers
//
// Output:
//   - []int: new sorted slice
//
// Examples:
//
//	BubbleSort([]int{64, 34, 25, 12, 22, 11, 90}) → []int{11, 12, 22, 25, 34, 64, 90}
//	BubbleSort([]int{}) → []int{}
//	BubbleSort([]int{5}) → []int{5}
//	BubbleSort([]int{3, 1, 2}) → []int{1, 2, 3}
//
// Complexity:
//   - Time: O(n²) in all cases (best, average, worst)
//   - Space: O(1) - in-place sorting
//
// Features:
//   - Stable sort (does not change relative order of equal elements)
//   - Adaptive - can be optimized to stop early if array is already sorted
//   - Simple to implement but inefficient for large datasets
func BubbleSort(arr []int) []int {
	panic("implement me")
}
