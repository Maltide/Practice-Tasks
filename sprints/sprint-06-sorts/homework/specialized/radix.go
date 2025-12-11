package problems

// RadixSort sorts a slice of integers in ascending order using the radix sort algorithm.
//
// Algorithm:
// Radix sort is a non-comparative sorting algorithm that sorts integers by processing
// individual digits. It sorts numbers digit by digit starting from the least significant
// digit to the most significant digit. A stable counting sort is used for sorting at
// each digit position.
//
// Function signature:
//
//	func RadixSort(arr []int) []int
//
// Input:
//   - arr []int: unsorted slice of integers
//
// Output:
//   - []int: new sorted slice
//
// Examples:
//
//	RadixSort([]int{170, 45, 75, 90, 2, 802, 24, 66}) → []int{2, 24, 45, 66, 75, 90, 170, 802}
//	RadixSort([]int{}) → []int{}
//	RadixSort([]int{5}) → []int{5}
//	RadixSort([]int{3, 1, 2}) → []int{1, 2, 3}
//
// Complexity:
//   - Time: O(d * (n + k)), where d is number of digits in max number, n is number of elements, k is digit value range (typically 10)
//   - Space: O(n + k) - requires additional memory for count arrays and results
//
// Features:
//   - Stable sort
//   - Not an in-place sorting algorithm
//   - Efficient for large datasets with fixed number of digits
//   - Works only with non-negative integers in basic implementation
//   - Can be adapted to work with negative numbers
func RadixSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	maxInt := getMax(arr)

	razryad := 1

	for maxInt/razryad > 0 {
		arr = countingSortByDigit(arr, razryad)
		razryad = razryad * 10
	}

	return arr
}

// getMax возвращает максимальное значение в слайсе
func getMax(arr []int) int {
	maxInt := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxInt {
			maxInt = arr[i]
		}
	}
	return maxInt
}

// countingSortByDigit сортирует слайс по определенному разряду с использованием сортировки подсчетом
func countingSortByDigit(arr []int, exp int) []int {
	counterslice := make([][]int, 10, 10)

	for i := range arr {
		expdigit := (arr[i] / exp) % 10
		if expdigit < 0 {
			return arr
		}
		counterslice[expdigit] = append(counterslice[expdigit], arr[i])
	}

	resSLise := make([]int, 0, len(arr))

	for _, val := range counterslice {
		for _, everydigit := range val {
			resSLise = append(resSLise, everydigit)
		}
	}
	return resSLise
}

// RadixSortExtended версия поразрядной сортировки, которая работает с отрицательными числами
func RadixSortExtended(arr []int) []int {
	panic("implement me")
}
