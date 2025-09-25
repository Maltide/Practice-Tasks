package problems

// ReverseSlice разворачивает слайс целых чисел "на месте".
//
// Вход:
//   - nums []int: исходный слайс, например []int{1, 2, 3}
//
// Выход:
//   - []int: тот же слайс, но элементы расположены в обратном порядке
//
// Пример:
//   - nums = []int{1, 2, 3} -> []int{3, 2, 1}
func ReverseSlice(nums []int) []int {
	reverse := []int{}

	for i := range nums {
		reverse = append(reverse, nums[len(nums)-1-i])
	}
	nums = reverse
	return nums
}
