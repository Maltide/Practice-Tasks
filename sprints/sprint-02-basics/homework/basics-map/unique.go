// Package problems give opportunity to learn how go works around different structures
package problems

// UniqueValues возвращает новый слайс с уникальными значениями исходного слайса.
//
// Вход:
//   - nums []int: исходный слайс, например []int{1, 2, 2, 3, 1}
//
// Выход:
//   - []int: слайс без повторяющихся элементов, порядок первого появления сохраняется
//
// Пример:
//   - nums = []int{1, 2, 2, 3, 1} -> []int{1, 2, 3}
func UniqueValues(nums []int) []int {
	duplicate := make(map[int]bool)
	out := []int{}
	for _, val := range nums {
		if !duplicate[val] {
			out = append(out, val)
			duplicate[val] = true
		}
	}
	return out
}
