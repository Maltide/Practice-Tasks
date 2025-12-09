package problems

// MergeSort sorts a slice of integers in ascending order using the merge sort algorithm.
//
// Algorithm:
// Merge sort is a "divide and conquer" algorithm. It divides the input array into two halves,
// recursively sorts them, and then merges the two sorted halves. The merge() algorithm is used
// for merging two halves. The function merge(arr, l, m, r) is the key process that assumes
// that arr[l..m] and arr[m+1..r] are sorted and merges the two sorted subarrays into one.
//
// Function signature:
//
//	func MergeSort(arr []int) []int
//
// Input:
//   - arr []int: unsorted slice of integers
//
// Output:
//   - []int: new sorted slice
//
// Examples:
//
//	MergeSort([]int{64, 34, 25, 12, 22, 11, 90}) → []int{11, 12, 22, 25, 34, 64, 90}
//	MergeSort([]int{}) → []int{}
//	MergeSort([]int{5}) → []int{5}
//	MergeSort([]int{3, 1, 2}) → []int{1, 2, 3}
//
// Complexity:
//   - Time: O(n log n) in all cases (best, average, worst)
//   - Space: O(n) - requires additional memory for temporary arrays
//
// Features:
//   - Stable sort
//   - Not adaptive - always performs same number of operations
//   - Works well with large datasets
//   - Works well with slow storage (disks, tapes) due to sequential access
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[0:mid]
	right := arr[mid:]
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

// mergeSortHelper вспомогательная рекурсивная функция для сортировки слиянием
// func mergeSortHelper(arr []int) []int {
// 	panic("implement me")
// }

// merge объединяет два отсортированных слайса в один отсортированный слайс
func merge(left, right []int) []int {
	result := []int{}

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
