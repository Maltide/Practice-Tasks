package problems

// HeapSort sorts a slice of integers in ascending order using the heapsort algorithm.
//
// Algorithm:
// Heapsort is a comparison-based sorting algorithm based on the heap data structure.
// It is similar to selection sort where we first find the minimum element and place
// it at the beginning. We repeat the same process for the remaining elements.
// The algorithm consists of two main parts:
// 1. Building a max-heap from the input data
// 2. Repeatedly removing the root (maximum element) from the heap and placing it at the end of the array
//
// Function signature:
//
//	func HeapSort(arr []int) []int
//
// Input:
//   - arr []int: unsorted slice of integers
//
// Output:
//   - []int: new sorted slice
//
// Examples:
//
//	HeapSort([]int{64, 34, 25, 12, 22, 11, 90}) → []int{11, 12, 22, 25, 34, 64, 90}
//	HeapSort([]int{}) → []int{}
//	HeapSort([]int{5}) → []int{5}
//	HeapSort([]int{3, 1, 2}) → []int{1, 2, 3}
//
// Complexity:
//   - Time: O(n log n) in all cases (best, average, worst)
//   - Space: O(1) - in-place sorting
//
// Features:
//   - Unstable sort
//   - In-place - requires no additional storage
//   - Not adaptive - always performs same number of operations
//   - Guaranteed O(n log n) time complexity
func HeapSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

	return arr
}

// heapify функция для преобразования поддерева с корнем в индексе i в max-heap
// n - размер кучи
func heapify(arr []int, n, i int) {
	left := 2*i + 1

	right := 2*i + 2

	largest := i

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, n, largest)
	}
}
